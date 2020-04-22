package ark

import (
	"log"
	"path/filepath"
	"reflect"
)

type IPlugin interface {
	GetPluginVersion() int
	GetPluginName() string
	Install()
	Uninstall()
	GetPluginManager() *PluginManager
	SetPluginManager(manager *PluginManager)
}

// ------------------- IPlugin implement -------------------
type Plugin struct {
	Modules       map[string]IModule
	pluginManager *PluginManager
}

func NewPlugin() Plugin {
	return Plugin{
		Modules:       make(map[string]IModule),
		pluginManager: nil,
	}
}

func (plugin *Plugin) GetPluginVersion() int { return 0 }
func (plugin *Plugin) GetPluginName() string { return "" }
func (plugin *Plugin) Install()              {}
func (plugin *Plugin) Uninstall()            {}
func (plugin *Plugin) GetPluginManager() *PluginManager {
	return plugin.pluginManager
}
func (plugin *Plugin) SetPluginManager(p *PluginManager) {
	plugin.pluginManager = p
}

func (plugin *Plugin) RegisterModule(t reflect.Type, update string) {
	pRegModule, ok := reflect.New(t).Interface().(IModule)
	if !ok {
		log.Fatalf("type %v should be a IModule\n", t)
	}
	pRegModuleName := filepath.Join(t.PkgPath(), t.Name())

	pluginManager := GetPluginManagerInstance()
	pRegModule.setPluginManager(pluginManager)
	pRegModule.setName(pRegModuleName)
	pluginManager.AddModule(pRegModuleName, pRegModule)
	plugin.Modules[pRegModuleName] = pRegModule

	if update != moduleUpdate {
		pluginManager.AddUpdateModule(pRegModule)
	}
}

func (plugin *Plugin) DeregisterModule(name string) {
	pluginManager := GetPluginManagerInstance()
	if pluginManager.FindModule(name) == nil {
		return
	}
	pluginManager.RemoveModule(name)
	pluginManager.RemoveUpdateModule(name)
	delete(plugin.Modules, name)
}
