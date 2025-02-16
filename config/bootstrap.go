package config

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

//go:embed *.yaml
var configs embed.FS

func init() {
	env := os.Getenv("env")
	vp := viper.New()
	configStream, err := configs.ReadFile("application." + env + ".yaml")
	if err != nil {
		panic(err)
	}
	vp.SetConfigType("yaml")
	err = vp.ReadConfig(bytes.NewReader(configStream))
	if err != nil {
		panic(err)
	}
	err = vp.UnmarshalKey("app", &App)
	if err != nil {
		panic(err)
	}
	err = vp.UnmarshalKey("DB", &DB)
	if err != nil {
		panic(err)
	}
	fmt.Println(*App, *DB)
}
