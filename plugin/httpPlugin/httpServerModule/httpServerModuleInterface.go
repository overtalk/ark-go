package httpServerModule

import (
	ark "github.com/ArkNX/ark-go/interface"
	"reflect"
)

var (
	ModuleName   string
	ModuleType   reflect.Type
	ModuleUpdate string
)

type IHttpServerModule interface {
	ark.IModule
	Start(port uint16) error
}
