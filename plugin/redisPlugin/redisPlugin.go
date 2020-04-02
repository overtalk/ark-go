package redisPlugin

import (
	ark "github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/redisPlugin/redisModule"

	_ "github.com/ArkNX/ark-go/plugin/redisPlugin/redisModule/src"
)

var PluginName = ark.GetName((*RedisPlugin)(nil))

type RedisPlugin struct {
	ark.Plugin
}

func init() {
	ark.GetPluginManagerInstance().AddPlugin(PluginName, NewPlugin())
}

func NewPlugin() *RedisPlugin {
	return &RedisPlugin{Plugin: ark.NewPlugin()}
}

func (redisPlugin *RedisPlugin) Install() {
	redisPlugin.Plugin.RegisterModule(redisModule.ModuleType, redisModule.ModuleUpdate)
}

func (redisPlugin *RedisPlugin) Uninstall() {
	redisPlugin.Plugin.DeregisterModule(redisModule.ModuleName)
}

func (redisPlugin *RedisPlugin) GetPluginName() string {
	return PluginName
}
