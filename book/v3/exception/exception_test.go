package exception_test

import (
	"testing"

	"github.com/YouthInThinking/GoProject/book/v3/exception"
)

func CheckIsError() error {
	return exception.ErrValidation("Invalid value for field %s", "title")
}

func TestErrorValidation(t *testing.T) {
	err := CheckIsError()
	t.Log(err)

	// 怎么获取ErrorCode,要通过断言这个接口的对象的具体实现来获取类型和消息
	if v, ok := err.(*exception.ApiException); ok {
		t.Log(v.Code)
		t.Log(v.String())

	}
	t.Log(exception.IsApiException(err, exception.CODE_PARAM_INVALIDATE))
}
