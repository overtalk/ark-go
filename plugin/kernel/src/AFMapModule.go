package kernelSrc

import (
	"reflect"
	"runtime"

	"github.com/ArkNX/ark-go/interface"
	kernelInterface "github.com/ArkNX/ark-go/plugin/kernel/interface"
)

var (
	MapModuleType   = ark.GetType((*AFCMapModule)(nil))
	MapModuleName   = ark.GetName((*AFCMapModule)(nil))
	MapModuleUpdate = runtime.FuncForPC(reflect.ValueOf((&AFCMapModule{}).Update).Pointer()).Name()
)

func init() {
	kernelInterface.AFIMapModuleName = ark.GetName((*AFCMapModule)(nil))
}

type AFCMapModule struct {
	ark.AFCModule
	// other value
}

func (MapModule *AFCMapModule) Init() error {
	return nil
}
