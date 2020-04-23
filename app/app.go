package app

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	ark "github.com/ArkNX/ark-go/interface"
)

const logo = `
*************************************************
   _____           __                             
  /  _  \ _______ |  | __           ____    ____  
 /  /_\  \\_  __ \|  |/ /  ______  / ___\  /  _ \ 
/    |    \|  | \/|    <  /_____/ / /_/  >(  <_> )
\____|__  /|__|   |__|_ \         \___  /  \____/ 
        \/             \/        /_____/         

Copyright 2019 (c) ArkNX. All Rights Reserved.
Website: https://arknx.com
Github:  https://github.com/ArkNX/ark-go
*************************************************

`

var (
	// version args
	commit         string
	branch         string
	version        = "no-version"
	showArkVersion bool
	// command line args
	serverName string
	configPath string
	logPath    string
)

func parseFlags() error {
	flag.StringVar(&serverName, "name", "", "Set application name")
	flag.StringVar(&configPath, "config", "", "plugin config path")
	flag.StringVar(&logPath, "log", "", "Set application log output path")
	flag.BoolVar(&showArkVersion, "v", false, "show the version")
	flag.Parse()

	// show the version of ark framwork
	if showArkVersion {
		return nil
	}

	// check the required flags
	for _, name := range []string{"name", "config", "log"} {
		found := false
		flag.Visit(func(f *flag.Flag) {
			if f.Name == name {
				found = true
			}
		})

		if !found {
			return errors.New("flag ` " + name + " ` is absent")
		}
	}

	// set app name
	ark.GetPluginManagerInstance().SetAppName(serverName)

	// set plugin config path
	ark.GetPluginManagerInstance().SetPluginConf(configPath)

	// set log path
	ark.GetPluginManagerInstance().SetLogPath(logPath)

	return nil
}

func printLogo() {
	fmt.Printf(logo)
}

func printVersion() {
	fmt.Printf("Version : %s \nBranch : %s \nCommitID : %s", version, branch, commit)
}

func Start() {
	if err := parseFlags(); err != nil {
		log.Fatal(err)
	}

	printLogo()

	if showArkVersion {
		printVersion()
		return
	}

	if err := ark.GetPluginManagerInstance().Start(); err != nil {
		log.Fatal(err)
	}

	defer ark.GetPluginManagerInstance().Stop()

	// start server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(time.Millisecond)
	for {
		select {
		case <-sigChan:
			return
		case <-ticker.C:
			ark.GetPluginManagerInstance().Update()
		}
	}
}
