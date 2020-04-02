package busPlugin

import (
	ark "github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/busPlugin/busModule"
	_ "github.com/ArkNX/ark-go/plugin/busPlugin/busModule/src"
	"github.com/ArkNX/ark-go/plugin/busPlugin/msgModule"
	_ "github.com/ArkNX/ark-go/plugin/busPlugin/msgModule/src"
)

var PluginName = ark.GetName((*AFBusPlugin)(nil))

type AFBusPlugin struct {
	ark.AFCPlugin
}

func init() {
	ark.GetAFPluginManagerInstance().AddPlugin(PluginName, NewPlugin())
}

func NewPlugin() *AFBusPlugin {
	return &AFBusPlugin{AFCPlugin: ark.NewAFCPlugin()}
}

func (busPlugin *AFBusPlugin) Install() {
	busPlugin.AFCPlugin.RegisterModule(busModule.ModuleType, busModule.ModuleUpdate)
	busPlugin.AFCPlugin.RegisterModule(msgModule.ModuleType, msgModule.ModuleUpdate)
}

func (busPlugin *AFBusPlugin) Uninstall() {
	busPlugin.AFCPlugin.DeregisterModule(busModule.ModuleName)
	busPlugin.AFCPlugin.DeregisterModule(msgModule.ModuleName)
}

func (busPlugin *AFBusPlugin) GetPluginName() string {
	return PluginName
}
