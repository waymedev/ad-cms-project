package initStep

import (
	clog "cwm.wiki/ad-CMS/common/log"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
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
	return nil
}

func GetString(key string) string {

	if !viper.IsSet(key) {
		fmt.Printf("Configuration key %s not found.\n", key)
		os.Exit(1)
	}

	return viper.GetString(key)

}
