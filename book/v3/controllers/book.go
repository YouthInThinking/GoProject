package controllers

import (
	"context"
	"fmt"

	"github.com/YouthInThinking/GoProject/book/v3/config"
	"github.com/YouthInThinking/GoProject/book/v3/exception"
	"github.com/YouthInThinking/GoProject/book/v3/models"
	"gorm.io/gorm"
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
	fmt.Printf("日志级别为：%v\n", config.C().Log.Level)

	config.L().Debug().Msgf("GetBooks: %d", in.BookNumber)
	fmt.Printf("日志级别为：%v\n", config.C().Log.Level)
	// 验证配置
	config.L().Info().Msg("logger initialization completed")
	fmt.Printf("日志级别为：%v\n", config.C().Log.Level)

	bookInstence := &models.Book{}
	if err := config.DB().Where("id = ?", in.BookNumber).First(bookInstence).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, exception.ErrNotFound("Book not found with id %d", in.BookNumber)
		} else if err == gorm.ErrInvalidValue {
			return nil, exception.ErrValidation("Invalid value for field %d", in.BookNumber)
		} else {
			return nil, err
		}
	}

	return bookInstence, nil

}

func (c *BookController) CreateBooks(ctx context.Context, in *models.BookSpec) (*models.Book, error) {
	bookInstance := &models.Book{
		BookSpec: *in,
	}
	// 数据入库(Grom), 补充自增Id的值
	if err := config.DB().Save(&bookInstance).Error; err != nil {
		return nil, err
	}
	return bookInstance, nil
}
