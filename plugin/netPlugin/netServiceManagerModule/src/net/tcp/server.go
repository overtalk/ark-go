package net

import (
	"errors"
	"sync"

	"github.com/ArkNX/ark-go/plugin/netPlugin/netServiceManagerModule"
)

type TcpServer struct {
	netServiceManagerModule.Net

	rwMutex  sync.RWMutex
	sessions map[int64]*netServiceManagerModule.NetSession
	busID    uint32

	netMsgCallback   netServiceManagerModule.NetMsgSessionFunc
	netEventCallback netServiceManagerModule.NetEventFunc

	// TODO: add tcp listener
}

func NewTcpServer(netMsgCallback netServiceManagerModule.NetMsgSessionFunc,
	netEventCallback netServiceManagerModule.NetEventFunc) *TcpServer {

	return &TcpServer{
		Net:              netServiceManagerModule.Net{},
		rwMutex:          sync.RWMutex{},
		sessions:         make(map[int64]*netServiceManagerModule.NetSession),
		busID:            0,
		netMsgCallback:   netMsgCallback,
		netEventCallback: netEventCallback,
	}
}

func (tcpServer *TcpServer) Update() {}

func (tcpServer *TcpServer) StartServer(headLen netServiceManagerModule.HeaderLength,
	busID uint32,
	ip string,
	port uint16,
	threadNum uint8,
	maxClient uint32,
	isIpv6 bool) error {

	// TODO: qinhan
	return errors.New("to finish")
}

func (tcpServer *TcpServer) Shutdown() error {
	// TODO: qinhan
	return errors.New("to finish")
}

func (tcpServer *TcpServer) SendMsg(head netServiceManagerModule.MsgHead, msgData []byte, sessionID int64) error {
	// TODO: qinhan
	return errors.New("to finish")
}

func (tcpServer *TcpServer) BroadcastMsg(head netServiceManagerModule.MsgHead, msgData []byte) error {
	// TODO: qinhan
	return errors.New("to finish")
}

func (tcpServer *TcpServer) CloseSession(sessionID int64) error {
	// TODO: qinhan
	return errors.New("to finish")
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// private func
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (tcpServer *TcpServer) sendMsgToAllClient(msg []byte, msgLen uint32) error {
	// TODO: qinhan
	return errors.New("to finish")
}

func (tcpServer *TcpServer) sendMsg(msg []byte, msgLen uint32, sessionID int64) error {
	// TODO: qinhan
	return errors.New("to finish")
}

// add a net session
func (tcpServer *TcpServer) addNetSession()                {}
func (tcpServer *TcpServer) getNetSession(sessionID int64) {}
func (tcpServer *TcpServer) closeSession()                 {}

func (tcpServer *TcpServer) updateNetSession() {}
func (tcpServer *TcpServer) updateNetEvent()   {}
func (tcpServer *TcpServer) updateNetMsg()     {}

func (tcpServer *TcpServer) closeAllSession() {}
