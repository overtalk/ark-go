package src

import (
	"errors"
	"reflect"

	"gopkg.in/yaml.v2"

	"github.com/ArkNX/ark-go/utils"
)

type Consul struct {
	IP            string `yaml:"ip" env:"Consul_IP"`
	Port          int    `yaml:"port" env:"Consul_Port"`
	CheckInterval int    `yaml:"check_interval" env:"Consul_CheckInterval"`
	CheckTimeout  int    `yaml:"check_timeout" env:"Consul_CheckTimeout"`
}

type Relation struct {
	Source string `yaml:"source" env:"Consul_Source"`
	Target string `yaml:"target" env:"Consul_Target"`
}

type Application struct {
	Name string `yaml:"name" env:"Application_Name"`
	Type int    `yaml:"type" env:"Application_Type"`
}

type Process struct {
	Busid            string `yaml:"busid"  env:"Process_Busid"`
	EndpointServer   string `yaml:"endpoint_server"  env:"Process_EndpointServer"`
	EndpointIntranet string `yaml:"endpoint_intranet"  env:"Process_EndpointIntranet"`
	MaxConnection    int    `yaml:"max_connection"  env:"Process_MaxConnection"`
	ThreadNum        int    `yaml:"thread_num"  env:"Process_ThreadNum"`
}

type BusConfig struct {
	Consul       *Consul        `yaml:"consul"`
	Applications []*Application `yaml:"applications"`
	Relations    []*Relation    `yaml:"relations"`
	Process      []*Process     `yaml:"process"`
}

func GetBusConfigFromYaml(configPath string) (*BusConfig, error) {
	if len(configPath) == 0 {
		return nil, errors.New("config for bus module is absent")
	}

	bc := &BusConfig{}

	data, err := utils.GetBytes(configPath)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, bc); err != nil {
		return nil, err
	}

	return bc, nil
}

func GetBusConfigFromEnv() (*BusConfig, error) {
	bc := &BusConfig{}

	// config
	c := &Consul{}
	flag, err := utils.ParseStruct(-1, reflect.TypeOf(*c), reflect.ValueOf(c).Elem())
	if err != nil {
		return nil, err
	}
	if flag {
		bc.Consul = c
	}

	// application
	index := 0
	for {
		r := &Application{}
		flag, err := utils.ParseStruct(index, reflect.TypeOf(*r), reflect.ValueOf(r).Elem())
		if err != nil {
			return nil, err
		}
		if !flag {
			break
		}
		bc.Applications = append(bc.Applications, r)
		index++
	}

	// relation
	index = 0
	for {
		r := &Relation{}
		flag, err := utils.ParseStruct(index, reflect.TypeOf(*r), reflect.ValueOf(r).Elem())
		if err != nil {
			return nil, err
		}
		if !flag {
			break
		}
		bc.Relations = append(bc.Relations, r)
		index++
	}

	// Process
	index = 0
	for {
		r := &Process{}
		flag, err := utils.ParseStruct(index, reflect.TypeOf(*r), reflect.ValueOf(r).Elem())
		if err != nil {
			return nil, err
		}
		if !flag {
			break
		}
		bc.Process = append(bc.Process, r)
		index++
	}

	return bc, nil
}
