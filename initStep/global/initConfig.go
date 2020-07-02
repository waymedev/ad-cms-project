package global

import (
	clog "cwm.wiki/ad-CMS/common/log"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

var (
	Appkey string
	Port   string
	Mode   string
)

func InitConfig(filename string) error {

	clog.Info("filename: ", filename)

	splits := strings.Split(filepath.Base(filename), ".")
	viper.SetConfigName(splits[0])
	viper.AddConfigPath(filepath.Dir(filename))
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	initGlobal()

	return nil
}

func initGlobal() {
	Appkey = GetString("jwt.appKey")
	Port = GetString("server.port")
	Mode = GetString("server.mode")

	clog.Info(
		"Appkey: ", Appkey,
		"Port: ", Port,
		"Mode: ", Mode,
	)
}

func GetString(key string) string {

	if !viper.IsSet(key) {
		fmt.Printf("Configuration key %s not found.\n", key)
		os.Exit(1)
	}

	return viper.GetString(key)

}
