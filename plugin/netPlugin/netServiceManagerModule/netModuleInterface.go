package netServiceManagerModule

import (
	"reflect"

	"github.com/ArkNX/ark-go/base"
	ark "github.com/ArkNX/ark-go/interface"
)

var (
	ModuleName   string
	ModuleUpdate string
	ModuleType   reflect.Type
)

type INetServiceManagerModule interface {
	ark.IModule

	// server service about
	CreateServer(headLen HeaderLength)
	GetSelfNetServer() ServerService

	AddNetConnectionBus(clientBusID uint32, netPtr INet) error
	RemoveNetConnectionBus(clientBusID uint32) error
	GetNetConnectionBus(srcBusID uint32, targetBusID uint32) INet
	GetClientService(appType base.AppType) ClientService

	GetSessionID(clientBusID uint32) int64
	RemoveSessionID(sessionID int64) error

	// client service about
	CreateClientService(busAddr base.BusAddr, ip string, port uint16) error
}
