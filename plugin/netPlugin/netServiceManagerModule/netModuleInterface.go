package netServiceManagerModule

import (
	"github.com/ArkNX/ark-go/base"
	netCommon "github.com/ArkNX/ark-go/plugin/netPlugin/common"
	"reflect"

	ark "github.com/ArkNX/ark-go/interface"
)

var (
	ModuleName   string
	ModuleUpdate string
	ModuleType   reflect.Type
)

type INetServiceManagerModule interface {
	ark.IModule
	CreateServer()
	GetSelfNetServer()
}

//////////////////////////////////////////////////////
//////////////////////////////////////////////////////
//////////////////////////////////////////////////////
//////////////////////////////////////////////////////
type NetMsgFunc func(*netCommon.NetMsg)
type NetMsgSessionFunc func(*netCommon.NetMsg, base.GUID)
type NetEventFunc func(*netCommon.NetEvent)

type INet interface {
	Update() error
	StartClient() error
	CloseSession(sessionID int64) error
}

type ConnectionData struct {
}

type ClientService interface {
	StartClient(headLen netCommon.HeaderLength, targetBusID uint32, endpoint base.Endpoint) error
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
	ProcessUpdate()
	ProcessAddConnection()
	CreateNet(proto base.ProtoType) INet
	OnConnect()
	OnDisonnect()
	OnNetMsg()
	OnNetEvent()
	KeepReport()
	LogServerInfo()
	KeepAlive()
	AddServerNode()
	GetServerNode()
	RemoveServerNode()
	RegToServer()
	SendReport()
}

type ServerService interface {
	GetNet() INet
	Start(netCommon.HeaderLength, uint32, base.Endpoint, uint8, uint32) error
	// register callback
	// those func is used for net
	RegMsgCallback(msgID uint16, cb NetMsgFunc) error
	RegForwardMsgCallback(cb NetMsgFunc)
	RegNetEventCallback(cb NetEventFunc)
	RegRegServerCallBack(cb NetMsgSessionFunc)
	// others
	OnNetMsg(msg *netCommon.NetMsg, sessionID base.GUID)
	OnNetEvent(event *netCommon.NetEvent)
	ProcessConnection()
	AddConnection(busID uint32, sessionID int64)
	UpdateConnection(busID uint32)
	RemoveConnection(busID uint32)
}
