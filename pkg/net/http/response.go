package http

import (
	"time"

	"github.com/gin-gonic/gin"
	"snowdream.tech/http-server/pkg/i18n"
)

// Response Response
type Response struct {
	Version   string `json:"version" xml:"version" yaml:"version" schema:"version"`
	Code      string `json:"code" xml:"code" yaml:"code" schema:"code"`
	Message   string `json:"message" xml:"message" yaml:"message" schema:"message"`
	Data      any    `json:"data" xml:"data" yaml:"data" schema:"data"`
	Timestamp int64  `json:"timestamp" xml:"timestamp" yaml:"timestamp" schema:"timestamp"`
	RequestID string `json:"requestid" xml:"requestid" yaml:"requestid" schema:"requestid"`
}

// NewResponse NewResponse
func NewResponse(code string, message string, data any) Response {
	if nil == data {
		return Response{
			Version:   "0.1",
			Code:      code,
			Message:   message,
			Data:      struct{}{},
			Timestamp: time.Now().Unix(),
		}
	}

	return Response{
		Version:   "0.1",
		Code:      code,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().Unix(),
	}
}

// ResponseSuccess ResponseSuccess
func ResponseSuccess(c *gin.Context) Response {
	i18 := i18n.Default(c)

	str := i18.T(c, "SUCCESS")

	return NewResponse("SUCCESS", str, nil)
}

// ResponseSuccessWithMessage ResponseSuccessWithMessage
func ResponseSuccessWithMessage(c *gin.Context, key string, params ...string) Response {
	i18 := i18n.Default(c)

	paramarrs := make([]any, 0)

	for i := 0; i < len(params); i++ {
		param := i18.T(c, params[i])

		paramarrs = append(paramarrs, param)
	}

	str := i18.T(c, key, paramarrs...)

	return NewResponse("SUCCESS", str, nil)
}

// ResponseSuccessWithData ResponseSuccessWithData
func ResponseSuccessWithData(c *gin.Context, data any) Response {
	i18 := i18n.Default(c)

	str := i18.T(c, "SUCCESS")

	return NewResponse("SUCCESS", str, data)
}

// ResponseSuccessWithMessageAndData ResponseSuccessWithMessageAndData
func ResponseSuccessWithMessageAndData(c *gin.Context, data any, key string, params ...string) Response {
	i18 := i18n.Default(c)

	paramarrs := make([]any, 0)

	for i := 0; i < len(params); i++ {
		param := i18.T(c, params[i])

		paramarrs = append(paramarrs, param)
	}

	str := i18.T(c, key, paramarrs...)

	return NewResponse("SUCCESS", str, data)
}

// ResponseFailure ResponseFailure
func ResponseFailure(c *gin.Context) Response {
	i18 := i18n.Default(c)

	str := i18.T(c, "FAILURE")

	return NewResponse("FAILURE", str, nil)
}

// ResponseFailureWithMessage ResponseFailureWithMessage
func ResponseFailureWithMessage(c *gin.Context, key string, params ...string) Response {
	i18 := i18n.Default(c)

	paramarrs := make([]any, 0)

	for i := 0; i < len(params); i++ {
		param := i18.T(c, params[i])

		paramarrs = append(paramarrs, param)
	}

	str := i18.T(c, key, paramarrs...)

	return NewResponse("FAILURE", str, nil)
}

// ResponseFailureWithData ResponseFailureWithData
func ResponseFailureWithData(c *gin.Context, data any) Response {
	i18 := i18n.Default(c)

	str := i18.T(c, "FAILURE")

	return NewResponse("FAILURE", str, data)
}

// ResponseFailureWithMessageAndData ResponseFailureWithMessageAndData
func ResponseFailureWithMessageAndData(c *gin.Context, data any, key string, params ...string) Response {
	i18 := i18n.Default(c)

	paramarrs := make([]any, 0)

	for i := 0; i < len(params); i++ {
		param := i18.T(c, params[i])

		paramarrs = append(paramarrs, param)
	}

	str := i18.T(c, key, paramarrs...)

	return NewResponse("FAILURE", str, data)
}
