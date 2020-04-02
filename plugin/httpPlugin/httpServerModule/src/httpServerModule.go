package src

import (
	"fmt"
	"github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/httpPlugin/httpServerModule"
	"github.com/ArkNX/ark-go/plugin/logPlugin/logModule"
	"log"
	"net/http"
	"path/filepath"
	"reflect"
	"runtime"
)

func init() {
	t := reflect.TypeOf((*CHttpServerModule)(nil))
	if !t.Implements(reflect.TypeOf((*httpServerModule.IHttpServerModule)(nil)).Elem()) {
		log.Fatal("IHttpServerModule is not implemented by CHttpServerModule")
	}

	httpServerModule.ModuleType = t.Elem()
	httpServerModule.ModuleName = filepath.Join(httpServerModule.ModuleType.PkgPath(), httpServerModule.ModuleType.Name())
	httpServerModule.ModuleUpdate = runtime.FuncForPC(reflect.ValueOf((&CHttpServerModule{}).Update).Pointer()).Name()
}

type CHttpServerModule struct {
	ark.Module
	// other data
	log logModule.ILogModule
}

func (httpServerModule *CHttpServerModule) Init() error {
	m := httpServerModule.GetPluginManager().FindModule(logModule.ModuleName)
	logModule, ok := m.(logModule.ILogModule)
	if !ok {
		log.Fatal("failed to get log module in httpServer module")
	}
	httpServerModule.log = logModule
	return nil
}

func (httpServerModule *CHttpServerModule) PostInit() error {
	go httpServerModule.Start(9999)
	return nil
}

func (httpServerModule *CHttpServerModule) Start(port uint16) error {
	http.HandleFunc("/hello", HelloServer)
	httpServerModule.log.GetLogger().Warn("start http server on port : ", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

//func (httpServerModule *CHttpServerModule) Update() error {
//	httpServerModule.log.GetLogger().WithField("test-key", "test-value").Warn("Update func in httpServerModule")
//	return nil
//}
