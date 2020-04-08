package src

import (
	"log"
	"path/filepath"
	"reflect"
	"runtime"

	"github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/netPlugin/netServiceManagerModule"
)

func init() {
	t := reflect.TypeOf((*CNetServiceManagerModule)(nil))
	if !t.Implements(reflect.TypeOf((*netServiceManagerModule.INetServiceManagerModule)(nil)).Elem()) {
		log.Fatal("INetServiceManagerModule is not implemented by CNetServiceManagerModule")
	}

	netServiceManagerModule.ModuleType = t.Elem()
	netServiceManagerModule.ModuleName = filepath.Join(netServiceManagerModule.ModuleType.PkgPath(), netServiceManagerModule.ModuleType.Name())
	netServiceManagerModule.ModuleUpdate = runtime.FuncForPC(reflect.ValueOf((&CNetServiceManagerModule{}).Update).Pointer()).Name()
}

type CNetServiceManagerModule struct {
	ark.Module
	// other data
}

func (netServiceManagerModule *CNetServiceManagerModule) Init() error {
	return nil
}
