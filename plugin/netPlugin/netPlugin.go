package netPlugin

import (
	ark "github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/netPlugin/busModule"
	_ "github.com/ArkNX/ark-go/plugin/netPlugin/busModule/src"
	"github.com/ArkNX/ark-go/plugin/netPlugin/msgModule"
	_ "github.com/ArkNX/ark-go/plugin/netPlugin/msgModule/src"
	"github.com/ArkNX/ark-go/plugin/netPlugin/factoryModule"
	_ "github.com/ArkNX/ark-go/plugin/netPlugin/factoryModule/src"
)

var PluginName = ark.GetName((*NetPlugin)(nil))

type NetPlugin struct {
	ark.Plugin
}

func init() {
	ark.GetPluginManagerInstance().AddPlugin(PluginName, NewPlugin())
}

func NewPlugin() *NetPlugin {
	return &NetPlugin{Plugin: ark.NewPlugin()}
}

func (netPlugin *NetPlugin) Install() {
	netPlugin.Plugin.RegisterModule(busModule.ModuleType, busModule.ModuleUpdate)
	netPlugin.Plugin.RegisterModule(msgModule.ModuleType, msgModule.ModuleUpdate)
	netPlugin.Plugin.RegisterModule(factoryModule.ModuleType, factoryModule.ModuleUpdate)
}

func (netPlugin *NetPlugin) Uninstall() {
	netPlugin.Plugin.DeregisterModule(busModule.ModuleName)
	netPlugin.Plugin.DeregisterModule(msgModule.ModuleName)
	netPlugin.Plugin.DeregisterModule(factoryModule.ModuleName)
}

func (netPlugin *NetPlugin) GetPluginName() string {
	return PluginName
}