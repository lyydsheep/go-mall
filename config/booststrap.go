package config

import (
	"bytes"
	"embed"
	"fmt"
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
	fmt.Println(fileName)
	vp := viper.New()
	configFileStream, err := configs.ReadFile("application." + fileName + ".yaml")
	if err != nil {
		panic(err)
	}
	vp.SetConfigType("yaml")
	err = vp.ReadConfig(bytes.NewReader(configFileStream))
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
