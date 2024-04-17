# client-go内存优化

### 修改client-go的序列化方式

client-go默认的序列化方式是json，当数据量较大的情况下，json的序列化消耗的资源较多，容易导致服务的内存增长。

这种情况，可以使用protobuf的方式序列化，具体操作如下：

```go
import (
	"k8s.io/client-go/metadata"
)

func createClient() 
.....
configCopy := metadata.ConfigFor(config)
.....
```

