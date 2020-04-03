package util_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"

	. "github.com/ArkNX/ark-go/util"
)

func TestA(t *testing.T) {
	q := NewLockFreeQueue(1)
	for i := 0; i < 10; i++ {
		q.Put(i)
	}

	for {
		i, flag := q.Get()
		if !flag {
			break
		}
		fmt.Println(i)
	}
}

func testQueueHigh(grp, cnt int) {
	var wg sync.WaitGroup

	wg.Add(grp)
	q := NewLockFreeQueue(1024 * 1024)
	for i := 0; i < grp; i++ {
		go func(g int) {
			for j := 0; j < cnt; j++ {
				q.Put(j)
			}
			wg.Done()
		}(i)
	}
	wg.Add(grp)
	for i := 0; i < grp; i++ {
		go func(g int) {
			for j := 0; j < cnt; j++ {
				q.Get()
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func TestQueue(t *testing.T) {
	var Sum int
	var Use time.Duration
	for i := 40; i <= runtime.NumCPU()*4; i++ {
		cnt := 10000 * 1000
		if i > 9 {
			cnt = 10000 * 100
		}
		sum := i * cnt
		start := time.Now()
		testQueueHigh(i, cnt)
		end := time.Now()
		use := end.Sub(start)
		op := use / time.Duration(sum)
		fmt.Printf("%v, Grp: %3d, Times: %10d, use: %12v, %8v/op\n",
			runtime.Version(), i, sum, use, op)
		Use += use
		Sum += sum
	}
	op := Use / time.Duration(Sum)
	fmt.Printf("%v %v, Grp: %3v, Times: %10d, miss:%6v, use: %12v, %8v/op\n",
		runtime.Version(), runtime.GOARCH, "Sum", Sum, 0, Use, op)
}
