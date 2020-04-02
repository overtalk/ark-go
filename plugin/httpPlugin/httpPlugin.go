package httpPlugin

import (
	ark "github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/httpPlugin/httpServerModule"

	_ "github.com/ArkNX/ark-go/plugin/httpPlugin/httpServerModule/src"
)

var PluginName = ark.GetName((*HttpPlugin)(nil))

type HttpPlugin struct {
	ark.Plugin
}

func init() {
	ark.GetPluginManagerInstance().AddPlugin(PluginName, NewPlugin())
}

func NewPlugin() *HttpPlugin {
	return &HttpPlugin{Plugin: ark.NewPlugin()}
}

func (httpPlugin *HttpPlugin) Install() {
	httpPlugin.Plugin.RegisterModule(httpServerModule.ModuleType, httpServerModule.ModuleUpdate)
}

func (httpPlugin *HttpPlugin) Uninstall() {
	httpPlugin.Plugin.DeregisterModule(httpServerModule.ModuleName)
}

func (httpPlugin *HttpPlugin) GetPluginName() string {
	return PluginName
}
