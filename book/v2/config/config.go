package config

// 此对象是程序的应用服务类型配置
type Config struct {
	Application *application `toml:"app" yaml:"app" json:"app"`
	MySQL       *mySQL       `toml:"mysql" yaml:"mysql" json:"mysql"`
}

// 此对象是程序的程序类型配置
type application struct {
	Host string `toml:"host" yaml:"host" json:"host" env:"HOST"`
	Port int    `toml:"port" yaml:"port" json:"port" env:"PORT"`
}

// 此对象是程序的数据库类型配置
type mySQL struct {
	Host     string `json:"host" yaml:"host" toml:"host" env:"DATASOURCE_HOST"`
	Port     int    `json:"port" yaml:"port" toml:"port" env:"DATASOURCE_PORT"`
	DB       string `json:"database" yaml:"database" toml:"database" env:"DATASOURCE_DB"`
	Username string `json:"username" yaml:"username" toml:"username" env:"DATASOURCE_USERNAME"`
	Password string `json:"password" yaml:"password" toml:"password" env:"DATASOURCE_PASSWORD"`
	Debug    bool   `json:"debug" yaml:"debug" toml:"debug" env:"DATASOURCE_DEBUG"`
}

// 通过创建构造函数，将对象类型的配置信息初始化为默认值。方便通过构造函数调用数据类型字段
func Default() *Config {
	return &Config{
		Application: &application{
			Host: "127.0.0.1",
			Port: 8080,
		},
		MySQL: &mySQL{
			Host:     "127.0.0.1",
			Port:     3306,
			DB:       "go18",
			Username: "root",
			Password: "123456",
			Debug:    true,
		},
	}
}
