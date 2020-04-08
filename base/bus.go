package base

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cast"
)

// BusAddr defines the bus address for Ark Frame
// like IP address, 8.8.8.8
type BusAddr struct {
	ChannelId uint8 // Android,Apple
	ZoneId    uint8 // China,US...
	AppType   uint8 // game,pvp
	InstId    uint8 // instance id
}

func NewBusAddr(cId uint8, zId uint8, pId uint8, iId uint8) *BusAddr {
	return &BusAddr{
		ChannelId: cId,
		ZoneId:    zId,
		AppType:   pId,
		InstId:    iId,
	}
}

func NewBusAddrFromUInt32(id uint32) *BusAddr {
	return &BusAddr{
		ChannelId: uint8(id >> 24),
		ZoneId:    uint8(id >> 16),
		AppType:   uint8(id >> 8),
		InstId:    uint8(id),
	}
}

func NewBusAddrFromStr(busName string) (*BusAddr, error) {
	if busName == "" {
		return nil, errors.New("bus name is empty")
	}

	strArr := strings.Split(busName, ".")
	if len(strArr) != 4 {
		return nil, errors.New("bus id ` " + busName + " ` is invalid, it likes 8.8.8.8")
	}

	var uint8Arr [4]uint8
	for index, str := range strArr {
		i, err := cast.ToUint8E(str)
		if err != nil {
			return nil, err
		}
		uint8Arr[index] = i
	}

	return NewBusAddr(uint8Arr[0], uint8Arr[1], uint8Arr[2], uint8Arr[3]), nil
}

func (a *BusAddr) BusID() uint32 {
	return binary.BigEndian.Uint32([]uint8{a.ChannelId, a.ZoneId, a.AppType, a.InstId})
}

func (a *BusAddr) ToString() string {
	return fmt.Sprintf("%d.%d.%d.%d", a.ChannelId, a.ZoneId, a.AppType, a.InstId)
}

// bus relation, app connect other app with direct way or waiting sync message
type BusRelation struct {
	AppType        uint8
	TargetAppType  uint8
	ConnectionType bool
}

type ProcConfig struct {
	BusId         uint32
	MaxConnection uint32
	ThreadNum     uint8
	IntranetEp    Endpoint
	ServerEp      Endpoint
	// to add other fields
}

type RegCenter struct {
	Ip            string
	Port          uint16
	ServiceName   string
	CheckInterval time.Duration
	CheckTimeout  time.Duration
}

type AppConfig struct {
	RegCenter           RegCenter
	Name2types          map[string]AppType
	Type2names          map[AppType]string
	ConnectionRelations map[AppType][]AppType
	SelfProc            ProcConfig
	OtherProcList       map[uint32]ProcConfig // bus id -> ProcConfig
}
