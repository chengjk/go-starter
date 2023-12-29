package resp

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-starter/internal/pkg/resp/codes"
	"net/http"
	"reflect"
)

type Resp struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	TraceId string      `json:"traceId,omitempty"`
}

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, newResp(codes.OK, data))
	ctx.Abort()
}
func SuccessWith(ctx *gin.Context, code codes.BizCode, data interface{}) {
	ctx.JSON(http.StatusOK, newResp(code, data))
	ctx.Abort()
}

func Error(ctx *gin.Context, err error) {
	var bizErr codes.BizErr
	var resp Resp
	if ok := errors.As(err, &bizErr); ok {
		resp = Resp{
			Code:    bizErr.Code.Code,
			Message: err.Error(),
			Data:    bizErr.Code.Name,
		}
	}
	resp = Resp{
		Code:    codes.InternalError.Code,
		Message: err.Error(),
		Data:    codes.InternalError.Name,
	}
	ctx.JSON(http.StatusOK, resp)
	ctx.Abort()
}

func newResp(code codes.BizCode, data interface{}) Resp {
	if reflect.TypeOf(data) == reflect.TypeOf(errors.New("")) {
		data = data.(error).Error()
	}
	return Resp{
		code.Code,
		code.Name,
		data,
		"",
	}
}
