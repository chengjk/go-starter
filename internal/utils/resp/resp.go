package resp

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-starter/internal/utils/resp/codes"
	"net/http"
	"reflect"
)

type Resp struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, NewResp(codes.OK, data))
}
func SuccessWith(ctx *gin.Context, code codes.Code, data interface{}) {
	ctx.JSON(http.StatusOK, NewResp(code, data))
}
func Failed(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusInternalServerError, NewResp(codes.InternalError, data))
}

func NewResp(code codes.Code, data interface{}) Resp {
	if reflect.TypeOf(data) == reflect.TypeOf(errors.New("")) {
		data = data.(error).Error()
	}
	return Resp{
		code.Value,
		code.Message,
		data,
	}
}
