package base_test

import (
	"fmt"
	"testing"

	"github.com/ArkNX/ark-go/base"
)

func TestEndpoint(t *testing.T) {
	ep, err := base.NewFromString("tcp://0.0.0.0:9001")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(ep)
}
