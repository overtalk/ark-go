package factoryModule

import (
	"reflect"

	ark "github.com/ArkNX/ark-go/interface"
)

var (
	ModuleName   string
	ModuleUpdate string
	ModuleType   reflect.Type
)

type IFactoryModule interface {
	ark.IModule

	NewClient() (Client, error)
	NewServer() (Server, error)
}

type MsgParser interface {
	Encode(data []byte)
	Decode()
}

// network client
type Client interface {
	Update()

	StartClient(
		dstBusID uint32,
		ip string,
		port uint16,
		isIpv6 bool,
	) error

	Shutdown() error

	// server
	SendMsg(msgData []byte) error

	IsWorking() bool
	SetWorking(value bool)
}

// network server
type Server interface {
	// 定期去处理已经积累的数据
	Update()

	StartServer(
		busID uint32,
		ip string,
		port uint16,
		threadNum uint8,
		maxClient uint32,
		isIpv6 bool,
	) error

	Shutdown() error

	// send message about
	SendMsg(msgData []byte, sessionID int64) error
	BroadcastMsg(msgData []byte) error

	// close one client
	CloseSession(sessionID int64) error

	// status
	IsWorking() bool
	SetWorking(value bool)

	SetMsgDecorator(parser MsgParser)
}
