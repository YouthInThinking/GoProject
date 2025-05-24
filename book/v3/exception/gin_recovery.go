package exception

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 自定义异常处理机制
func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, err any) {
		// 非业务异常
		c.JSON(CODE_INTERNAL, NewApiException(CODE_INTERNAL, fmt.Sprintf("%#v", err)))
		c.Abort()
	})
}
