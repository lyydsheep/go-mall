package config

var (
	App *AppConfig
	DB  *DBConfig
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
	Type string `mapstructure:"type"`
	Dsn  string `mapstructure:"dsn"`
	// TODO  int 可能需要改成 duration 类型
	MaxIdleCon  int `mapstructure:"maxIdleCons"`
	MaxIdleTime int `mapstructure:"maxIdleTime"`
	MaxLifeTime int `mapstructure:"maxLifeTime"`
}
