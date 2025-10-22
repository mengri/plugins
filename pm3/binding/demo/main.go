package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mengri/plugins/pm3"
	"github.com/mengri/plugins/pm3/binding"
	"net/http"
)

type Request struct {
	Name      string `json:"url_name" uri:"name"  binding:"required"`
	Value     string `json:"url_value" uri:"value" binding:"required"`
	QueryName string `json:"query_name" query:"name"`
	Append    string `json:"append" form:"append" binding:"required"`
	Header    string `header:"user-agent" json:"header" `
}

func main() {
	engine := gin.New()
	engine.POST("/test/:name/:value", handle)
	engine.GET("/test/:name/:value", handle)
	err := engine.Run(":8080")
	if err != nil {
		panic(err)
	}
}
func handle(ctx *gin.Context) {
	v := new(Request)
	err := binding.AutoBind(ctx, &v)
	if err != nil {
		pm3.ResponseErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	pm3.Response(ctx, v)
}
