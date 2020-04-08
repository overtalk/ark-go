package service

import (
	"fmt"
	"github.com/ArkNX/ark-go/base"
	ark "github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/busPlugin/msgModule"
	"github.com/ArkNX/ark-go/plugin/logPlugin/logModule"
	"github.com/ArkNX/ark-go/plugin/netPlugin/netServiceManagerModule"
	net "github.com/ArkNX/ark-go/plugin/netPlugin/netServiceManagerModule/src/net/tcp"
	"github.com/ArkNX/ark-go/utils"
)

type eConnectState uint8

const (
	DISCONNECT eConnectState = iota
	CONNECTING
	CONNECTED
	RECONNECT
)

type Connection struct {
	busID          uint32
	sessionID      int64
	lastActiveTime int64
	eConnectState  eConnectState
}

func NewConnection(busID uint32, sessionID int64, lastActiveTime int64, status eConnectState) *Connection {
	return &Connection{
		busID:          busID,
		sessionID:      sessionID,
		lastActiveTime: lastActiveTime,
		eConnectState:  status,
	}
}

//////////////////////////////////////////////////////
// NetServerService
// this defines the server service, which contains the connections & callbacks
//////////////////////////////////////////////////////
type NetServerService struct {
	pluginManager           *ark.PluginManager
	netServiceManagerModule netServiceManagerModule.INetServiceManagerModule
	msgModule               msgModule.IMsgModule
	logModule               logModule.ILogModule
	pNet                    netServiceManagerModule.INet // pNet defined like tcpServer / wsServer

	regServerCallback netServiceManagerModule.NetMsgSessionFunc

	netMsgCallbacks        map[uint16]netServiceManagerModule.NetMsgFunc
	netForwardMsgCallbacks []netServiceManagerModule.NetMsgFunc
	netEventCallbacks      []netServiceManagerModule.NetEventFunc

	connectionList map[uint32]Connection
}

func NewNetServerService() *NetServerService {
	pluginManager := ark.GetPluginManagerInstance()
	return &NetServerService{
		pluginManager:           pluginManager,
		netServiceManagerModule: pluginManager.FindModule(netServiceManagerModule.ModuleName).(netServiceManagerModule.INetServiceManagerModule),
		msgModule:               pluginManager.FindModule(msgModule.ModuleName).(msgModule.IMsgModule),
		logModule:               pluginManager.FindModule(logModule.ModuleName).(logModule.ILogModule),
		pNet:                    nil,
		regServerCallback:       nil,
		netMsgCallbacks:         make(map[uint16]netServiceManagerModule.NetMsgFunc),
		netForwardMsgCallbacks:  make([]netServiceManagerModule.NetMsgFunc, 0),
		netEventCallbacks:       make([]netServiceManagerModule.NetEventFunc, 0),
		connectionList:          make(map[uint32]Connection),
	}
}

// start server
func (netServer *NetServerService) Start(len netServiceManagerModule.HeaderLength, busID uint32, ep base.Endpoint,
	threadCount uint8, maxConnection uint32) error {
	switch ep.Proto() {
	case base.ProtoTypeTcp:
		netServer.pNet = net.NewTcpServer(netServer.OnNetMsg, netServer.OnNetEvent)
		return netServer.pNet.StartServer(len, busID, ep.GetIP(), ep.GetPort(), threadCount, maxConnection, ep.IsV6())
	case base.ProtoTypeUdp:
		// TODO
	case base.ProtoTypeWs:
		// TODO
	default:
		return fmt.Errorf("invalid endpoint proto type : %v", ep.Proto())
	}

	return nil
}

func (netServer *NetServerService) Update() error {
	if netServer.pNet == nil {
		return fmt.Errorf("pNet is nil")
	}

	netServer.pNet.Update()

	netServer.ProcessConnection()
	return nil
}

func (netServer *NetServerService) GetNet() netServiceManagerModule.INet {
	return netServer.pNet
}

func (netServer *NetServerService) RegMsgCallback(msgID uint16, cb netServiceManagerModule.NetMsgFunc) error {
	if _, isExist := netServer.netMsgCallbacks[msgID]; !isExist {
		return fmt.Errorf("multiple registration for message id : %d", msgID)
	}

	netServer.netMsgCallbacks[msgID] = cb
	return nil
}

func (netServer *NetServerService) RegForwardMsgCallback(cb netServiceManagerModule.NetMsgFunc) {
	netServer.netForwardMsgCallbacks = append(netServer.netForwardMsgCallbacks, cb)
}

func (netServer *NetServerService) RegNetEventCallback(cb netServiceManagerModule.NetEventFunc) {
	netServer.netEventCallbacks = append(netServer.netEventCallbacks, cb)
}

func (netServer *NetServerService) RegRegServerCallBack(cb netServiceManagerModule.NetMsgSessionFunc) {
	netServer.regServerCallback = cb
}

func (netServer *NetServerService) OnNetMsg(msg *netServiceManagerModule.NetMsg, sessionID int64) {
	msgID := msg.GetMsgID()
	switch msgID {
	// TODO: add some default message type
	default:
		f, isExist := netServer.netMsgCallbacks[msgID]
		if !isExist {
			// TODO: add log
			// TODO: remove connections in m_pNetServiceManagerModule
		}
		f(msg)
	}
}

func (netServer *NetServerService) OnNetEvent(event *netServiceManagerModule.NetEvent) {
	switch event.GetType() {
	//case netCommon.NONE:
	//case netCommon.RECV_DATA:
	case netServiceManagerModule.NetEventConnected:
		// TODO: add log
	case netServiceManagerModule.NetEventDisconnected:
		// TODO: add log
	}

	for _, callback := range netServer.netEventCallbacks {
		callback(event)
	}
}

func (netServer *NetServerService) ProcessConnection() {
	now := utils.GetNowTime()
	for index, connection := range netServer.connectionList {
		if connection.eConnectState != CONNECTED {
			continue
		} else if connection.lastActiveTime+30*1000 < now {
			connection.eConnectState = DISCONNECT

			// close session
			delete(netServer.connectionList, index)
			netServer.pNet.CloseSession(connection.sessionID)

			event := &netServiceManagerModule.NetEvent{}
			event.SetType(netServiceManagerModule.NetEventDisconnected)
			event.SetBusID(connection.busID)
			netServer.OnNetEvent(event)
			break
		}
	}
}

func (netServer *NetServerService) AddConnection(busID uint32, sessionID int64) {
	if _, isExist := netServer.connectionList[busID]; isExist {
		return
	}

	netServer.connectionList[busID] = Connection{
		busID:          busID,
		sessionID:      sessionID,
		lastActiveTime: utils.GetNowTime(),
		eConnectState:  CONNECTED,
	}
}

func (netServer *NetServerService) UpdateConnection(busID uint32) {
	connection, isExist := netServer.connectionList[busID]
	if !isExist {
		return
	}

	connection.lastActiveTime = utils.GetNowTime()
}

func (netServer *NetServerService) RemoveConnection(busID uint32) {
	delete(netServer.connectionList, busID)
}
