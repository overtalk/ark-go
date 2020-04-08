package netServiceManagerModule_test

import (
	"fmt"
	"testing"

	"github.com/ArkNX/ark-go/plugin/netPlugin/netServiceManagerModule"
)

func TestCopy(t *testing.T) {
	bytes1 := []byte("hasdfasdfahah")
	bytes2 := []byte("asdfas")

	bytes2 = make([]byte, len(bytes1))

	copy(bytes2, bytes1)

	fmt.Println("bytes1 = ", string(bytes1))
	fmt.Println("bytes1 = ", string(bytes2))

	bytes2 = []byte("sd")

	fmt.Println("bytes1 = ", string(bytes1))
	fmt.Println("bytes1 = ", string(bytes2))
}

func TestNewNetMsg(t *testing.T) {
	netMsg := netServiceManagerModule.NewNetMsgFromData([]byte("wsqinhan"))
	netMsg1 := netServiceManagerModule.NewNetMsgFromMetMsg(netMsg)

	t.Log(netMsg)
	t.Log(netMsg1)

	netMsg1.SetActorID(12)

	t.Log(netMsg)
	t.Log(netMsg1)
}
