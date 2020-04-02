package kernelSrc

import "github.com/ArkNX/ark-go/interface"

var PluginName = ark.GetName((*AFKernelPlugin)(nil))

type AFKernelPlugin struct {
	ark.Plugin
}

func NewPlugin() *AFKernelPlugin {
	return &AFKernelPlugin{Plugin: ark.NewPlugin()}
}

func (kernelPlugin *AFKernelPlugin) Install() {
	kernelPlugin.Plugin.RegisterModule(MetaClassModuleType, MetaClassModuleUpdate)
	kernelPlugin.Plugin.RegisterModule(ConfigModuleType, ConfigModuleUpdate)
	kernelPlugin.Plugin.RegisterModule(MapModuleType, MapModuleUpdate)
	kernelPlugin.Plugin.RegisterModule(KernelModuleType, KernelModuleUpdate)
}

func (kernelPlugin *AFKernelPlugin) Uninstall() {

	kernelPlugin.Plugin.DeregisterModule(MetaClassModuleName)
	kernelPlugin.Plugin.DeregisterModule(ConfigModuleName)
	kernelPlugin.Plugin.DeregisterModule(MapModuleName)
	kernelPlugin.Plugin.DeregisterModule(KernelModuleName)
}

func (kernelPlugin *AFKernelPlugin) GetPluginName() string {
	return PluginName
}
