package msgModule

import (
	"reflect"

	"github.com/ArkNX/ark-go/base"
	ark "github.com/ArkNX/ark-go/interface"
)

var (
	ModuleName   string
	ModuleUpdate string
	ModuleType   reflect.Type
)

type IMsgModule interface {
	ark.IModule

	SendMsgByAppType(appType base.AppType, msgID uint16, msg interface{}, guid base.GUID) error
	SendMsgByBusID(busID uint32, msgID uint16, msg interface{}, guid base.GUID) error
}
