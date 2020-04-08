package src

import (
	"github.com/ArkNX/ark-go/plugin/busPlugin/busModule"
	"github.com/ArkNX/ark-go/plugin/consulPlugin/consulModule"
	"github.com/ArkNX/ark-go/plugin/logPlugin/logModule"
	"github.com/ArkNX/ark-go/plugin/netPlugin/netServiceManagerModule/src/service"
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

type Temp struct {
	selfBusID   uint32
	ClientBusID uint32
}

type CNetServiceManagerModule struct {
	ark.Module
	// other modules
	busModule    busModule.IBusModule
	consulModule consulModule.IConsulModule
	logModule    logModule.ILogModule

	// other data
	serverServices  map[uint32]netServiceManagerModule.ServerService
	clientServices  map[uint32]netServiceManagerModule.ClientService
	netBusRelations map[Temp]netServiceManagerModule.INet
	busSessionList  map[uint32]int64
}

func (module *CNetServiceManagerModule) Init() error {
	return nil
}

func (module *CNetServiceManagerModule) CreateServer(headLen netServiceManagerModule.HeaderLength) {
	selfProc := module.busModule.GetSelfProc()
	pServer := service.NewNetServerService()

	if err := pServer.Start(headLen, selfProc.BusId, selfProc.ServerEp, selfProc.ThreadNum, selfProc.MaxConnection); err != nil {
		module.logModule.GetLogger().Infof("Start net server successful, url = %s", selfProc.ServerEp.ToString())
	} else {
		module.logModule.GetLogger().Infof("Cannot start server net, url = %s", selfProc.ServerEp.ToString())
	}

	module.serverServices[module.busModule.GetSelfBusID()] = pServer
	if pNetService := module.GetSelfNetServer(); pNetService != nil {
		pNetService.RegRegServerCallBack(module.OnRegServerCallBack)
	}

	module.registerToConsul(module.busModule.GetSelfBusID())
}

func (module *CNetServiceManagerModule) GetSelfNetServer() netServiceManagerModule.ServerService {
	return module.serverServices[module.busModule.GetSelfBusID()]
}

func (module *CNetServiceManagerModule) OnRegServerCallBack(msg *netServiceManagerModule.NetMsg, sessionID int64) {
}

func (module *CNetServiceManagerModule) registerToConsul(busID uint32) {

}
