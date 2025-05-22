package controllers_test

import (
	"context"
	"testing"

	"github.com/YouthInThinking/GoProject/book/v3/controllers"
)

func TestBooks(t *testing.T) {
	book, err := controllers.Book.GetBooks(context.Background(), &controllers.GetBookRequest{BookNumber: 3})

	if err != nil {
		t.Fatal("not implemented", err)
	}
	t.Logf("Book: %v", book)

}
