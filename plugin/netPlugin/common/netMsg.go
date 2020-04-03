package netCommon

////////////////////////////////////////////////////////
// header about
////////////////////////////////////////////////////////
type HeaderLength uint32

const (
	CS_HEAD_LENGTH HeaderLength = 6  // cs head
	SS_HEAD_LENGTH HeaderLength = 22 // ss head
)

type MsgHead struct {
	id     uint16 // msg id
	length uint32 // msg length (without header length)
}

/*
| msg id | msg len |
|    2   |    4    | = 6
*/
type CSMsgHead struct {
	MsgHead
}

/*
| msg id | msg len | actor id | src bus | dst bus |
|    2   |    4    |     8    |    4    |    4    | = 22
*/
type SSMsgHead struct {
	MsgHead
	actorID  int64
	srcBusID uint32
	dstBusID uint32
}

////////////////////////////////////////////////////////
// net message
////////////////////////////////////////////////////////
type NetMsg struct {
	head    SSMsgHead
	msgData []byte
}

func NewNetMsgFromData(data []byte) *NetMsg {
	netMsg := &NetMsg{
		head: SSMsgHead{
			MsgHead: MsgHead{
				length: uint32(len(data)),
			},
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
func (netMsg *NetMsg) GetHead() *SSMsgHead  { return &netMsg.head }
func (netMsg *NetMsg) GetMsgID() uint16     { return netMsg.head.id }
func (netMsg *NetMsg) GetMsgLength() uint32 { return netMsg.head.length }
func (netMsg *NetMsg) GetActorID() int64    { return netMsg.head.actorID }
func (netMsg *NetMsg) GetSrcBusID() uint32  { return netMsg.head.srcBusID }
func (netMsg *NetMsg) GetDstBusID() uint32  { return netMsg.head.dstBusID }

// set
func (netMsg *NetMsg) SetMsgID(value uint16)     { netMsg.head.id = value }
func (netMsg *NetMsg) SetMsgLength(value uint32) { netMsg.head.length = value }
func (netMsg *NetMsg) SetActorID(value int64)    { netMsg.head.actorID = value }
func (netMsg *NetMsg) SetSrcBusID(value uint32)  { netMsg.head.srcBusID = value }
func (netMsg *NetMsg) SetDstBusID(value uint32)  { netMsg.head.dstBusID = value }
