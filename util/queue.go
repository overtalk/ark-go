package util

import (
	"sync"
)

type LockFreeQueue struct {
	mux        sync.Mutex
	readIndex  uint8
	writeIndex uint8
	cap        int
	chanelArr  []chan interface{}
}

func NewLockFreeQueue(cap int) *LockFreeQueue {
	return &LockFreeQueue{
		mux: sync.Mutex{},
		cap: cap * 2,
		chanelArr: []chan interface{}{
			make(chan interface{}, cap),
		},
	}
}

func (queue *LockFreeQueue) Put(value interface{}) {
	select {
	case queue.chanelArr[queue.writeIndex] <- value:
	default:
		queue.mux.Lock()
		queue.cap += queue.cap
		queue.chanelArr = append(queue.chanelArr, make(chan interface{}, queue.cap))
		queue.writeIndex++
		queue.mux.Unlock()
		queue.Put(value)
	}

}

func (queue *LockFreeQueue) Get() (interface{}, bool) {
	select {
	case ret := <-queue.chanelArr[queue.readIndex]:
		return ret, true
	default:
		queue.mux.Lock()
		if queue.writeIndex <= queue.readIndex {
			queue.mux.Unlock()
			return nil, false
		}
		queue.readIndex++
		queue.mux.Unlock()

		return queue.Get()
	}
}
