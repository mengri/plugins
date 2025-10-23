package pm3

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateApiSimple[OUTPUT any](method string, path string, handler func(ctx *gin.Context) (*OUTPUT, error)) Api {
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

func GenSimpleHandler[OUTPUT any](handler func(ctx *gin.Context) (*OUTPUT, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		output, err := handler(ctx)
		if err != nil {
			ResponseErr(ctx, http.StatusInternalServerError, fmt.Sprintf("server error:%s", err.Error()))
			return
		}

		Response(ctx, output)
	}
}
