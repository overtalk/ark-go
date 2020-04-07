package utils_test

import (
	"fmt"
	"github.com/ArkNX/ark-go/utils"
	"testing"
)

func TestBitmap(t *testing.T) {
	bm := utils.NewBitSet()

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
