package kernelSrc

import (
	"reflect"
	"runtime"

	"github.com/ArkNX/ark-go/interface"
	kernelInterface "github.com/ArkNX/ark-go/plugin/kernel/interface"
)

var (
	MetaClassModuleType   = ark.GetType((*AFCMetaClassModule)(nil))
	MetaClassModuleName   = ark.GetName((*AFCMetaClassModule)(nil))
	MetaClassModuleUpdate = runtime.FuncForPC(reflect.ValueOf((&AFCMetaClassModule{}).Update).Pointer()).Name()
)

func init() {
	kernelInterface.AFIMetaClassModuleName = ark.GetName((*AFCMetaClassModule)(nil))
}

type AFCMetaClassModule struct {
	ark.Module
	// other value
}

func (MetaClassModule *AFCMetaClassModule) Init() error {
	return nil
}
