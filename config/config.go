package config

var (
	App *AppConfig
	DB  *DBConfig
)

type AppConfig struct {
	Env  string `yaml:"env"`
	Name string `yaml:"name"`
}

type DBConfig struct {
	Type string `yaml:"type"`
	Dsn  string `yaml:"dsn"`
	// TODO  int 可能需要改成 duration 类型
	MaxIdleCon  int `yaml:"maxIdleCons"`
	MaxIdleTime int `yaml:"maxIdleTime"`
	MaxLifeTime int `yaml:"maxLifeTime"`
}
