package server

import (
	"encoding/binary"
	"errors"
)

////////////////////////////////////////////////////////
// header about
////////////////////////////////////////////////////////
type HeadLength uint32

const (
	CSHeadLength HeadLength = 6  // cs head
	SSHeadLength HeadLength = 22 // ss head
)

/*
| msg id | msg len | actor id | src bus | dst bus |
|    2   |    4    |     8    |    4    |    4    | = 22
*/
type MsgHead struct {
	id       uint16 // msg id
	length   uint32 // msg length (without header length)
	actorID  uint64
	srcBusID uint32
	dstBusID uint32
}

////////////////////////////////////////////////////////
// net message
////////////////////////////////////////////////////////
type NetMsg struct {
	head    MsgHead
	msgData []byte
}

func NewNetMsgFromData(data []byte) *NetMsg {
	netMsg := &NetMsg{
		head: MsgHead{
			length: uint32(len(data)),
		},
		msgData: make([]byte, len(data)),
	}

	// copy
	copy(netMsg.msgData, data)
	return netMsg
}

func NewNetMsgFromMetMsg(msg *NetMsg) *NetMsg {
	netMsg := NewNetMsgFromData(msg.msgData)
	netMsg.head = msg.head
	return netMsg
}

// get
func (netMsg *NetMsg) GetHead() *MsgHead    { return &netMsg.head }
func (netMsg *NetMsg) GetMsgID() uint16     { return netMsg.head.id }
func (netMsg *NetMsg) GetMsgLength() uint32 { return netMsg.head.length }
func (netMsg *NetMsg) GetActorID() uint64   { return netMsg.head.actorID }
func (netMsg *NetMsg) GetSrcBusID() uint32  { return netMsg.head.srcBusID }
func (netMsg *NetMsg) GetDstBusID() uint32  { return netMsg.head.dstBusID }

// set
func (netMsg *NetMsg) SetMsgID(value uint16)     { netMsg.head.id = value }
func (netMsg *NetMsg) SetMsgLength(value uint32) { netMsg.head.length = value }
func (netMsg *NetMsg) SetActorID(value uint64)   { netMsg.head.actorID = value }
func (netMsg *NetMsg) SetSrcBusID(value uint32)  { netMsg.head.srcBusID = value }
func (netMsg *NetMsg) SetDstBusID(value uint32)  { netMsg.head.dstBusID = value }

////////////////////////////////////////////////////////
// deserialize net message
////////////////////////////////////////////////////////
func DeserializeMsgHead(l HeadLength, data []byte) (*MsgHead, error) {
	if len(data) != int(l) {
		return nil, errors.New("invalid header length")
	}

	header := &MsgHead{}
	header.id = binary.BigEndian.Uint16(data[:2])
	header.length = binary.BigEndian.Uint32(data[2:CSHeadLength])

	if l == SSHeadLength {
		header.actorID = binary.BigEndian.Uint64(data[CSHeadLength : CSHeadLength+8])
		header.srcBusID = binary.BigEndian.Uint32(data[CSHeadLength+8 : CSHeadLength+12])
		header.dstBusID = binary.BigEndian.Uint32(data[CSHeadLength+12:])
	}

	return header, nil
}
