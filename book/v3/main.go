package main

import (
	"fmt"
	"os"

	"github.com/YouthInThinking/GoProject/book/v3/config"
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

	server := gin.Default()
	handlers.Book.Registry(server)
	if err := server.Run(":8080"); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
