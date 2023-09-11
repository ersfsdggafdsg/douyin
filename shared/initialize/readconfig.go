package initialize

import (
	"log"

	"github.com/spf13/viper"
)

func readConf() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("../config.yaml")
	if err := v.ReadInConfig(); err != nil {
		log.Fatal("load config file failed", err)
	}
	return v
}
