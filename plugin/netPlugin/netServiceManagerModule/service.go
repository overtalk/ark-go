package netServiceManagerModule

import (
	"github.com/ArkNX/ark-go/base"
)

type NetMsgFunc func(*NetMsg)
type NetMsgSessionFunc func(msg *NetMsg, sessionID int64)
type NetEventFunc func(*NetEvent)

//////////////////////////////////
// client service
//////////////////////////////////
type ConnectState uint8

const (
	ConnectStateDisconnect ConnectState = iota
	ConnectStateConnecting
	ConnectStateConnected
	ConnectStateReconnect
)

type ConnectionData struct {
	HeadLen        HeaderLength
	ServerBusID    uint32
	EndPoint       base.Endpoint
	NetClient      INet
	ConnectStatus  ConnectState
	LastActiveTime int64
}

// ClientService is to manager client net
// the implement should contain a INet(client)
type ClientService interface {
	StartClient(headLen HeaderLength, targetBusID uint32, endpoint base.Endpoint) error
	Update()

	GetConnectionInfo(busID uint32) *ConnectionData
	GetSuitableConnect(key string) *ConnectionData

	RegMsgCallback(msgID uint16, cb NetMsgFunc) error
	RegForwardMsgCallback(cb NetMsgFunc) error
	RegNetEventCallback(cb NetEventFunc) error

	AddAccountBusID(account string, busID uint32)
	RemoveAccountBusID(account string)
	GetAccountBusID(account string) int

	AddActorBusID(actor base.GUID, busID uint32)
	RemoveActorBusID(actor base.GUID)
	GetActorBusID(actor base.GUID) int

	// others
	//ProcessUpdate()
	//ProcessAddConnection()
	//CreateNet(proto base.ProtoType) INet
	//OnConnect()
	//OnDisconnect()
	//OnNetMsg()
	//OnNetEvent()
	//KeepReport()
	//LogServerInfo()
	//KeepAlive()
	//AddServerNode()
	//GetServerNode()
	//RemoveServerNode()
	//RegToServer()
	//SendReport()
}

//////////////////////////////////
// server service
//////////////////////////////////

// ServerService is to manager server net
// the implement should contain a INet(server)
type ServerService interface {
	GetNet() INet
	Start(HeaderLength, uint32, base.Endpoint, uint8, uint32) error
	// register callback
	// those func is used for net
	RegMsgCallback(msgID uint16, cb NetMsgFunc) error
	RegForwardMsgCallback(cb NetMsgFunc)
	RegNetEventCallback(cb NetEventFunc)
	RegRegServerCallBack(cb NetMsgSessionFunc)
	// others
	//OnNetMsg(msg *NetMsg, sessionID base.GUID)
	//OnNetEvent(event *NetEvent)
	//ProcessConnection()
	//AddConnection(busID uint32, sessionID int64)
	//UpdateConnection(busID uint32)
	//RemoveConnection(busID uint32)
}
