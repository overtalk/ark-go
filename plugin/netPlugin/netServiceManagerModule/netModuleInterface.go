package netServiceManagerModule

import (
	"reflect"

	ark "github.com/ArkNX/ark-go/interface"
)

var (
	ModuleName   string
	ModuleUpdate string
	ModuleType   reflect.Type
)

type INetServiceManagerModule interface {
	ark.IModule
}