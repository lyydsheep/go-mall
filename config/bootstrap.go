package config

import (
	"bytes"
	"embed"
	"github.com/spf13/viper"
	"os"
)

const CONF_DIR = "config/"

func initV1() {
	fileName := os.Getenv("fileName")
	vp := viper.New()
	configFilePath := CONF_DIR + "application." + fileName + ".yaml"
	vp.SetConfigFile(configFilePath)
	err := vp.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = vp.UnmarshalKey("app", &App)
	if err != nil {
		panic(err)
	}
	err = vp.UnmarshalKey("database", &DataBase)
	if err != nil {
		panic(err)
	}
}

//go:embed *.yaml
var configs embed.FS

func init() {
	fileName := os.Getenv("fileName")
	bf, err := configs.ReadFile("application." + fileName + ".yaml")
	if err != nil {
		panic(err)
	}
	vp := viper.New()
	vp.SetConfigType("yaml")
	err = vp.ReadConfig(bytes.NewReader(bf))
	if err != nil {
		panic(err)
	}
	err = vp.UnmarshalKey("app", &App)
	if err != nil {
		panic(err)
	}
	err = vp.UnmarshalKey("database", &DataBase)
	if err != nil {
		panic(err)
	}
}
