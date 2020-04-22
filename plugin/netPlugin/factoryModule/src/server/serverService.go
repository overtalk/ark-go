package server

import (
	"errors"
	"fmt"
	"github.com/ArkNX/ark-go/base"
	"github.com/ArkNX/ark-go/plugin/netPlugin/factoryModule"
	"github.com/ArkNX/ark-go/utils/ringQueue"
)

type NetMsgHandler func(msg *base.NetMsg, sessionID int64)

type ServerService struct {
	working      bool
	ep           *base.Endpoint
	handler      NetMsgHandler
	sessions     map[int64]*base.NetSession
	connectQueue *ringQueue.RingQueue
	server       factoryModule.Server
}

func NewServerService(handler NetMsgHandler, ep *base.Endpoint) (*ServerService, error) {
	ret := &ServerService{
		working:      false,
		sessions:     make(map[int64]*base.NetSession),
		connectQueue: ringQueue.New(128),
		handler:      handler,
		ep:           ep,
	}

	switch ep.Proto() {
	case base.ProtoTypeTcp:
		ret.server = NewGNetServer(ret)
	default:
		return nil, fmt.Errorf("invalid proto : %s", ep.Proto())
	}
	return ret, nil
}

func (ss *ServerService) StartServer(
	hl uint32,
	busID uint32,
	threadNum uint8,
	maxClient uint32,
	isIpv6 bool,
) error {
	ss.working = true
	return ss.server.Start(hl, ss.ep.GetIP(), ss.ep.GetPort(), threadNum, maxClient, isIpv6)
}

func (ss *ServerService) Update() {
	// accept new sessions
	n := ss.connectQueue.GetBufferLength()
	connectSession := make([]interface{}, n)
	ss.connectQueue.Pop(connectSession)
	for _, session := range connectSession {
		// TODO : gen a session id
		var sessionID int64 = 12
		ss.sessions[sessionID] = session.(*base.NetSession)
	}

	var needRemove []int64
	for sessionID, session := range ss.sessions {
		if session.NeedRemove() {
			needRemove = append(needRemove, sessionID)
			continue
		}
		// TODO: handle msg
		session.ParseBufferToMsg()
		for {
			msg, flag := session.PopNetMsg()
			if !flag {
				break
			}
			// TODO: handle msg
			fmt.Println("handle message :", msg)
		}
	}

	// remove session
	for _, sessionID := range needRemove {
		delete(ss.sessions, sessionID)
	}
}

func (ss *ServerService) Shutdown() error {
	return errors.New("unrealized")
}

func (ss *ServerService) SendMsg(msgData []byte, sessionID int64) error {
	return ss.sessions[sessionID].GetConn().AsyncWrite(msgData)
}

func (ss *ServerService) BroadcastMsg(msgData []byte) error {
	for _, session := range ss.sessions {
		session.GetConn().AsyncWrite(msgData)
	}
	return nil
}

func (ss *ServerService) CloseSession(sessionID int64) error {
	delete(ss.sessions, sessionID)
	return ss.sessions[sessionID].GetConn().Close()
}

func (ss *ServerService) IsWorking() bool { return ss.working }

func (ss *ServerService) SetWorking(value bool) { ss.working = value }

func (ss *ServerService) AddSession(session *base.NetSession) {
	ss.connectQueue.PushOne(session)
}

func (ss *ServerService) RemoveSession(sessionID int64) {
	ss.sessions[sessionID].SetNeedRemove(true)
}

func (ss *ServerService) AddBuffer(id int64, msgData []byte) {
	ss.sessions[id].AddBuffer(msgData)
}
