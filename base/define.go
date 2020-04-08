package base

type (
	AppType   uint8
	GUID      int64
	BusID     uint32
	SessionID int64
	ProtoType string
)

const (
	APPDefault AppType = iota // none
	APPMax     AppType = 255  // max of all processes
)
