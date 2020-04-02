package kernelInterface

import (
	ark "github.com/ArkNX/ark-go/interface"
)

var AFIMetaClassModuleName string

type AFIMetaClassModule interface {
	ark.IModule
	Load() error
	AddClassCallBack(className string)
}
