package netServiceManagerModule

import (
	"errors"
	"github.com/ArkNX/ark-go/base"
	"github.com/ArkNX/ark-go/utils/ringBuffer"
	"github.com/ArkNX/ark-go/utils/ringQueue"
)

type NetSession struct {
	headLen   uint32
	sessionID int64
	objectID  base.GUID
	buffer    ringbuffer.RingBuffer

	msgQueue   ringQueue.RingQueue
	eventQueue ringQueue.RingQueue
	session    interface{} // tcp session / http session

	connected  bool
	needRemove bool
}

func (netSession *NetSession) GetSession() interface{} { return netSession.session }

func (netSession *NetSession) AddBuffer(data []byte) int {
	cnt, _ := netSession.buffer.Write(data)
	return cnt
}

func (netSession *NetSession) GetBuffer(n int) ([]byte, error) {
	header, tail := netSession.buffer.LazyRead(n)

	if len(header)+len(tail) != n {
		return nil, errors.New("unmatched bytes length")
	}

	if len(tail) != 0 {
		header = append(header, tail...)
	}

	return header, nil
}

func (netSession *NetSession) GetBufferLen() int { return netSession.buffer.Len() }

func (netSession *NetSession) GetHeadLen() uint32 { return netSession.headLen }

func (netSession *NetSession) GetSessionID() int64 { return netSession.sessionID }

func (netSession *NetSession) SetSessionID(value int64) { netSession.sessionID = value }

func (netSession *NetSession) NeedRemove() bool { return netSession.needRemove }

func (netSession *NetSession) SetNeedRemove(value bool) { netSession.needRemove = value }

func (netSession *NetSession) AddNetEvent(event *NetEvent) {
	netSession.eventQueue.PushOne(event)
}

func (netSession *NetSession) PopNetEvent() (*NetEvent, bool) {
	e, err := netSession.eventQueue.PopOne()
	if err != nil {
		return nil, false
	}
	return e.(*NetEvent), true
}

func (netSession *NetSession) AddNetMsg(msg *NetMsg) { netSession.msgQueue.PushOne(msg) }

func (netSession *NetSession) PopNetMsg() (*NetMsg, bool) {
	e, err := netSession.msgQueue.PopOne()
	if err != nil {
		return nil, false
	}
	return e.(*NetMsg), true
}

// ParseBufferToMsg
func (netSession *NetSession) ParseBufferToMsg() {
	for {
		msg, err := netSession.getNetMsg()
		if err != nil {
			break
		}

		netSession.AddNetMsg(msg)
	}
}

// getNetMsg defines the func to read msg from queue
func (netSession *NetSession) getNetMsg() (*NetMsg, error) {
	headerBytes, err := netSession.GetBuffer(int(netSession.headLen))
	if err != nil {
		return nil, err
	}

	header, err := DeserializationMsgHead(headerBytes)
	if err != nil {
		return nil, err
	}

	msg, err := netSession.GetBuffer(int(header.length))
	if err != nil {
		return nil, err
	}

	netSession.buffer.Shift(int(netSession.headLen + header.length))

	return &NetMsg{
		head: SSMsgHead{
			MsgHead:  *header,
			actorID:  0,
			srcBusID: 0,
			dstBusID: 0,
		},
		msgData: msg,
	}, nil
}
