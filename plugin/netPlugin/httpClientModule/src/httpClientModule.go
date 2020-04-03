package src

import (
	"log"
	"path/filepath"
	"reflect"
	"runtime"

	"github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/netPlugin/httpClientModule"
)

func init() {
	t := reflect.TypeOf((*CHttpClientModule)(nil))
	if !t.Implements(reflect.TypeOf((*httpClientModule.IHttpClientModule)(nil)).Elem()) {
		log.Fatal("IHttpClientModule is not implemented by CHttpClientModule")
	}

	httpClientModule.ModuleType = t.Elem()
	httpClientModule.ModuleName = filepath.Join(httpClientModule.ModuleType.PkgPath(), httpClientModule.ModuleType.Name())
	httpClientModule.ModuleUpdate = runtime.FuncForPC(reflect.ValueOf((&CHttpClientModule{}).Update).Pointer()).Name()
}

type CHttpClientModule struct {
	ark.Module
	// other data
}

func (httpClientModule *CHttpClientModule) Init() error {
	return nil
}