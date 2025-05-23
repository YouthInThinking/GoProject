package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env"
	"github.com/rs/zerolog"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

//配置加载

//定义全局变量，通过函数对外部提供访问

var config *Config

func C() *Config {
	//判断配置属性是否存在，如果不存在则初始化默认的配置
	if config == nil {
		config = Defalut()
	}

	//存在就返回配置对象，这样就能保证全局配置是一定存在的。
	return config
}
func L() *zerolog.Logger {
	return C().Logger()
}

func DB() *gorm.DB {
	return C().MySQL.GetDB()
}

//定义一个函数，用于将外部的文件加载到config配置属性中
/*
func LoadConfigFromYaml(configPath string) error {
	content, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	//然后将读取的内容进行解析，赋值给默认全局变量config中。这样就能保证在后续的调用中，都能使用到已经加载好的配置信息。同时也能避免重复加载带来的问题，提高代码的效率和稳定性。
	 // 错误点：直接修改全局 config 实例
	config = C() // 这里获取的是已存在的默认配置实例，而不是重新创建一个新的实例。这会导致后续的修改会影响到全局配置。
	return yaml.Unmarshal(content, config)

	这会直接修改全局 config 实例，但 YAML 解析时 嵌套结构体字段（如 log.rotate）可能无法正确覆盖。因此，建议在解析 YAML 后创建一个新的配置实例，并将其赋值给全局变量。
    如果 YAML 文件中的 log.level 字段未被正确映射，会保留默认值 debug
}
*/

// 修改 LoadConfigFromYaml 函数，确保每次加载都创建新实例：
func LoadConfigFromYaml(configPath string) error {
	content, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	// 创建新的配置实例（关键修正）
	newConfig := C()
	if err := yaml.Unmarshal(content, newConfig); err != nil {
		return fmt.Errorf("failed to parse YAML: %w", err)
	}

	// 替换全局配置
	config = newConfig
	return nil
}

// 从环境变量中读取配置
// config.MySQL.DB = os.Getenv("MYSQL_DB")
func LoadConfigFromEnv() error {
	config = C()
	return env.Parse(config)
}
