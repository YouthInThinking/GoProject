package controllers_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/YouthInThinking/GoProject/book/v3/config"
	"github.com/YouthInThinking/GoProject/book/v3/controllers"
	"github.com/YouthInThinking/GoProject/book/v3/models"
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

func TestCreateBook(t *testing.T) {
	book, err := controllers.Book.CreateBooks(context.Background(), &models.BookSpec{
		Title:  "Unit Test Book for go Controller test",
		Author: "Test Author",
		Price:  99.99,
		IsSale: nil, // Add this line to fix the missing field error
	})

	if err != nil {
		t.Fatal("not implemented", err)
	}
	t.Logf("Book: %v", book)
}
