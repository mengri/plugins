package pm3

import (
	"fmt"
	"github.com/mengri/plugins/pm3/binding"
	"net/http"

	"github.com/gin-gonic/gin"
	"log"
)

func CreateApiStandard[INPUT any, OUTPUT any](method string, path string, handler func(*gin.Context, *INPUT) (*OUTPUT, error)) Api {
	if handler == nil {

		log.Fatalf("handler is nil for %s %s", method, path)
		return nil
	}
	h := GenStandardHandler(handler)

	return &formApi{
		method:  method,
		path:    path,
		handler: h,
	}
}

func GenStandardHandler[INPUT any, OUTPUT any](handler func(*gin.Context, *INPUT) (*OUTPUT, error)) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var input = new(INPUT)
		if err := binding.AutoBind(ctx, &input); err != nil {
			ResponseErr(ctx, http.StatusBadRequest, fmt.Sprintf("invald request:%s", err.Error()))
			return
		}

		output, err := handler(ctx, input)
		if err != nil {
			ResponseErr(ctx, http.StatusInternalServerError, fmt.Sprintf("server error:%s", err.Error()))
			return
		}

		Response(ctx, output)
	}
}
func CreateApiStandardWidthPage[INPUT any, OUTPUT any](method string, path string, handler func(ctx *gin.Context, page int, pageSize int, input *INPUT) (data []*OUTPUT, total int64, err error)) Api {
	if handler == nil {
		log.Fatalf("handler is nil for %s %s", method, path)
		return nil
	}
	h := GenStandardWidthPageHandler(handler)
	return &formApi{
		method:  method,
		path:    path,
		handler: h,
	}
}

type pageInput struct {
	Page int `form:"page" binding:"required,default=1"`
	Size int `form:"page_size" binding:"required,default=15"`
}

func GenStandardWidthPageHandler[INPUT any, OUTPUT any](handler func(ctx *gin.Context, page int, pageSize int, input *INPUT) (data []*OUTPUT, total int64, err error)) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var pi = new(pageInput)
		_ = binding.AutoBind(ctx, pi)
		input := new(INPUT)
		if err := binding.AutoBind(ctx, &input); err != nil {
			ResponseErr(ctx, http.StatusBadRequest, fmt.Sprintf("invald request:%s", err.Error()))
		}
		output, total, err := handler(ctx, pi.Page, pi.Size, input)
		if err != nil {
			return
		}
		ResponsePage(ctx, output, total, pi.Size, pi.Page)

	}
}
