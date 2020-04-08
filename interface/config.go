package ark

type pluginConf struct {
	PluginConfDir string `yaml:"plugin_conf_dir"`
	Plugins       []struct {
		Name string `yaml:"name"`
		Conf string `yaml:"conf"`
	} `yaml:"plugins"`
}
