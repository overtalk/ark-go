package module

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ArkNX/ark-go/tools/pluginBuilder/utils"
)

const str = `package src

import (
	"log"
	"path/filepath"
	"reflect"
	"runtime"

	"{{.ProjectName}}/interface"
	"{{.ProjectName}}/plugin/{{.PluginName}}Plugin/{{.ModuleName}}Module"
)

func init() {
	t := reflect.TypeOf((*C{{.UcfirstModuleName}}Module)(nil))
	if !t.Implements(reflect.TypeOf((*{{.ModuleName}}Module.I{{.UcfirstModuleName}}Module)(nil)).Elem()) {
		log.Fatal("I{{.UcfirstModuleName}}Module is not implemented by C{{.UcfirstModuleName}}Module")
	}

	{{.ModuleName}}Module.ModuleType = t.Elem()
	{{.ModuleName}}Module.ModuleName = filepath.Join({{.ModuleName}}Module.ModuleType.PkgPath(), {{.ModuleName}}Module.ModuleType.Name())
	{{.ModuleName}}Module.ModuleUpdate = runtime.FuncForPC(reflect.ValueOf((&C{{.UcfirstModuleName}}Module{}).Update).Pointer()).Name()
}

type C{{.UcfirstModuleName}}Module struct {
	ark.Module
	// other data
}

func ({{.ModuleName}}Module *C{{.UcfirstModuleName}}Module) Init() error {
	return nil
}`

const iStr = `package {{.ModuleName}}Module

import (
	"reflect"

	ark "github.com/ArkNX/ark-go/interface"
)

var (
	ModuleName   string
	ModuleUpdate string
	ModuleType   reflect.Type
)

type I{{.UcfirstModuleName}}Module interface {
	ark.IModule
}`

type Config struct {
	ProjectName       string
	PluginName        string
	ModuleName        string
	UcfirstModuleName string
}

func BuildModule(c *Config, outPath string) error {
	srcStr, err := utils.ParseTemplate(str, c)
	if err != nil {
		return err
	}

	interfaceStr, err := utils.ParseTemplate(iStr, c)
	if err != nil {
		return err
	}

	// write to disk
	srcPath := filepath.Join(outPath, fmt.Sprintf("%sPlugin", c.PluginName), fmt.Sprintf("%sModule", c.ModuleName), "src")
	if !utils.PathExists(srcPath) {
		if err := os.MkdirAll(srcPath, os.ModePerm); err != nil {
			return fmt.Errorf("failed to mkdir : %s\n", srcPath)
		}
	}
	srcPath = filepath.Join(srcPath, fmt.Sprintf("%sModule.go", c.ModuleName))
	if !utils.PathExists(srcPath) {
		if err := utils.Write(srcPath, []byte(srcStr)); err != nil {
			return err
		}
	} else {
		fmt.Printf("path %s is already exist.\n", srcPath)
		return nil
	}

	interfacePath := filepath.Join(outPath, fmt.Sprintf("%sPlugin", c.PluginName), fmt.Sprintf("%sModule", c.ModuleName), fmt.Sprintf("%sModuleInterface.go", c.ModuleName))
	if utils.PathExists(interfacePath) {
		fmt.Printf("path %s is already exist.\n", interfacePath)
		return nil
	}

	if err := utils.Write(interfacePath, []byte(interfaceStr)); err != nil {
		return err
	}

	return nil
}
