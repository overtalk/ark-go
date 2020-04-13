package server

import (
	"errors"
	"fmt"
	"github.com/ArkNX/ark-go/utils/ringQueue"

	"github.com/ArkNX/ark-go/plugin/netPlugin/factoryModule"
)

type SessionManager struct {
	working         bool
	sessions        map[int64]*NetSession
	connectQueue    *ringQueue.RingQueue
	disconnectQueue *ringQueue.RingQueue
	parser          factoryModule.MsgParser
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		working:         false,
		sessions:        make(map[int64]*NetSession),
		connectQueue:    ringQueue.New(512),
		disconnectQueue: ringQueue.New(512),
	}
}

func (sm *SessionManager) StartServer(
	busID uint32,
	ip string,
	port uint16,
	threadNum uint8,
	maxClient uint32,
	isIpv6 bool,
) error {
	if sm.parser == nil {
		return errors.New("msg decoder is nil")
	}

	return errors.New("unrealized")
}

func (sm *SessionManager) Update() {
	for k, v := range sm.sessions {
		// TODO: handle msg
		fmt.Println("update for ", k, v)
	}
}

func (sm *SessionManager) Shutdown() error {
	return errors.New("unrealized")
}

func (sm *SessionManager) SendMsg(msgData []byte, sessionID int64) error {
	return sm.sessions[sessionID].GetConn().AsyncWrite(msgData)
}

func (sm *SessionManager) BroadcastMsg(msgData []byte) error {
	for _, session := range sm.sessions {
		session.GetConn().AsyncWrite(msgData)
	}
	return nil
}

func (sm *SessionManager) CloseSession(sessionID int64) error {
	delete(sm.sessions, sessionID)
	return sm.sessions[sessionID].GetConn().Close()
}

func (sm *SessionManager) IsWorking() bool { return sm.working }

func (sm *SessionManager) SetWorking(value bool) { sm.working = value }

func (sm *SessionManager) SetMsgDecorator(parser factoryModule.MsgParser) {
	sm.parser = parser
}

func (sm *SessionManager) AddSession(session *NetSession) {
	sm.connectQueue.PushOne(session)
}

func (sm *SessionManager) RemoveSession(sessionID int64) {
	sm.disconnectQueue.PushOne(sessionID)
}

func (sm *SessionManager) AddBuffer(id int64, msgData []byte) {
	sm.sessions[id].AddBuffer(msgData)
}
