package ark

import (
	"errors"
	"log"
	"path/filepath"
	"sync"

	"gopkg.in/yaml.v2"

	"github.com/ArkNX/ark-go/utils"
)

var (
	once          sync.Once
	pluginManager *PluginManager
)

type plugin struct {
	entryPoint func(*PluginManager)
	exitPoint  func(*PluginManager)
}

// ------------------- PluginManager -------------------
type PluginManager struct {
	timestamp int64  // loop timestamp
	confPath  string // plugin configuration filepath
	appName   string // app name
	logPath   string // log output path
	conf      *pluginConf

	pluginLibList             map[string]*plugin // dynamic libraries
	pluginInstanceList        map[string]IPlugin // plugin instances
	moduleInstanceList        map[string]IModule // module instances
	orderedModuleInstanceList []IModule          // ordered module instances

	moduleWithUpdateFuncList map[string]IModule // the list of modules who have the `update` function
}

func GetPluginManagerInstance() *PluginManager {
	once.Do(func() {
		pluginManager = &PluginManager{
			timestamp:                 utils.GetNowTime(),
			pluginLibList:             make(map[string]*plugin),
			pluginInstanceList:        make(map[string]IPlugin),
			moduleInstanceList:        make(map[string]IModule),
			orderedModuleInstanceList: make([]IModule, 0),
			moduleWithUpdateFuncList:  make(map[string]IModule),
		}
	})

	return pluginManager
}

func (a *PluginManager) AddPlugin(name string, p IPlugin) {
	a.pluginLibList[name] = &plugin{
		entryPoint: func(manager *PluginManager) {
			manager.Register(p)
		},
		exitPoint: func(manager *PluginManager) {
			manager.Deregister(name)
		},
	}
}

// ------------------- public func -------------------
func (a *PluginManager) Start() error {
	funcMap := []func() error{
		a.init,
		a.postInit,
		a.checkConfig,
		a.preUpdate,
	}

	for _, function := range funcMap {
		if err := function(); err != nil {
			return err
		}
	}

	return nil
}

func (a *PluginManager) Stop() error {
	funcMap := []func() error{
		a.preShut,
		a.shut,
	}

	for _, function := range funcMap {
		function()
	}

	return nil
}

func (a *PluginManager) Update() error {
	a.timestamp = utils.GetNowTime()

	for _, moduleWithUpdateFunc := range a.moduleWithUpdateFuncList {
		if moduleWithUpdateFunc == nil {
			continue
		}

		moduleWithUpdateFunc.Update()
	}
	return nil
}

func (a *PluginManager) FindModule(name string) IModule {
	return a.moduleInstanceList[name]
}

func (a *PluginManager) Register(plugin IPlugin) {
	a.register(plugin)
}

func (a *PluginManager) Deregister(name string) {
	a.deregister(a.findPlugin(name))
}

func (a *PluginManager) AddModule(moduleName string, modulePtr IModule) {
	if modulePtr == nil {
		return
	}

	a.moduleInstanceList[moduleName] = modulePtr
	a.orderedModuleInstanceList = append(a.orderedModuleInstanceList, modulePtr)
}

func (a *PluginManager) RemoveModule(moduleName string) {
	module, isExist := a.moduleInstanceList[moduleName]
	if !isExist {
		return
	}

	delete(a.moduleInstanceList, module.GetName())

	index := -1
	for tempIndex, tempModule := range a.orderedModuleInstanceList {
		if module == tempModule {
			index = tempIndex
			break
		}
	}

	length := len(a.orderedModuleInstanceList)
	if index != -1 {
		switch index {
		case 0:
			a.orderedModuleInstanceList = a.orderedModuleInstanceList[1:]
		case length:
			a.orderedModuleInstanceList = a.orderedModuleInstanceList[:length-1]
		default:
			a.orderedModuleInstanceList = append(a.orderedModuleInstanceList[:index], a.orderedModuleInstanceList[index+1:]...)
		}
	}
}

func (a *PluginManager) AddUpdateModule(module IModule) error {
	if module == nil {
		return errors.New("update module to add is nil")
	}

	a.moduleWithUpdateFuncList[module.GetName()] = module
	return nil
}

func (a *PluginManager) RemoveUpdateModule(moduleName string) {
	delete(a.moduleWithUpdateFuncList, moduleName)
}

func (a *PluginManager) GetNowTime() int64 {
	return a.timestamp
}

//func (a *PluginManager) GetBusID() uint32 {
//	return a.busID
//}
//
//func (a *PluginManager) SetBusID(id uint32) {
//	a.busID = id
//}

func (a *PluginManager) GetAppName() string {
	return a.appName
}

func (a *PluginManager) SetAppName(appName string) {
	a.appName = appName
}

//func (a *PluginManager) GetResPath() string {
//	return a.resPath
//}

func (a *PluginManager) SetPluginConf(path string) {
	if path == "" {
		return
	}

	//if !strings.Contains(path, ".plugin") {
	//	fmt.Println("failed to SetPluginConf  :", path)
	//	return
	//}

	a.confPath = path
}

func (a *PluginManager) GetLogPath() string {
	return a.logPath
}

func (a *PluginManager) SetLogPath(path string) {
	a.logPath = path
}

func (a *PluginManager) GetConfigDir(name string) string {
	for _, v := range a.conf.Plugins {
		v.Name = name
		return filepath.Join(a.conf.PluginConfDir, v.Conf)
	}
	return ""
}

// ------------------- private func -------------------

func (a *PluginManager) register(plugin IPlugin) {
	pluginName := plugin.GetPluginName()
	if a.findPlugin(pluginName) != nil {
		return
	}

	plugin.SetPluginManager(a)
	a.pluginInstanceList[pluginName] = plugin
	plugin.Install()
}

func (a *PluginManager) deregister(plugin IPlugin) {
	if plugin == nil {
		return
	}

	plugin.Uninstall()
	delete(a.pluginInstanceList, plugin.GetPluginName())
}

func (a *PluginManager) findPlugin(pluginName string) IPlugin {
	return a.pluginInstanceList[pluginName]
}

func (a *PluginManager) init() error {
	// load plugin configuration
	if err := a.loadPluginConf(); err != nil {
		return err
	}

	// load plugin
	for _, plugin := range a.pluginLibList {
		plugin.entryPoint(a)
	}

	// initialize all modules
	for _, module := range a.orderedModuleInstanceList {
		if module == nil {
			continue
		}

		if err := module.Init(); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func (a *PluginManager) postInit() error {
	for _, module := range a.orderedModuleInstanceList {
		if module == nil {
			continue
		}

		if err := module.PostInit(); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func (a *PluginManager) checkConfig() error {
	for _, module := range a.orderedModuleInstanceList {
		if module == nil {
			continue
		}

		if err := module.CheckConfig(); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func (a *PluginManager) preUpdate() error {
	for _, module := range a.orderedModuleInstanceList {
		if module == nil {
			continue
		}

		if err := module.PreUpdate(); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func (a *PluginManager) preShut() error {
	for _, module := range a.orderedModuleInstanceList {
		if module == nil {
			continue
		}

		module.PreShut()
	}

	return nil
}

func (a *PluginManager) shut() error {
	for _, module := range a.orderedModuleInstanceList {
		if module == nil {
			continue
		}

		module.Shut()
	}

	for _, plugin := range a.pluginLibList {
		plugin.exitPoint(a)
	}

	return nil
}

func (a *PluginManager) loadPluginConf() error {
	cfg := &pluginConf{}
	data, err := utils.GetBytes(a.confPath)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return err
	}

	a.conf = cfg
	return nil
}
