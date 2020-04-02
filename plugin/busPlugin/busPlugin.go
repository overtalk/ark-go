package busPlugin

import (
	ark "github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/busPlugin/busModule"
	_ "github.com/ArkNX/ark-go/plugin/busPlugin/busModule/src"
	"github.com/ArkNX/ark-go/plugin/busPlugin/msgModule"
	_ "github.com/ArkNX/ark-go/plugin/busPlugin/msgModule/src"
)

var PluginName = ark.GetName((*BusPlugin)(nil))

type BusPlugin struct {
	ark.Plugin
}

func init() {
	ark.GetPluginManagerInstance().AddPlugin(PluginName, NewPlugin())
}

func NewPlugin() *BusPlugin {
	return &BusPlugin{Plugin: ark.NewPlugin()}
}

func (busPlugin *BusPlugin) Install() {
	busPlugin.Plugin.RegisterModule(busModule.ModuleType, busModule.ModuleUpdate)
	busPlugin.Plugin.RegisterModule(msgModule.ModuleType, msgModule.ModuleUpdate)
}

func (busPlugin *BusPlugin) Uninstall() {
	busPlugin.Plugin.DeregisterModule(busModule.ModuleName)
	busPlugin.Plugin.DeregisterModule(msgModule.ModuleName)
}

func (busPlugin *BusPlugin) GetPluginName() string {
	return PluginName
}
