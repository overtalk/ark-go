package kernelSrc

import (
	"reflect"
	"runtime"

	"github.com/ArkNX/ark-go/interface"
	kernelInterface "github.com/ArkNX/ark-go/plugin/kernel/interface"
)

var (
	ConfigModuleType   = ark.GetType((*AFCConfigModule)(nil))
	ConfigModuleName   = ark.GetName((*AFCConfigModule)(nil))
	ConfigModuleUpdate = runtime.FuncForPC(reflect.ValueOf((&AFCConfigModule{}).Update).Pointer()).Name()
)

func init() {
	kernelInterface.AFIConfigModuleName = ark.GetName((*AFCConfigModule)(nil))
}

type AFCConfigModule struct {
	ark.Module
	// other value
}

func (ConfigModule *AFCConfigModule) Init() error {
	return nil
}
