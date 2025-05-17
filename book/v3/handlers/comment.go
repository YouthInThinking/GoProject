package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Comment = &CommentApiHandler{}

type CommentApiHandler struct {
}

func (h *CommentApiHandler) AddComment(c *gin.Context) {
	// 实现添加评论的逻辑

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})

}
