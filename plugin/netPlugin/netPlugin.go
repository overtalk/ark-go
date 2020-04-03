package netPlugin

import (
	ark "github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/netPlugin/netServiceManagerModule"
	_ "github.com/ArkNX/ark-go/plugin/netPlugin/netServiceManagerModule/src"
	"github.com/ArkNX/ark-go/plugin/netPlugin/httpClientModule"
	_ "github.com/ArkNX/ark-go/plugin/netPlugin/httpClientModule/src"
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
	netPlugin.Plugin.RegisterModule(netServiceManagerModule.ModuleType, netServiceManagerModule.ModuleUpdate)
	netPlugin.Plugin.RegisterModule(httpClientModule.ModuleType, httpClientModule.ModuleUpdate)
}

func (netPlugin *NetPlugin) Uninstall() {
	netPlugin.Plugin.DeregisterModule(netServiceManagerModule.ModuleName)
	netPlugin.Plugin.DeregisterModule(httpClientModule.ModuleName)
}

func (netPlugin *NetPlugin) GetPluginName() string {
	return PluginName
}