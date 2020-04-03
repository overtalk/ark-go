package netCommon

import (
	"bytes"
	"github.com/ArkNX/ark-go/base"
	"github.com/ArkNX/ark-go/util"
)

type NetSession struct {
	headLen   uint32
	sessionID int64
	objectID  base.GUID
	buffer    bytes.Buffer

	msgQueue   util.LockFreeQueue
	eventQueue util.LockFreeQueue
	session    interface{} // tcp session / http session

	connected  bool
	needRemove bool
}

func (netSession *NetSession) GetSession() interface{} { return netSession.session }

func (netSession *NetSession) AddBuffer(data []byte) int {
	cnt, _ := netSession.buffer.Write(data)
	return cnt
}

func (netSession *NetSession) GetBuffer(b []byte) (int, error) { return netSession.buffer.Read(b) }

func (netSession *NetSession) GetBufferLen() int { return netSession.buffer.Len() }

func (netSession *NetSession) GetHeadLen() uint32 { return netSession.headLen }

func (netSession *NetSession) GetSessionID() int64 { return netSession.sessionID }

func (netSession *NetSession) SetSessionID(value int64) { netSession.sessionID = value }

func (netSession *NetSession) NeedRemove() bool { return netSession.needRemove }

func (netSession *NetSession) SetNeedRemove(value bool) { netSession.needRemove = value }

func (netSession *NetSession) AddNetEvent(event *NetEvent) { netSession.eventQueue.Put(event) }

func (netSession *NetSession) PopNetEvent() (*NetEvent, bool) {
	e, flag := netSession.eventQueue.Get()
	if !flag {
		return nil, false
	}
	return e.(*NetEvent), true
}

func (netSession *NetSession) AddNetMsg(msg *NetMsg) { netSession.msgQueue.Put(msg) }

func (netSession *NetSession) PopNetMsg() (*NetMsg, bool) {
	e, flag := netSession.msgQueue.Get()
	if !flag {
		return nil, false
	}
	return e.(*NetMsg), true
}

func (netSession *NetSession) ParseBufferToMsg() {
	header := make([]byte, netSession.headLen)
	n, err := netSession.GetBuffer(header)
	if uint32(n) != netSession.headLen || err != nil {
		return
	}

	// TODO : get message from buffer to queue
}
