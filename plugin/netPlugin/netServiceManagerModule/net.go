package netServiceManagerModule

import (
	"fmt"
)

// INet defines the both client & server part of different net protocol
// tcp, ws...
type INet interface {
	Update()

	StartClient(
		headLen HeaderLength,
		dstBusID uint32,
		ip string,
		port uint16,
		isIpv6 bool,
	) error

	StartServer(
		headLen HeaderLength,
		busID uint32,
		ip string,
		port uint16,
		threadNum uint8,
		maxClient uint32,
		isIpv6 bool,
	) error

	Shutdown() error

	// server
	SendMsg(head MsgHead, msgData []byte, sessionID int64) error
	BroadcastMsg(head MsgHead, msgData []byte) error

	CloseSession(sessionID int64) error

	IsWorking() bool
	SetWorking(value bool)
}

////////////////////////////////
// Net defines the base inet
////////////////////////////////
type Net struct {
	working           bool
	StatisticRecvSize uint64
	StatisticSendSize uint64
}

func (net *Net) StartClient(
	headLen HeaderLength,
	dstBusID uint32,
	ip string,
	port uint16,
	isIpv6 bool,
) error {
	return fmt.Errorf("StartClient func is not be implemented")
}

func (net *Net) StartServer(
	headLen HeaderLength,
	busID uint32,
	ip string,
	port uint16,
	threadNum uint8,
	maxClient uint32,
	isIpv6 bool,
) error {
	return fmt.Errorf("StartServer func is not be implemented")
}
func (net *Net) BroadcastMsg(head MsgHead, msgData []byte) error {
	return fmt.Errorf("BroadcastMsg func is not be implemented")
}

func (net *Net) IsWorking() bool {
	return net.working
}

func (net *Net) SetWorking(value bool) {
	net.working = value
}
