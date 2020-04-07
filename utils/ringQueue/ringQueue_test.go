package ringQueue_test

import (
	"github.com/ArkNX/ark-go/utils/ringQueue"
	"reflect"
	"testing"
)

func TestRingQueue(t *testing.T) {
	queue := ringQueue.New(12)
	t.Log(queue.GetBufferLength())

	toWrite := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if _, err := queue.Push(toWrite); err != nil {
		t.Error(err)
		return
	}

	queue.PopOne()
	t.Log(queue.GetBufferLength())

	if _, err := queue.Push(toWrite); err != nil {
		t.Error(err)
		return
	}

	for k, v := range queue.LazyPopAll() {
		t.Log(k, reflect.TypeOf(v), v)
	}
}
