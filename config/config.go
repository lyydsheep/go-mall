package config

import "time"

var (
	App      appConfig
	DataBase dataBaseConfig
)

type appConfig struct {
	Env  string `yaml:"env"`
	Name string `yaml:"name"`
	Log  struct {
		FilePath    string `yaml:"path"`
		FileMaxSize int64  `yaml:"max_size"`
		FileMaxAge  int64  `yaml:"max_age"`
	}
}

type dataBaseConfig struct {
	Type        string        `yaml:"type"`
	Dsn         string        `yaml:"dsn"`
	MaxOpen     int64         `yaml:"max_open"`
	MaxIdle     int64         `yaml:"max_idle"`
	MaxFileTime time.Duration `yaml:"max_file_time"`
}
