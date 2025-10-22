package binding

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func AutoBind(ctx *gin.Context, v any) error {

	err := bindUri(ctx, v)
	if err != nil {
		return err
	}
	err = bindQuery(ctx, v)
	if err != nil {
		return err
	}
	err = ctx.ShouldBindHeader(v)
	if err != nil {
		return err
	}
	err = ctx.ShouldBind(v)
	if err != nil {
		return err
	}
	return sourceValidator.ValidateStruct(v)
}

func bindUri(c *gin.Context, obj any) error {
	m := make(map[string][]string, len(c.Params))
	for _, v := range c.Params {
		m[v.Key] = []string{v.Value}
	}
	return binding.MapFormWithTag(obj, m, "uri")

}
func bindQuery(c *gin.Context, obj any) error {
	m := c.Request.URL.Query()
	return binding.MapFormWithTag(obj, m, "query")
}
