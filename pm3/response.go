package pm3

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type _ResponseMessage struct {
	Code    int    `json:"code"`
	Success string `json:"success,omitempty"`
	Message string `json:"msg,omitempty"`
}

func ResponseErr(ctx *gin.Context, code int, err string) {
	resp := &_ResponseMessage{
		Code:    code,
		Success: "fail",
		Message: err,
	}

	ctx.JSON(http.StatusOK, resp)
}
func ResponseOk(ctx *gin.Context, msg ...string) {
	resp := &_ResponseMessage{
		Code:    0,
		Success: "success",
	}
	if len(msg) > 0 {
		resp.Message = msg[0]
	} else {
		resp.Message = "ok"
	}

	ctx.JSON(http.StatusOK, resp)
}

type _Response[T any] struct {
	Data     T      `json:"data,omitempty"`
	Code     int    `json:"code"`
	Success  string `json:"success,omitempty"`
	Message  string `json:"msg,omitempty"`
	Total    int64  `json:"total,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
	Page     int    `json:"page,omitempty"`
}

func Response[T any](ctx *gin.Context, data T, msg ...string) {
	resp := &_Response[T]{
		Code:    0,
		Success: "success",
		Data:    data,
	}
	if len(msg) > 0 {
		resp.Message = msg[0]
	}
	ctx.JSON(http.StatusOK, resp)
}
func ResponsePage[T any](ctx *gin.Context, data []*T, total int64, pageSize int, page int, msg ...string) {
	resp := &_Response[[]*T]{
		Data:     data,
		Code:     0,
		Success:  "success",
		Total:    total,
		PageSize: pageSize,
		Page:     page,
	}
	if len(msg) > 0 {
		resp.Message = msg[0]
	} else {
		resp.Message = "ok"
	}

	ctx.JSON(http.StatusOK, resp)
}
