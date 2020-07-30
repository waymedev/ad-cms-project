package main

import (
	clog "cwm.wiki/ad-CMS/common/log"
	"cwm.wiki/ad-CMS/controller"
	"cwm.wiki/ad-CMS/initStep/global"
	"cwm.wiki/ad-CMS/model"
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
	//err := global.InitConfig(*configFile)
	err := global.InitConfig("./config.toml")
	if err != nil {
		fmt.Println("configuration load failed.")
		os.Exit(1)
	}

	// get mode
	if global.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// get port
	port := global.Port

	if port == "" {
		port = "9998"
		clog.Warning("Default port is: ", port)
	}

	model.InitGorm()

	controller.InitGin(port)

}
