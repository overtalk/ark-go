package netCommon

type NetEventType uint8

const (
	NONE NetEventType = iota
	CONNECTED
	DISCONNECTED
	RECV_DATA
)

type NetEvent struct {
	type_ NetEventType
	id    int64
	ip    string
	busID uint32
}

func NewNetEvent() *NetEvent { return &NetEvent{} }

// get
func (netEvent *NetEvent) GetType() NetEventType { return netEvent.type_ }
func (netEvent *NetEvent) GetID() int64          { return netEvent.id }
func (netEvent *NetEvent) GetIP() string         { return netEvent.ip }
func (netEvent *NetEvent) GetBusID() uint32      { return netEvent.busID }

// set
func (netEvent *NetEvent) SetType(value NetEventType) { netEvent.type_ = value }
func (netEvent *NetEvent) SetID(value int64)          { netEvent.id = value }
func (netEvent *NetEvent) SetIP(value string)         { netEvent.ip = value }
func (netEvent *NetEvent) SetBusID(value uint32)      { netEvent.busID = value }
