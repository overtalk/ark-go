package util_test

import (
	"fmt"
	"github.com/ArkNX/ark-go/util"
	"testing"
)

func TestBitmap(t *testing.T) {
	bm := util.NewBitSet()

	for i := uint64(0); i < 10; i++ {
		bm.Add(i)
	}

	ret, err := bm.MarshalJSON()
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(string(ret))
}
