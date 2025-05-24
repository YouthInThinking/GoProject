package response

import (
	"net/http"

	"github.com/YouthInThinking/GoProject/book/v3/exception"
	"github.com/gin-gonic/gin"
)

// 当前请求成功的时候，我们应用返回的数据
// 1. {code: 0, data: {}}
// 2. 正常直接返回数据, Restful接口 怎么知道这些请求是成功还是失败喃? 通过HTTP判断 2xx
// 如果后面 所有的返回数据 要进过特殊处理，都在这个函数内进行扩展，方便维护，比如 数据脱敏
func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	},
	)
	//终止请求处理，防止继续执行后续的中间件或路由处理函数。
	c.Abort()

}

// 当前请求失败的时候, 我们返回的数据格式
// 1. {code: xxxx, data: null, message: "错误信息"}
// 请求HTTP Code 非 2xx 就返回我们自定义的异常
//
//	{
//		"code": 404,
//		"message": "book 1 not found"
//	}

func Failed(c *gin.Context, err error) {

	// 通过类型断言，判断是否是我们自己的业务异常
	if e, ok := err.(*exception.ApiException); ok {

		//如果是，那么返回业务异常信息
		c.JSON(e.HttpCode, e)
		c.Abort()
		return
	}

	//反之，不是我们的业务异常就统一返回一个通用的错误信息。
	c.JSON(exception.CODE_SERVER_ERROR, exception.NewApiException(exception.CODE_SERVER_ERROR, "服务器内部错误").WithMessage(err.Error()))

	//终止请求处理，防止继续执行后续的中间件或路由处理函数。
	c.Abort()

}
