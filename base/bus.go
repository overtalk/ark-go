package base

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cast"
)

// AFBusAddr defines the bus address for Ark Frame
// like IP address, 8.8.8.8
type AFBusAddr struct {
	ChannelId uint8 // Android,Apple
	ZoneId    uint8 // China,US...
	AppType   uint8 // game,pvp
	InstId    uint8 // instance id
}

func NewAFBusAddr(cId uint8, zId uint8, pId uint8, iId uint8) *AFBusAddr {
	return &AFBusAddr{
		ChannelId: cId,
		ZoneId:    zId,
		AppType:   pId,
		InstId:    iId,
	}
}

func NewAFBusAddrFromUInt32(id uint32) *AFBusAddr {
	return &AFBusAddr{
		ChannelId: uint8(id >> 24),
		ZoneId:    uint8(id >> 16),
		AppType:   uint8(id >> 8),
		InstId:    uint8(id),
	}
}

func NewAFBusAddrFromStr(busName string) (*AFBusAddr, error) {
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

	return NewAFBusAddr(uint8Arr[0], uint8Arr[1], uint8Arr[2], uint8Arr[3]), nil
}

func (a *AFBusAddr) BusID() uint32 {
	return binary.BigEndian.Uint32([]uint8{a.ChannelId, a.ZoneId, a.AppType, a.InstId})
}

func (a *AFBusAddr) ToString() string {
	return fmt.Sprintf("%d.%d.%d.%d", a.ChannelId, a.ZoneId, a.AppType, a.InstId)
}

// bus relation, app connect other app with direct way or waiting sync message
type AFBusRelation struct {
	AppType        uint8
	TargetAppType  uint8
	ConnectionType bool
}

type AFProcConfig struct {
	BusId         int
	MaxConnection uint32
	ThreadNum     uint8
	IntranetEp    AFEndpoint
	ServerEp      AFEndpoint
	// to add other fields
}

type AFRegCenter struct {
	Ip            string
	Port          uint16
	ServiceName   string
	CheckInterval time.Duration
	CheckTimeout  time.Duration
}

type AFAppConfig struct {
	RegCenter           AFRegCenter
	Name2types          map[string]ARKAppType
	Type2names          map[ARKAppType]string
	ConnectionRelations map[ARKAppType][]ARKAppType
	SelfProc            AFProcConfig
	OtherProcList       map[uint32]AFProcConfig // bus id -> AFProcConfig
}
