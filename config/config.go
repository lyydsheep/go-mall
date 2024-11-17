package config

import "time"

var (
	App      appConfig
	DataBase dataBaseConfig
)

type appConfig struct {
	// 用mapstructure标签就有效
	Env  string `mapstructure:"env"`
	Name string `mapstructure:"name"`
	Log  struct {
		FilePath    string `mapstructure:"path"`
		FileMaxSize int    `mapstructure:"max_size"`
		FileMaxAge  int    `mapstructure:"max_age"`
	}
}

type dataBaseConfig struct {
	Type        string        `mapstructure:"type"`
	Dsn         string        `mapstructure:"dsn"`
	MaxOpen     int64         `mapstructure:"max_open"`
	MaxIdle     int64         `mapstructure:"max_idle"`
	MaxFileTime time.Duration `mapstructure:"max_file_time"`
}
