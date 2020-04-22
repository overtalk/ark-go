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

	NewClient() (ClientService, error)
	NewServer() (ServerService, error)
}

//////////////////////////////
// network server
//////////////////////////////
type ServerService interface {
	Update()
	StartServer(hl uint32, busID uint32, threadNum uint8, maxClient uint32, isIpv6 bool) error
	Shutdown() error
	SendMsg(msgData []byte, sessionID int64) error
	BroadcastMsg(msgData []byte) error
	CloseSession(sessionID int64) error
	IsWorking() bool
	SetWorking(value bool)
}

type Server interface {
	Start(hl uint32, ip string, port uint16, threadNum uint8, maxClient uint32, isIpv6 bool) error
}

//////////////////////////////
// network client
//////////////////////////////
type ClientService interface {
	Update()
	StartClient(dstBusID uint32, ip string, port uint16, isIpv6 bool) error
	Shutdown() error
	SendMsg(msgData []byte) error
	IsWorking() bool
	SetWorking(value bool)
}
