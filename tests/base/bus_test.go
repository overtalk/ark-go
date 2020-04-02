package base_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/ArkNX/ark-go/base"
)

func TestBusAddr(t *testing.T) {
	var (
		ChannelId uint8 = 1
		ZoneId    uint8 = 32
		AppType   uint8 = 45
		InstId    uint8 = 1
	)
	bus := base.NewBusAddr(ChannelId, ZoneId, AppType, InstId)
	busID := bus.BusID()
	fmt.Println(busID)

	bus1 := base.NewBusAddrFromUInt32(busID)
	fmt.Println(bus1.ToString())
}

func TestBusAddr_FromString(t *testing.T) {
	busStr := "12.32.43.67"
	bus, err := base.NewBusAddrFromStr(busStr)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(bus.ToString())
}

func TestA(t *testing.T) {
	arch := runtime.GOARCH
	os := runtime.GOOS
	fmt.Println(arch, os)
}
