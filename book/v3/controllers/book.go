package controllers

import (
	"context"

	"github.com/YouthInThinking/GoProject/book/v3/config"
	"github.com/YouthInThinking/GoProject/book/v3/models"
)

var Book = &BookController{}

type BookController struct {
	// 定义控制器的属性和方法
}

type GetBookRequest struct {
	BookNumber int
}

func NewGetBookRequest(bookNumber int) *GetBookRequest {
	return &GetBookRequest{
		BookNumber: bookNumber,
	}
}

func (c *BookController) GetBooks(ctx context.Context, in *GetBookRequest) (*models.Book, error) {
	// 实现获取所有书籍的逻辑

	config.L().Error().Msgf("GetBooks: %d", in.BookNumber)
	config.L().Debug().Msgf("GetBooks: %d", in.BookNumber)
	bookInstence := &models.Book{}
	if err := config.DB().Where("id = ?", in.BookNumber).First(bookInstence).Error; err != nil {
		return nil, err
	}

	return bookInstence, nil

}
