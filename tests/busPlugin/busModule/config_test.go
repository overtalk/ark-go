package busModule

import (
	"fmt"
	"github.com/ArkNX/ark-go/plugin/netPlugin/busModule/src"
	"os"
	"strings"
	"testing"
)

func TestGetBusConfigFromYaml(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
		return
	}
	path := strings.ReplaceAll(pwd, "tests/busPlugin/busModule", "build/bus.module.yaml")

	bc, err := src.GetBusConfigFromYaml(path)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(bc.Consul)
	t.Log("---------Process---------")
	for k, v := range bc.Process {
		t.Log(k, v)
	}
	t.Log("--------Applications----------")
	for k, v := range bc.Applications {
		t.Log(k, v)
	}
	t.Log("--------Relations----------")
	for k, v := range bc.Relations {
		t.Log(k, v)
	}
}

func TestGetBusConfigFromEnv(t *testing.T) {
	setEnv()

	bc, err := src.GetBusConfigFromEnv()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(bc.Consul)
	t.Log("---------Process---------")
	for k, v := range bc.Process {
		t.Log(k, v)
	}
	t.Log("--------Applications----------")
	for k, v := range bc.Applications {
		t.Log(k, v)
	}
	t.Log("--------Relations----------")
	for k, v := range bc.Relations {
		t.Log(k, v)
	}
}

func setEnv() {
	// set consul
	os.Setenv("Consul_IP", "0.0.0.0")
	os.Setenv("Consul_Port", "9999")
	os.Setenv("Consul_CheckInterval", "12")
	os.Setenv("Consul_CheckTimeout", "10")

	// Relation
	for i := 0; i < 10; i++ {
		os.Setenv(fmt.Sprintf("Consul_Source_%d", i), fmt.Sprintf("game-%d", i))
		os.Setenv(fmt.Sprintf("Consul_Target_%d", i), fmt.Sprintf("pvp-%d", i))
	}

	// Application
	for i := 0; i < 10; i++ {
		os.Setenv(fmt.Sprintf("Application_Name_%d", i), fmt.Sprintf("game%d", i))
		os.Setenv(fmt.Sprintf("Application_Type_%d", i), fmt.Sprintf("%d", i))
	}

	// Process
	for i := 0; i < 10; i++ {
		os.Setenv(fmt.Sprintf("Process_Busid_%d", i), fmt.Sprintf("%d.%d.%d.%d", i, i, i, i))
		os.Setenv(fmt.Sprintf("Process_EndpointServer_%d", i), fmt.Sprintf("47.101.184.248:%d", 100+i))
		os.Setenv(fmt.Sprintf("Process_EndpointIntranet_%d", i), fmt.Sprintf("127.0.0.1:%d", 100+i))
		os.Setenv(fmt.Sprintf("Process_MaxConnection_%d", i), fmt.Sprintf("%d", i))
		os.Setenv(fmt.Sprintf("Process_ThreadNum_%d", i), fmt.Sprintf("%d", i))
	}
}
