package src

import (
	"log"
	"path/filepath"
	"reflect"
	"runtime"

	"github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/netPlugin/msgModule"
)

func init() {
	t := reflect.TypeOf((*CMsgModule)(nil))
	if !t.Implements(reflect.TypeOf((*msgModule.IMsgModule)(nil)).Elem()) {
		log.Fatal("IMsgModule is not implemented by CMsgModule")
	}

	msgModule.ModuleType = t.Elem()
	msgModule.ModuleName = filepath.Join(msgModule.ModuleType.PkgPath(), msgModule.ModuleType.Name())
	msgModule.ModuleUpdate = runtime.FuncForPC(reflect.ValueOf((&CMsgModule{}).Update).Pointer()).Name()
}

type CMsgModule struct {
	ark.Module
	// other data
}

func (msgModule *CMsgModule) Init() error {
	return nil
}