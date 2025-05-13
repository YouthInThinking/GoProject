package config

import (
	"os"

	"github.com/caarlos0/env"
	"gopkg.in/yaml.v3"
)

var config *Config

// 如果没有配置文件，那么就是用默认的配置文件来进行初始化
func C() *Config {
	if config == nil {
		config = Default()
	}
	return config
}

// 将外部的yaml配置文件读取到全局变量 config 中来。
func LoadConfigFromYaml(configPath string) error {
	content, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	//读取默认值
	config = C()

	// 将yaml文件的内容解析到 config 中来。
	return yaml.Unmarshal(content, &config)
}

// 从环境变量中读取配置
func LoadConfigFromEnv() error {
	//读取默认值
	config = C()

	// 将环境变量中的配置解析到 config 中来。
	return env.Parse(config)
}
