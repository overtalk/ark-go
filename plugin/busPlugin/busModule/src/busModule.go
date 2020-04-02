package src

import (
	"fmt"
	"github.com/ArkNX/ark-go/base"
	"log"
	"path/filepath"
	"reflect"
	"runtime"

	"github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/busPlugin/busModule"
)

func init() {
	t := reflect.TypeOf((*CBusModule)(nil))
	if !t.Implements(reflect.TypeOf((*busModule.IBusModule)(nil)).Elem()) {
		log.Fatal("IBusModule is not implemented by CBusModule")
	}

	busModule.ModuleType = t.Elem()
	busModule.ModuleName = filepath.Join(busModule.ModuleType.PkgPath(), busModule.ModuleType.Name())
	busModule.ModuleUpdate = runtime.FuncForPC(reflect.ValueOf((&CBusModule{}).Update).Pointer()).Name()
}

type CBusModule struct {
	ark.Module
	// other data
	appConfig *base.AppConfig
}

func (busModule *CBusModule) Init() error {
	if err := busModule.LoadBusRelationConfig(); err != nil {
		return err
	}

	if err := busModule.LoadRegCenterConfig(); err != nil {
		return err
	}

	if err := busModule.LoadProcConfig(); err != nil {
		return err
	}

	return nil
}

// get unique signboard like `pvp-8.8.8.8`
func (busModule *CBusModule) GetAppWholeName(busID uint32) string {
	addr := base.NewBusAddrFromUInt32(busID)
	name := busModule.GetAppName(base.AppType(addr.AppType))
	return fmt.Sprintf("%s-%s", name, addr.ToString())
}

// get signboard like `pvp`
func (busModule *CBusModule) GetAppName(appType base.AppType) string {
	return busModule.appConfig.Type2names[appType]
}

// get ARKAppType by name
func (busModule *CBusModule) GetAppType(name string) base.AppType {
	return busModule.appConfig.Name2types[name]
}

// get self process information
func (busModule *CBusModule) GetSelfProc() *base.ProcConfig {
	return &busModule.appConfig.SelfProc
}

func (busModule *CBusModule) GetSelfAppType() base.AppType {
	return base.AppType(base.NewBusAddrFromUInt32(busModule.GetSelfBusID()).AppType)
}

func (busModule *CBusModule) GetSelfBusID() uint32 {
	return ark.GetPluginManagerInstance().GetBusID()
}

func (busModule *CBusModule) GetSelfBusName() string {
	return base.NewBusAddrFromUInt32(busModule.GetSelfBusID()).ToString()
}

// get bus id from `appType` & `instanceID`
// `channelID` & `zoneID` is the same as self
func (busModule *CBusModule) CombineBusID(appType base.AppType, instanceID uint8) uint32 {
	if appType < base.ARK_APP_DEFAULT || appType > base.ARK_APP_MAX {
		return 0
	}

	selfBusAddr := base.NewBusAddrFromUInt32(busModule.GetSelfBusID())
	return base.NewBusAddr(selfBusAddr.ChannelId, selfBusAddr.ZoneId, uint8(appType), instanceID).BusID()

}

// get the service registry
func (busModule *CBusModule) GetRegCenter() *base.RegCenter {
	return &busModule.appConfig.RegCenter
}

// get the app type list to connect with
func (busModule *CBusModule) GetTargetBusRelations() []base.AppType {
	selfBusAddr := base.NewBusAddrFromUInt32(busModule.GetSelfBusID())
	return busModule.appConfig.ConnectionRelations[base.AppType(selfBusAddr.AppType)]
}

////////////////////////////////////////////////////////
////////////////////////////////////////////////////////
////////////////////////////////////////////////////////
// get other process info
// use this if you do not want to use register center
func (busModule *CBusModule) GetOtherProc(busID uint32) base.ProcConfig {
	return busModule.appConfig.OtherProcList[busID]
}

// get other process list by app type
// use this if you do not want to use register center
func (busModule *CBusModule) GetOtherProcListByAppType(appType base.AppType) []base.ProcConfig {
	var ret []base.ProcConfig
	for busID, procConfig := range busModule.appConfig.OtherProcList {
		busAddr := base.NewBusAddrFromUInt32(busID)
		if base.AppType(busAddr.AppType) == appType {
			ret = append(ret, procConfig)
		}
	}
	return ret
}

////////////////////////////////////////////////////////
// private func
////////////////////////////////////////////////////////
func (busModule *CBusModule) LoadBusRelationConfig() error {
	// TODO: parse config file
	return nil
}

func (busModule *CBusModule) LoadRegCenterConfig() error {
	// TODO: parse config file
	return nil
}

func (busModule *CBusModule) LoadProcConfig() error {
	// TODO: parse config file
	return nil
}
