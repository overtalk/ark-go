package service

import (
	"github.com/ArkNX/ark-go/base"
	ark "github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/busPlugin/busModule"
	"github.com/ArkNX/ark-go/plugin/busPlugin/msgModule"
	"github.com/ArkNX/ark-go/plugin/logPlugin/logModule"
	"github.com/ArkNX/ark-go/plugin/netPlugin/netServiceManagerModule"
)

type NetClientService struct {
	pluginManager           *ark.PluginManager
	netServiceManagerModule netServiceManagerModule.INetServiceManagerModule
	pNet                    netServiceManagerModule.INet // pNet defined like tcpServer / wsServer
	msgModule               msgModule.IMsgModule
	logModule               logModule.ILogModule
	busModule               busModule.IBusModule

	// Connected connections(may the ConnectState is different)
	realConnections map[uint32]*netServiceManagerModule.ConnectionData
	actorBusMap     map[base.GUID]uint32 // actor id bus id
	accountBusMap   map[string]uint32    // account bus id

	// TODO : add consistent_hashmap_

	tmpConnections         []netServiceManagerModule.ConnectionData
	netMsgCallbacks        map[uint16]netServiceManagerModule.NetMsgFunc
	netEventCallbacks      []netServiceManagerModule.NetEventFunc
	netMsgForwardCallbacks []netServiceManagerModule.NetMsgFunc
}
