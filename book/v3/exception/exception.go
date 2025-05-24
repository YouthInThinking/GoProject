package exception

import (
	"errors"
	"fmt"

	"github.com/infraboard/mcube/v2/tools/pretty"
)

// 用于描述业务异常
// 实现自定义异常
type ApiException struct {
	//业务异常编码，用于标识异常类型，500001 表示Token过期
	Code int `json:"code"`
	//业务异常消息，用于描述异常信息
	Message string `json:"message"`
	//HTTP状态码，用于标识HTTP响应状态 不会出现在Body里面，不会序列化为json，http response 进行set
	HttpCode int `json:"-"`
}

// 将业务异常信息自定义格式输出错误信息
func (e *ApiException) Error() string {
	return fmt.Sprintf("error code: %d, message: %s", e.Code, e.Message)
}

// 自定义输出格式
func (e *ApiException) String() string {
	/* 	dj, _ := json.MarshalIndent(e, "", "  ")
	   	return string(dj) */
	return pretty.ToJSON(e)
}

// 构造函数，返回业务异常字段属性以供调用
func NewApiException(code int, message string) *ApiException {
	return &ApiException{
		Code:    code,
		Message: message,
	}
}

// 设置业务异常消息，返回当前对象以供链式调用
func (e *ApiException) WithMessage(msg string) *ApiException {
	e.Message = msg
	return e
}

// 设置HTTP状态码，返回当前对象以供链式调用
func (e *ApiException) WithHttpCode(httpCode int) *ApiException {
	e.HttpCode = httpCode
	return e
}

// 通过Code来比较错误
func IsApiException(err error, code int) bool {
	var apiErr *ApiException
	if err != nil && errors.As(err, &apiErr) {
		return apiErr.Code == code
	}
	return false

}
