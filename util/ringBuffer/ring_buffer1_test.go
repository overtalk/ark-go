package ringbuffer_test

import (
	"testing"

	"github.com/ArkNX/ark-go/util/ringBuffer"
)

func TestRingBuffer_ByteBuffer(t *testing.T) {
	rf := ringbuffer.New(12)
	t.Log("ring buffer length = ", rf.Len())
	t.Log("ring buffer length = ", rf.IsEmpty())

	toWrite := []byte("qinhanyes")
	n, err := rf.Write(toWrite)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(n)
	t.Log("ring buffer length = ", rf.Len())
	t.Log("ring buffer length = ", rf.IsEmpty())

	header, tail := rf.LazyReadAll()
	t.Log("header = ", string(header))
	t.Log("tail = ", string(tail))
	//
	for i := 0; i < 3; i++ {
		bit, err := rf.ReadByte()
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(string(bit))
	}

	n, err = rf.Write(toWrite)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(n)

	header1, tail1 := rf.LazyReadAll()
	t.Log("header = ", string(header1))
	t.Log("tail = ", string(tail1))

	////////////////////////////////
	n, err = rf.Write(toWrite)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(n)
	n, err = rf.Write(toWrite)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(n)
	t.Log("ring buffer length = ", rf.Len())
	t.Log("ring buffer length = ", rf.IsEmpty())
}
