package main

import (
	"fmt"
	"os"

	"github.com/YouthInThinking/GoProject/book/v3/config"
	"github.com/YouthInThinking/GoProject/book/v3/exception"
	"github.com/YouthInThinking/GoProject/book/v3/handlers"
	"github.com/gin-gonic/gin"
)

func init() {
	if err := config.LoadConfigFromYaml("application.yaml"); err != nil {

		panic(err)
	}
	config.LoadConfigFromEnv()
	// 打印加载后的配置

	// 初始化日志
	logger := config.L()

	fmt.Printf("Loaded Config:\n%s\n", config.C().String())
	logger.Info().Msg("Configuration loaded")
}
func main() {

	server := gin.New()
	server.Use(gin.Logger(), exception.Recovery())
	handlers.Book.Registry(server)
	ac := config.C().Application
	// 启动服务
	if err := server.Run(fmt.Sprintf("%s:%d", ac.Host, ac.Port)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
