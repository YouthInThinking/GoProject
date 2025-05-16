package config

import (
	"fmt"
	"sync"

	"github.com/YouthInThinking/GoProject/book/v3/models"
	"github.com/infraboard/mcube/v2/tools/pretty"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type application struct {
	Host string `toml:"host" yaml:"jost" json:"host" env:"HOST"`
	Port int    `toml:"port" yaml:"port" json:"port" env:"PORT"`
}

//定义数据库配置结构体

type mySQL struct {
	Host     string `json:"host" yaml:"host" toml:"host" env:"DATASOURCE_HOST"`
	Port     int    `json:"port" yaml:"port" toml:"port" env:"DATASOURCE_PORT"`
	DB       string `json:"database" yaml:"database" toml:"database" env:"DATASOURCE_DB"`
	Username string `json:"username" yaml:"username" toml:"username" env:"DATASOURCE_USERNAME"`
	Password string `json:"password" yaml:"password" toml:"password" env:"DATASOURCE_PASSWORD"`
	Debug    bool   `json:"debug" yaml:"debug" toml:"debug" env:"DATASOURCE_DEBUG"`

	//gorm db对象只需要有一个，不能运行重复生成。
	db *gorm.DB `json:"-"` //不序列化到json中。

	//互斥锁，这里用锁是为了保证连接池被多个连接安全地进行读写操作，避免多个goroutine 一次性获取数据库连接导致资源竞争和错误。同时操作。
	lock sync.Mutex
}

// 使用结构体嵌套的方式定义全局变量配置结构体
type Config struct {
	Application *application `toml:"app" yaml:"app" json:"app"`
	MySQL       *mySQL       `toml:"mysql" yaml:"mysql" json:"mysql"`
}

// 定义构造函数，提供外部调用默认的配置文件
func Defalut() *Config {
	return &Config{
		Application: &application{
			Host: "localhost",
			Port: 8080,
		},
		MySQL: &mySQL{
			Host:     "172.16.160.12",
			Port:     3306,
			Username: "root",
			Password: "123456",
			DB:       "go18",
		},
	}
}

func (c *Config) String() string {
	return pretty.ToJSON(c)
}

//开始初始化数据库连接

func (m *mySQL) GetDB() *gorm.DB {
	//初始化之前就先枷锁
	m.lock.Lock()
	//defer执行在return开始之前，释放锁，这样就能保证初始化完成之前不会有别的连接进来进行操作，影响数据库初始化。
	defer m.lock.Unlock()

	//如果db已经初始化过了，就直接返回即可,如果没有初始化，那么开始执行初始化数据库。
	if m.db == nil {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			m.Username,
			m.Password,
			m.Host,
			m.Port,
			m.DB,
		)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("数据库连接失败，请检查配置文件是否正确！")
		}
		db.AutoMigrate(&models.Book{}) //自动迁移表结构

		//初始化完成之后，将db赋值给m.db，并返回即可。这样就能保证在后续的调用中，都能使用到已经初始化好的数据库连接，而不需要再次进行初始化操作。同时也能避免重复初始化带来的问题，提高代码的效率和稳定性。
		m.db = db
	}

	//返回gorm.DB对象开启debug模式
	return m.db.Debug()
}
