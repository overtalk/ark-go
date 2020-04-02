package msgModule

import (
	"reflect"

	ark "github.com/ArkNX/ark-go/interface"
)

var (
	ModuleName   string
	ModuleUpdate string
	ModuleType   reflect.Type
)

type IMsgModule interface {
	ark.IModule
}
