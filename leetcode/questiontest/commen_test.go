package questiontest

import (
	"fmt"
	"sync"
	"testing"

	"github.com/hutaochu/note/leetcode/questionbank"
)

func TestSumTarget(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 6, 8, 9}
	res := questionbank.SumTarget(nums, 15)
	fmt.Println(res)
}

func TestLongestPalindromeSubStr(t *testing.T) {
	res := questionbank.LongestPalindromeSubStr("abcaacdabcdcbaaaaaaaaaaa")
	fmt.Println(res)
}

func TestFindMinCoveringSubStr(t *testing.T) {
	s := "ADOBECODEBANC"
	s1 := "ABC"
	fmt.Println(questionbank.FindMinCoveringSubStr(s, s1))
	s = "a"
	s1 = "a"
	fmt.Println(questionbank.FindMinCoveringSubStr(s, s1))
	s1 = "aa"
	fmt.Println(questionbank.FindMinCoveringSubStr(s, s1))
}

func TestFindSubStr(t *testing.T) {
	s := "ADOBECODEBANC"
	s1 := "NC"
	fmt.Println(questionbank.FindSubStr(s1, s))
}

func TestCoinsChange(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(10)
	jobs := make(chan int, 10)

	go func() {
		for {
			_, ok := <-jobs
			if ok {
				println("a")
				wg.Done()
			} else {
				return
			}
		}
	}()
	go func() {
		for {
			_, ok := <-jobs
			if ok {
				println("b")
				wg.Done()
			} else {
				return
			}
		}
	}()
	go func() {
		for {
			_, ok := <-jobs
			if ok {
				println("c")
				wg.Done()
			} else {
				return
			}
		}
	}()

	for i := 0; i < 10; i++ {
		jobs <- i
	}

	wg.Wait()
	close(jobs)
}
