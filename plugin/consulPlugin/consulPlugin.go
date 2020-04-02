package consulPlugin

import (
	ark "github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/consulPlugin/consulModule"

	_ "github.com/ArkNX/ark-go/plugin/consulPlugin/consulModule/src"
)

var PluginName = ark.GetName((*ConsulPlugin)(nil))

type ConsulPlugin struct {
	ark.Plugin
}

func init() {
	ark.GetPluginManagerInstance().AddPlugin(PluginName, NewPlugin())
}

func NewPlugin() *ConsulPlugin {
	return &ConsulPlugin{Plugin: ark.NewPlugin()}
}

func (consulPlugin *ConsulPlugin) Install() {
	consulPlugin.Plugin.RegisterModule(consulModule.ModuleType, consulModule.ModuleUpdate)
}

func (consulPlugin *ConsulPlugin) Uninstall() {
	consulPlugin.Plugin.DeregisterModule(consulModule.ModuleName)
}

func (consulPlugin *ConsulPlugin) GetPluginName() string {
	return PluginName
}
