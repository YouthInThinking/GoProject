package config

import (
	"os"

	"github.com/caarlos0/env"
	"gopkg.in/yaml.v3"
)

//配置加载

//定义全局变量，通过函数对外部提供访问

var config *Config

func C() *Config {
	//判断配置属性是否存在，如果不存在则初始化默认的配置
	if config == nil {
		config = Defalut()
	}

	//然后返回配置对象，这样就能保证全局配置是一定存在的。
	return config
}

//定义一个函数，用于将外部的文件加载到config配置属性中

func LoadConfigFromYaml(configPath string) error {
	content, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	//然后将读取的内容进行解析，赋值给默认全局变量config中。这样就能保证在后续的调用中，都能使用到已经加载好的配置信息。同时也能避免重复加载带来的问题，提高代码的效率和稳定性。
	config = C()
	return yaml.Unmarshal(content, &config)
}

// 从环境变量中读取配置
// config.MySQL.DB = os.Getenv("MYSQL_DB")
func LoadConfigFromEnv() error {
	config = C()
	return env.Parse(&config)
}
