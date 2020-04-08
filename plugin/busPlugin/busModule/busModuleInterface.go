package busModule

import (
	"github.com/ArkNX/ark-go/base"
	"reflect"

	ark "github.com/ArkNX/ark-go/interface"
)

var (
	ModuleName   string
	ModuleUpdate string
	ModuleType   reflect.Type
)

type IBusModule interface {
	ark.IModule

	// get unique signboard like `pvp-8.8.8.8`
	GetAppWholeName(busID uint32) string

	// get signboard like `pvp`
	GetAppName(appType base.AppType) string

	// get ARKAppType by name
	GetAppType(name string) base.AppType

	// get self process information
	GetSelfProc() *base.ProcConfig
	GetSelfAppType() base.AppType
	GetSelfBusID() uint32
	GetSelfBusName() string

	// get bus id from `appType` & `instanceID`
	// `channelID` & `zoneID` is the same as self
	CombineBusID(appType base.AppType, instanceID uint8) uint32

	// get the service registry
	GetRegCenter() *base.RegCenter

	// get the app type list to connect with
	GetTargetBusRelations() []base.AppType

	////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////
	// get other process info
	// use this if you do not want to use register center
	GetOtherProc(busID uint32) base.ProcConfig

	// get other process list by app type
	// use this if you do not want to use register center
	GetOtherProcListByAppType(appType base.AppType) []base.ProcConfig
}
