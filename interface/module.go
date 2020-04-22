package ark

import (
	"reflect"
	"runtime"
)

// IModule defines the module interface
type IModule interface {
	// lifecycle
	Init() error
	PostInit() error
	CheckConfig() error
	PreUpdate() error
	Update() error
	PreShut() error
	Shut() error
	// set & get
	GetPluginManager() *PluginManager
	GetName() string
	setPluginManager(manager *PluginManager)
	setName(name string)
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

func (module *Module) GetPluginManager() *PluginManager {
	return module.pluginManager
}

func (module *Module) setName(name string) {
	module.name = name
}

// Do nothing in the module interface
func (module *Module) setPluginManager(manager *PluginManager) {
	if manager != nil {
		module.pluginManager = manager
	}
}
