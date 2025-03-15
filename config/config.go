package config

var (
	App AppConfig
	DB  DBConfig
)

type AppConfig struct {
	Env  string `mapstructure:"env"`
	Name string `mapstructure:"name"`
	Log  struct {
		Path    string `mapstructure:"path"`
		MaxSize int    `mapstructure:"maxSize"`
		MaxAge  int    `mapstructure:"maxAge"`
	} `mapstructure:"log"`
	Pagination struct {
		DefaultSize int `mapstructure:"defaultSize"`
		MaxSize     int `mapstructure:"max_size"`
	} `mapstructure:"pagination"`
}

type DBConfig struct {
	Master DBConfigOptions `mapstructure:"master"`
	Slave  DBConfigOptions `mapstructure:"slave"`
}

type DBConfigOptions struct {
	Type        string `mapstructure:"type"`
	Dsn         string `mapstructure:"dsn"`
	MaxOpen     int    `mapstructure:"maxOpen"`
	MaxIdle     int    `mapstructure:"maxIdle"`
	MaxLifeTime int    `mapstructure:"maxLifeTime"`
}
