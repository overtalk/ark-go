package src

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"github.com/spf13/cast"

	"github.com/ArkNX/ark-go/base"
	"github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/netPlugin/busModule"
)

var busIDStr string

func init() {
	t := reflect.TypeOf((*CBusModule)(nil))
	if !t.Implements(reflect.TypeOf((*busModule.IBusModule)(nil)).Elem()) {
		log.Fatal("IBusModule is not implemented by CBusModule")
	}

	busModule.ModuleType = t.Elem()
	busModule.ModuleName = filepath.Join(busModule.ModuleType.PkgPath(), busModule.ModuleType.Name())
	busModule.ModuleUpdate = runtime.FuncForPC(reflect.ValueOf((&CBusModule{}).Update).Pointer()).Name()

	flag.StringVar(&busIDStr, "busid", "", "Set application id(like IP address: 8.8.8.8)")
}

type CBusModule struct {
	ark.Module
	// other data
	busID     uint32
	appConfig *base.AppConfig
	busConfig *BusConfig
}

func (busModule *CBusModule) Init() error {
	// parse bus id
	busID, err := parseBusID()
	if err != nil {
		return err
	}
	busModule.busID = busID

	// parse config

	cfg, err := GetBusConfigFromYaml(busModule.GetPluginManager().GetConfigDir("BusModule"))
	if err != nil {
		return err
	}
	busModule.busConfig = cfg

	fmt.Println(busModule.busConfig)

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
	return busModule.busID
}

func (busModule *CBusModule) GetSelfBusName() string {
	return base.NewBusAddrFromUInt32(busModule.GetSelfBusID()).ToString()
}

// get bus id from `appType` & `instanceID`
// `channelID` & `zoneID` is the same as self
func (busModule *CBusModule) CombineBusID(appType base.AppType, instanceID uint8) uint32 {
	if appType < base.APPDefault || appType > base.APPMax {
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
	os.Environ()
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

////////////////////////////////////////////////////////
// utils
////////////////////////////////////////////////////////

func parseBusID() (uint32, error) {
	// parse bus id
	strArr := strings.Split(busIDStr, ".")
	if len(strArr) != 4 {
		return 0, errors.New("Bus id ` " + busIDStr + " ` is invalid, it likes 8.8.8.8")
	}

	var uint8Arr []uint8
	for _, str := range strArr {
		i, err := cast.ToUint8E(str)
		if err != nil {
			return 0, err
		}
		uint8Arr = append(uint8Arr, i)
	}
	return base.NewBusAddr(uint8Arr[0], uint8Arr[1], uint8Arr[2], uint8Arr[3]).BusID(), nil
}
