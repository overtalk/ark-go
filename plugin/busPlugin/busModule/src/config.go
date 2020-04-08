package src

type BusConfig struct {
	Consul struct {
		IP            string `yaml:"ip"`
		Port          int    `yaml:"port"`
		CheckInterval int    `yaml:"check_interval"`
		CheckTimeout  int    `yaml:"check_timeout"`
	} `yaml:"consul"`
	Applications []struct {
		Name string `yaml:"name"`
		Type int    `yaml:"type"`
	} `yaml:"applications"`
	Relations []struct {
		Source string `yaml:"source"`
		Target string `yaml:"target"`
	} `yaml:"relations"`
	Process []struct {
		Busid            string `yaml:"busid"`
		EndpointServer   string `yaml:"endpoint_server"`
		EndpointIntranet string `yaml:"endpoint_intranet"`
		MaxConnection    int    `yaml:"max_connection"`
		ThreadNum        int    `yaml:"thread_num"`
	} `yaml:"process"`
}
