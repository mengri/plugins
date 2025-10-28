package pm3

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mengri/plugins/pm3/binding"
	"log"
	"net/http"
)

// 生成无输入的api
func CreateApiSimple[OUTPUT any](method string, path string, handler func(ctx *gin.Context) (OUTPUT, error)) Api {
	if handler == nil {
		log.Fatalf("handler is nil for %s %s", method, path)
		return nil
	}
	h := GenSimpleHandler(handler)

	return &formApi{
		method:  method,
		path:    path,
		handler: h,
	}
}

func GenSimpleHandler[OUTPUT any](handler func(ctx *gin.Context) (OUTPUT, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		output, err := handler(ctx)
		if err != nil {
			ResponseErr(ctx, http.StatusInternalServerError, fmt.Sprintf("server error:%s", err.Error()))
			return
		}

		Response(ctx, output)
	}
}

// 生成无响应的api( 仅响应状态)
func CreateApiNone[INPUT any](method string, path string, handler func(ctx *gin.Context, input *INPUT) error) Api {
	if handler == nil {
		log.Fatalf("handler is nil for %s %s", method, path)
		return nil
	}
	h := GenNoneHandler(handler)
	return &formApi{
		method:  method,
		path:    path,
		handler: h,
	}
}

func GenNoneHandler[INPUT any](handler func(ctx *gin.Context, input *INPUT) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input = new(INPUT)
		if err := binding.AutoBind(ctx, &input); err != nil {
			ResponseErr(ctx, http.StatusBadRequest, fmt.Sprintf("invald request:%s", err.Error()))
			return
		}
		err := handler(ctx, input)
		if err != nil {
			ResponseErr(ctx, http.StatusInternalServerError, fmt.Sprintf("server error:%s", err.Error()))
			return
		}

		ResponseOk(ctx)
	}
}
