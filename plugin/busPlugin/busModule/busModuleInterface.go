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

type AFIBusModule interface {
	ark.AFIModule

	// get unique signboard like `pvp-8.8.8.8`
	GetAppWholeName(busID uint32) string

	// get signboard like `pvp`
	GetAppName(appType base.ARKAppType) string

	// get ARKAppType by name
	GetAppType(name string) base.ARKAppType

	// get self process information
	GetSelfProc() *base.AFProcConfig
	GetSelfAppType() base.ARKAppType
	GetSelfBusID() uint32
	GetSelfBusName() string

	// get bus id from `appType` & `instanceID`
	// `channelID` & `zoneID` is the same as self
	CombineBusID(appType base.ARKAppType, instanceID uint8)

	// get the service registry
	GetRegCenter() *base.AFRegCenter

	// get the app type list to connect with
	GetTargetBusRelations() ([]base.ARKAppType, error)

	////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////
	// get other process info
	// use this if you do not want to use register center
	GetOtherProc(busID uint32) (base.AFProcConfig, error)

	// get other process list by app type
	// use this if you do not want to use register center
	GetOtherProcListByAppType(appType base.ARKAppType) (int, []base.AFProcConfig)
}
