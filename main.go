package main

import (
	"fmt"

	"./config"
	"./logmgr"
	Log "./third_party/logrus"
)

var LOG *Log.Logger = nil

func main() {
	// load config file
	fmt.Println("Load config ...")
	//if err := config.CONF.LoadConfig(cfgFile); err != nil {
	if err := config.CONF.LoadConfig("./config/config.json"); err != nil {
		fmt.Println("Load config failed: ", err.Error())
		return
	}
	fmt.Println("Load config success")

	// init log file
	fmt.Println("Init log ...")
	if err := logmgr.InitLog(); err != nil {
		fmt.Println("Init log error: ", err.Error())
		return
	}
	LOG = logmgr.LOG
	fmt.Println("Init log success")

	caught := 12
	LOG.Debugln("Debug value:", caught)
	LOG.Infoln("Caught a signal:", caught)
}
