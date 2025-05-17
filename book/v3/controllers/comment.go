package controllers

import (
	"context"
	"fmt"

	"github.com/YouthInThinking/GoProject/book/v3/models"
)

var Comment = &CommentController{}

type CommentController struct {
	// 定义控制器的属性和方法
}
type AddCommentRequest struct {
	// 定义添加评论请求的结构体
	BookNumber string
}

func (c *CommentController) AddComment(ctx context.Context, in *AddCommentRequest) (*models.Comment, error) {
	// 实现添加评论的逻辑
	book, err := Book.GetBooks(ctx, NewGetBookRequest(in.BookNumber))
	if err != nil {
		return nil, err
	}

	fmt.Println("Book Number:", book)
	return nil, nil

}
