package src

import (
	"log"
	"path/filepath"
	"reflect"
	"runtime"

	"github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/netPlugin/factoryModule"
)

func init() {
	t := reflect.TypeOf((*CFactoryModule)(nil))
	if !t.Implements(reflect.TypeOf((*factoryModule.IFactoryModule)(nil)).Elem()) {
		log.Fatal("IFactoryModule is not implemented by CFactoryModule")
	}

	factoryModule.ModuleType = t.Elem()
	factoryModule.ModuleName = filepath.Join(factoryModule.ModuleType.PkgPath(), factoryModule.ModuleType.Name())
	factoryModule.ModuleUpdate = runtime.FuncForPC(reflect.ValueOf((&CFactoryModule{}).Update).Pointer()).Name()
}

type CFactoryModule struct {
	ark.Module
	// other data
}

func (factoryModule *CFactoryModule) Init() error {
	return nil
}