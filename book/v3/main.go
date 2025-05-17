package main

import (
	"fmt"
	"os"

	"github.com/YouthInThinking/GoProject/book/v3/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	handlers.Book.Registry(server)

	if err := server.Run(":8080"); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
