package plugin

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ArkNX/ark-go/tools/pluginBuilder/utils"
)

const str = `package {{.PluginName}}Plugin

import (
	ark "github.com/ArkNX/ark-go/interface"
{{- range .ModuleNames }}
	"{{$.ProjectName}}/plugin/{{$.PluginName}}Plugin/{{.}}Module"
	_ "{{$.ProjectName}}/plugin/{{$.PluginName}}Plugin/{{.}}Module/src"
{{- end }}
)

var PluginName = ark.GetName((*{{.UcfirstPluginName}}Plugin)(nil))

type {{.UcfirstPluginName}}Plugin struct {
	ark.Plugin
}

func init() {
	ark.GetPluginManagerInstance().AddPlugin(PluginName, NewPlugin())
}

func NewPlugin() *{{.UcfirstPluginName}}Plugin {
	return &{{.UcfirstPluginName}}Plugin{Plugin: ark.NewPlugin()}
}

func ({{.PluginName}}Plugin *{{.UcfirstPluginName}}Plugin) Install() {
{{- range .ModuleNames }}
	{{$.PluginName}}Plugin.Plugin.RegisterModule({{.}}Module.ModuleType, {{.}}Module.ModuleUpdate)
{{- end }}
}

func ({{.PluginName}}Plugin *{{.UcfirstPluginName}}Plugin) Uninstall() {
{{- range .ModuleNames }}
	{{$.PluginName}}Plugin.Plugin.DeregisterModule({{.}}Module.ModuleName)
{{- end }}
}

func ({{.PluginName}}Plugin *{{.UcfirstPluginName}}Plugin) GetPluginName() string {
	return PluginName
}`

type Config struct {
	ProjectName       string
	PluginName        string
	UcfirstPluginName string
	ModuleNames       []string
}

func BuildPlugin(c *Config, outPath string) error {
	pluginStr, err := utils.ParseTemplate(str, c)
	if err != nil {
		return err
	}

	path := filepath.Join(outPath, fmt.Sprintf("%sPlugin", c.PluginName))
	if !utils.PathExists(path) {
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			return fmt.Errorf("failed to mkdir : %s\n", path)
		}
	}

	path = filepath.Join(path, fmt.Sprintf("%sPlugin.go", c.PluginName))
	if utils.PathExists(path) {
		fmt.Printf("path %s is already exist.\n", path)
		return nil
	}

	if err := utils.Write(path, []byte(pluginStr)); err != nil {
		return err
	}
	return nil
}
