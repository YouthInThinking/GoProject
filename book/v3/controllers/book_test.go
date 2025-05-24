package controllers_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/YouthInThinking/GoProject/book/v3/config"
	"github.com/YouthInThinking/GoProject/book/v3/controllers"
)

func init() {
	if err := config.LoadConfigFromYaml(fmt.Sprintf("%s/GoProject/book/v3/application.yaml", os.Getenv("workspaceFolder"))); err != nil {

		panic(err)
	}
	config.LoadConfigFromEnv()
	// 打印加载后的配置

	// 初始化日志
	logger := config.L()

	fmt.Printf("Loaded Config:\n%s\n", config.C().String())
	logger.Info().Msg("Configuration loaded")
}

func TestBooks(t *testing.T) {
	book, err := controllers.Book.GetBooks(context.Background(), &controllers.GetBookRequest{BookNumber: 3})

	if err != nil {
		t.Fatal("not implemented", err)
	}
	t.Logf("Book: %v", book)

}
