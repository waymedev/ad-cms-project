package main

import (
	clog "cwm.wiki/ad-CMS/common/log"
	"cwm.wiki/ad-CMS/initStep"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

var (
	configFile = flag.String("c", "", "Configuration File name")
)

func main() {

	// load config file
	flag.Parse()
	// load configuration file./
	err := initStep.InitConfig(*configFile)
	if err != nil {
		fmt.Println("configuration load failed.")
		os.Exit(1)
	}

	// get mode
	mode := initStep.GetString("server.mode")
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// get port
	port := initStep.GetString("server.port")

	if port == "" {
		port = "9998"
		clog.Warning("Default port is: ", port)
	}

	initStep.InitGin(port)

}
