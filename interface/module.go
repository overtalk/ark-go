package ark

import (
	"reflect"
	"runtime"
)

type IModule interface {
	Init() error
	PostInit() error
	CheckConfig() error
	PreUpdate() error
	Update() error
	PreShut() error
	Shut() error
	GetPluginManager() *PluginManager
	SetPluginManager(manager *PluginManager)
	GetName() string
	SetName(name string)
}

var moduleUpdate = runtime.FuncForPC(reflect.ValueOf((&Module{}).Update).Pointer()).Name()

// ------------------- IModule implement -------------------
// ------------------- Eclectic solution for c++ macro -------------------
type Module struct {
	pluginManager *PluginManager
	name          string
}

func (module *Module) Init() error        { return nil }
func (module *Module) PostInit() error    { return nil }
func (module *Module) CheckConfig() error { return nil }
func (module *Module) PreUpdate() error   { return nil }
func (module *Module) Update() error      { return nil }
func (module *Module) PreShut() error     { return nil }
func (module *Module) Shut() error        { return nil }

// Do nothing in the module interface
func (module *Module) GetName() string {
	return module.name
}

func (module *Module) SetName(name string) {
	module.name = name
}

func (module *Module) GetPluginManager() *PluginManager {
	return module.pluginManager
}

// Do nothing in the module interface
func (module *Module) SetPluginManager(manager *PluginManager) {
	if manager != nil {
		module.pluginManager = manager
	}
}
