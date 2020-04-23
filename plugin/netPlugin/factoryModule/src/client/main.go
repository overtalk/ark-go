package main

import (
	"fmt"
	"log"

	"github.com/ArkNX/ark-go/base"
	"github.com/ArkNX/ark-go/plugin/netPlugin/factoryModule/src/server"
)

func main() {
	ep, err := base.NewFromString("tcp://127.0.0.1:9999")
	if err != nil {
		log.Fatal(err)
	}

	serverService, err := server.NewServerService(handler, ep)
	if err != nil {
		log.Fatal(err)
	}

	serverService.StartServer(uint32(base.CSHeadLength), 12, 12, 12, false)
}

func handler(msg *base.NetMsg, sessionId int64) {
	fmt.Print("handler")
}
