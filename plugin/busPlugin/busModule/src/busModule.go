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
	t := reflect.TypeOf((*AFCBusModule)(nil))
	if !t.Implements(reflect.TypeOf((*busModule.AFIBusModule)(nil)).Elem()) {
		log.Fatal("AFIBusModule is not implemented by AFCBusModule")
	}

	busModule.ModuleType = t.Elem()
	busModule.ModuleName = filepath.Join(busModule.ModuleType.PkgPath(), busModule.ModuleType.Name())
	busModule.ModuleUpdate = runtime.FuncForPC(reflect.ValueOf((&AFCBusModule{}).Update).Pointer()).Name()
}

type AFCBusModule struct {
	ark.AFCModule
	// other data
	appConfig *base.AFAppConfig
}

func (busModule *AFCBusModule) Init() error {
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
func (busModule *AFCBusModule) GetAppWholeName(busID uint32) string {
	addr := base.NewAFBusAddrFromUInt32(busID)
	name := busModule.GetAppName(base.ARKAppType(addr.AppType))
	return fmt.Sprintf("%s-%s", name, addr.ToString())
}

// get signboard like `pvp`
func (busModule *AFCBusModule) GetAppName(appType base.ARKAppType) string {
	return busModule.appConfig.Type2names[appType]
}

// get ARKAppType by name
func (busModule *AFCBusModule) GetAppType(name string) base.ARKAppType {
	return busModule.appConfig.Name2types[name]
}

// get self process information
func (busModule *AFCBusModule) GetSelfProc() *base.AFProcConfig {
	return &busModule.appConfig.SelfProc
}

func (busModule *AFCBusModule) GetSelfAppType() base.ARKAppType {
	return base.ARKAppType(base.NewAFBusAddrFromUInt32(busModule.GetSelfBusID()).AppType)
}

func (busModule *AFCBusModule) GetSelfBusID() uint32 {
	return ark.GetAFPluginManagerInstance().GetBusID()
}

func (busModule *AFCBusModule) GetSelfBusName() string {
	return base.NewAFBusAddrFromUInt32(busModule.GetSelfBusID()).ToString()
}

// get bus id from `appType` & `instanceID`
// `channelID` & `zoneID` is the same as self
func (busModule *AFCBusModule) CombineBusID(appType base.ARKAppType, instanceID uint8) uint32 {
	if appType < base.ARK_APP_DEFAULT || appType > base.ARK_APP_MAX {
		return 0
	}

	selfBusAddr := base.NewAFBusAddrFromUInt32(busModule.GetSelfBusID())
	return base.NewAFBusAddr(selfBusAddr.ChannelId, selfBusAddr.ZoneId, uint8(appType), instanceID).BusID()

}

// get the service registry
func (busModule *AFCBusModule) GetRegCenter() *base.AFRegCenter {
	return &busModule.appConfig.RegCenter
}

// get the app type list to connect with
func (busModule *AFCBusModule) GetTargetBusRelations() []base.ARKAppType {
	selfBusAddr := base.NewAFBusAddrFromUInt32(busModule.GetSelfBusID())
	return busModule.appConfig.ConnectionRelations[base.ARKAppType(selfBusAddr.AppType)]
}

////////////////////////////////////////////////////////
////////////////////////////////////////////////////////
////////////////////////////////////////////////////////
// get other process info
// use this if you do not want to use register center
func (busModule *AFCBusModule) GetOtherProc(busID uint32) base.AFProcConfig {
	return busModule.appConfig.OtherProcList[busID]
}

// get other process list by app type
// use this if you do not want to use register center
func (busModule *AFCBusModule) GetOtherProcListByAppType(appType base.ARKAppType) []base.AFProcConfig {
	var ret []base.AFProcConfig
	for busID, procConfig := range busModule.appConfig.OtherProcList {
		busAddr := base.NewAFBusAddrFromUInt32(busID)
		if base.ARKAppType(busAddr.AppType) == appType {
			ret = append(ret, procConfig)
		}
	}
	return ret
}

////////////////////////////////////////////////////////
// private func
////////////////////////////////////////////////////////
func (busModule *AFCBusModule) LoadBusRelationConfig() error {
	// TODO: parse config file
	return nil
}

func (busModule *AFCBusModule) LoadRegCenterConfig() error {
	// TODO: parse config file
	return nil
}

func (busModule *AFCBusModule) LoadProcConfig() error {
	// TODO: parse config file
	return nil
}
