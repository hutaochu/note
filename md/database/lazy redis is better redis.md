## [Lazy Redis is better Redis](http://antirez.com/news/93)

[antirez](http://antirez.com/user/antirez) 3052 days ago. 335798 views.

Everybody knows Redis is single threaded. The best informed ones will tell you that, actually, Redis is *kinda* single threaded, since there are threads in order to perform certain slow operations on disk. So far threaded operations were so focused on I/O that our small library to perform asynchronous tasks on a different thread was called bio.c: Background I/O, basically.
However some time ago I opened an issue where I promised a new Redis feature that many wanted, me included, called “lazy free”. The original issue is here: https://github.com/antirez/redis/issues/1748.
The gist of the issue is that Redis DEL operations are normally blocking, so if you send Redis “DEL mykey” and your key happens to have 50 million objects, the server will block for seconds without serving anything in the meantime. Historically this was accepted mostly as a side effect of the Redis design, but is a limit in certain use cases. DEL is not the only blocking command, but is a special one, since usually we say: Redis is very fast as long as you use O(1) and O(log_N) commands. You are free to use O(N) commands but be aware that it’s not the case we optimized for, be prepared for latency spikes.
This sounds reasonable, but at the same time, even objects created with fast operations need to be deleted. And in this case, Redis blocks.
The first attempt
—
In a single-threaded server the easy way to make operations non-blocking is to do things incrementally instead of stopping the world. So if there is to free a 1 million allocations, instead of blocking everything in a for() loop, we can free 1000 elements each millisecond, for example. The CPU time used is the same, or a bit more, since there is more logic, but the latency from the point of view of the user is ways better. Maybe those cycles to free 1000 elements per millisecond were not even used. Avoiding to block for seconds is the key here. This is how many things inside Redis work: LRU eviction and keys expires are two obvious examples, but there are more, like incremental rehashing of hash tables.
So this was the first thing I tried: create a new timer function, and perform the eviction there. Objects were just queued into a linked list, to be reclaimed slowly and incrementally each time the timer function was called. This requires some trick to work well. For example objects implemented with hash tables were also reclaimed incrementally using the same mechanism used inside Redis SCAN command: taking a cursor inside the dictionary and iterating it to free element after element. This way, in each timer call, we don’t have to free a whole hash table. The cursor will tell us where we left when we re-enter the timer function.
Adaptive is hard
—
Do you know what is the hard part with this? That this time, we are doing a very special task incrementally: we are freeing memory. So if while we free memory incrementally, the server memory raises very fast, we may end, for the sake of latency, to consume an *unbound* amount of memory. Which is very bad. Imagine this, for example:
    WHILE 1
        SADD myset element1 element2 … many many many elements
        DEL myset
    END
If deleting myset in the background is slower compared to our SADD call adding tons of elements per call, our memory usage will grow forever.
However after a few experiments, I found a way to make it working very well. The timer function used two ideas in order to be adaptive to the memory pressure:

1. Check the memory tendency: it is raising or lowering? In order to adapt how aggressively to free.
2. Also adapt the timer frequency itself based on “1”, so that we don’t waste CPU time when there is little to free, with continuous interruptions of the event loop. At the same time the timer could reach ~300 HZ when really needed.
    A small portion of code, from the now no longer existing function implementing this ideas:
   /* Compute the memory trend, biased towards thinking memory is raising
    * for a few calls every time previous and current memory raise. */
   if (prev_mem < mem) mem_trend = 1;
      mem_trend *= 0.9; /* Make it slowly forget. */
      int mem_is_raising = mem_trend > .1;
      /* Free a few items. */
      size_t workdone = lazyfreeStep(LAZYFREE_STEP_SLOW);
      /* Adjust this timer call frequency according to the current state. */
      if (workdone) {
       if (timer_period == 1000) timer_period = 20;
       if (mem_is_raising && timer_period > 3)
           timer_period--; /* Raise call frequency. */
       else if (!mem_is_raising && timer_period < 20)
           timer_period++; /* Lower call frequency. */
      } else {
       timer_period = 1000;    /* 1 HZ */
      }
    It’s a good trick and it worked very well. But still it was kinda sad we have to do this operation in a single thread. There was a lot of logic to handle that well, and anyway when the lazy free cycle was very busy, operations per second were reduced to around 65% of the norm.
    Freeing objects in a different thread would be much simpler: freeing is almost always faster than adding new values in the dataset, if there is a thread which is busy doing only free operations.
    For sure there is some contention between the main thread calling the allocator and the lazy free thread doing the same, but Redis spends a fraction of its time on allocations, and much more time on I/O, command dispatching, cache misses, and so forth.
    However there was a big problem with implementing a threaded lazy free: Redis itself. The internal design was totally biased towards sharing objects around. After all they are reference counted right? So why don’t share as much as possible? We can save memory and time. A few examples: if you do SUNIONSTORE you end with shared objects in the target set. Similarly, client output buffers have lists of objects to send to the socket as reply, so during a call like SMEMBERS all the set members may end shared in the output buffer list. So sharing objects sounds so useful, lovely, wonderfully, greatly cool.
    But, hey, there is something more here. If after an SUNIONSTORE I re-load the database, the objects will be unshared, so the memory suddenly may jump to more than it was. Not great. Moreover what happens when we send replies to clients? We actually *glue together* objects into plain buffers when they are small, since otherwise doing many write() calls is not efficient! (free hint, writev() does not help). So we are mostly copying already. And in programming when something is not useful, but exists, it is likely a problem.
    And indeed each time you had to access a value, inside a key containing an aggregate data type, you had to traverse the following:
      key -> value_obj -> hash table -> robj -> sds_string
    So what about getting rid of “robj” structures entirely and converting aggregate values to be made just of hash tables (or skiplists) of SDS strings? (SDS is the library we use inside Redis for strings).
    There is a problem with that. Immagine a command like SADD myset myvalue. We can’t take client->argv[2], for example, and just reference it in the hash table implementing the set. We have to *duplicate* values sometimes and can’t reuse the ones already existing in the client argument vector, created when the command was parsed. However Redis performance is dominated by cache misses, so maybe we can compensate this with one indirection less?
    So I started to work to this new lazyfree branch, and tweeting about it on Twitter without any context so that everybody was thinking I was like desperate or crazy (a few eventually asked WTF this lazyfree thing was). So what I did?
3. Change client output buffers to use just dynamic strings instead of robj structures. Values are always copied when there is to create a reply.
4. Convert all the Redis data types to use SDS strings instead of shared robj structures. Sounds trivial? ~800 highly bug sensitive lines changed in the course of multiple weeks. But now all tests are passing.
5. Rewrite lazyfree to be threaded.
    The result is that Redis is now more memory efficient since there are no robj structures around in the implementation of the data structures (but they are used in the code paths where there is a lot of sharing going on, for example during the command dispatch and replication). Threaded lazy free works great and is faster than the incremental one to reclaim memory, even if the implementation of the incremental one is something I like a lot and was not so terrible even compared with the threaded one. But now, you can delete a huge key and the performance drop is negligible which is very useful. But, the most interesting thing is, Redis is now faster in all the operations I tested so far. Less indirection was a real winner here. It is faster even in unrelated benchmarks just because the client output buffers are now simpler and faster.
    In the end I deleted the incremental lazy freeing implementation from the branch, to retain only the threaded one.
    A note about the API
    —
    However what about the API? We still have a blocking DEL, the default is the same, since DEL in Redis means: reclaim memory now. I didn’t like the idea of changing that. So now you have a new command called UNLINK which more clearly states what is happening to the value.
    UNLINK is a smart command: it calculates the deallocation cost of an object, and if it is very small it will just do what DEL is supposed to do and free the object ASAP. Otherwise the object is sent to the background queue for processing. Otherwise the two commands are identical from the point of view of the keys space semantics.
    FLUSHALL / FLUSHDB non blocking variants were also implemented, but still not at API level, they’ll just take a LAZY option that if given will change the behavior.
    Not just lazy freeing
    —
    Now that values of aggregated data types are fully unshared, and client output buffers don’t contain shared objects as well, there is a lot to exploit. For example it is finally possible to implement threaded I/O in Redis, so that different clients are served by different threads. This means that we’ll have a global lock only when accessing the database, but the clients read/write syscalls and even the parsing of the command the client is sending, can happen in different threads. This is a design similar to memcached, and one I look forward to implement and test.
    Moreover it is now possible to implement certain slow operations on aggregated data types in another thread, in a way that only a few keys are “blocked” but all the other clients can continue. This can be achieved in a very similar way to what we do currently with blocking operations (see blocking.c), plus an hash table to store what keys are currently busy and with what client. So if a client asks for something like SMEMBERS it is possible to lock just the key, process the request creating the output buffer, and later release the key again. Only clients trying to access the same key will be blocked if the key is blocked.
    All this requires even more drastic internal changes, but the bottom line here is, we have a taboo less. We can compensate object copying times with less cache misses and a smaller memory footprint for aggregated data types, so we are now free to think in terms of a threaded Redis with a share-nothing design, which is the only design that could easily outperform our single threaded one. In the past a threaded Redis was always seen as a bad idea if thought as a set of mutexes in data structures and objects in order to implement concurrent access, but fortunately there are alternatives to get the best of both the worlds. And we have the option to still serve all the fast operations like we did in the past from the main thread, if we want. There should be only to gain performance-wise, at the cost of some contained complexity.
    ETA
    —
    I touched a lot of internals, this is not something which is going live tomorrow. So my plan is to call 3.2. what we have already into unstable, do the work to put it into Release Candidate state, and merge this branch into unstable targeting 3.4.
    However before merging, a very scrupulous check for speed regression should be performed. There is definitely more work to do.
    If you want to give it a try check the “lazyfree” branch on Github. Btw be aware that currently I’m working at it very actively so certain things may be at moments totally broken.



### 总结

1. Redis不是真正的单线程，只是指令的执行在一个线程，但是Redis也有其他的线程在工作，比如：写磁盘、清理数据、处理过期key等；
2. 当delete大key的时候，可能导致Redis阻塞；
3. 因为单线程模型很简单，所以Redis早期选择了单线程模型；