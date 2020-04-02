package src

import (
	"log"
	"path/filepath"
	"reflect"
	"runtime"

	"github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/busPlugin/msgModule"
)

func init() {
	t := reflect.TypeOf((*AFCMsgModule)(nil))
	if !t.Implements(reflect.TypeOf((*msgModule.AFIMsgModule)(nil)).Elem()) {
		log.Fatal("AFIMsgModule is not implemented by AFCMsgModule")
	}

	msgModule.ModuleType = t.Elem()
	msgModule.ModuleName = filepath.Join(msgModule.ModuleType.PkgPath(), msgModule.ModuleType.Name())
	msgModule.ModuleUpdate = runtime.FuncForPC(reflect.ValueOf((&AFCMsgModule{}).Update).Pointer()).Name()
}

type AFCMsgModule struct {
	ark.AFCModule
	// other data
}

func (msgModule *AFCMsgModule) Init() error {
	return nil
}
