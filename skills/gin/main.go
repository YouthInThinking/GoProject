package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 设置 Gin 模式为发布模式
	// gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	if err := server.Run(":8080"); err != nil {
		fmt.Println("Failed to start server:", err.Error())
		os.Exit(1)
	} else {
		fmt.Println("Server started on port 8080")

	}
}
