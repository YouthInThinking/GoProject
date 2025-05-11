package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// 设置 Gin 模式为发布模式
	// gin.SetMode(gin.ReleaseMode)

	// 配置日志格式
	zerolog.TimeFieldFormat = zerolog.TimestampFunc().Local().Format("2006.01.02T15:04:05+0800")
	log.Info().Msgf("Starting server...")
	server := gin.Default()

	if err := server.Run(":8080"); err != nil {
		fmt.Println("Failed to start server:", err.Error())
		os.Exit(1)
	} else {
		fmt.Println("Server started on port 8080")

	}
}
