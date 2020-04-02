package logPlugin

import (
	ark "github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/logPlugin/logModule"

	_ "github.com/ArkNX/ark-go/plugin/logPlugin/logModule/src"
)

var PluginName = ark.GetName((*LogPlugin)(nil))

type LogPlugin struct {
	ark.Plugin
}

func init() {
	ark.GetPluginManagerInstance().AddPlugin(PluginName, NewPlugin())
}

func NewPlugin() *LogPlugin {
	return &LogPlugin{Plugin: ark.NewPlugin()}
}

func (logPlugin *LogPlugin) Install() {
	logPlugin.Plugin.RegisterModule(logModule.ModuleType, logModule.ModuleUpdate)
}

func (logPlugin *LogPlugin) Uninstall() {
	logPlugin.Plugin.DeregisterModule(logModule.ModuleName)
}

func (logPlugin *LogPlugin) GetPluginName() string {
	return PluginName
}
