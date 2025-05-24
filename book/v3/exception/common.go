package exception

import (
	"fmt"
	"net/http"
)

const (
	CODE_SERVER_ERROR     = 5000
	CODE_INTERNAL         = 500
	CODE_PARAM_INVALIDATE = 400
	CODE_FORBIDDEN        = 403
	CODE_NOT_FOUND        = 404
	CODE_GATEWAY_TIMEOUT  = 504
	CODE_BAD_GATEWAY      = 502
)

func ErrServerInternal(format string, a ...any) *ApiException {
	return &ApiException{
		Code:     CODE_SERVER_ERROR,
		Message:  fmt.Sprintf(format, a...),
		HttpCode: http.StatusInternalServerError,
	}

}

func ErrNotFound(format string, a ...any) *ApiException {
	return &ApiException{
		Code:     CODE_NOT_FOUND,
		Message:  fmt.Sprintf(format, a...),
		HttpCode: http.StatusNotFound,
	}

}

func ErrValidation(format string, a ...any) *ApiException {
	return &ApiException{
		Code:     CODE_PARAM_INVALIDATE,
		Message:  fmt.Sprintf(format, a...),
		HttpCode: http.StatusBadRequest,
	}

}

func ErrForbidden(format string, a ...any) *ApiException {
	return &ApiException{
		Code:     CODE_FORBIDDEN,
		Message:  fmt.Sprintf(format, a...),
		HttpCode: http.StatusForbidden,
	}

}

func ErrInternal(format string, a ...any) *ApiException {
	return &ApiException{
		Code:     CODE_INTERNAL,
		Message:  fmt.Sprintf(format, a...),
		HttpCode: http.StatusInternalServerError,
	}

}

func ErrGatewayTimeout(format string, a ...any) *ApiException {
	return &ApiException{
		Code:     CODE_GATEWAY_TIMEOUT,
		Message:  fmt.Sprintf(format, a...),
		HttpCode: http.StatusGatewayTimeout,
	}

}

func ErrBadGateway(format string, a ...any) *ApiException {
	return &ApiException{
		Code:     CODE_BAD_GATEWAY,
		Message:  fmt.Sprintf(format, a...),
		HttpCode: http.StatusBadGateway,
	}

}
