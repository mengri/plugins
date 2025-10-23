package plugins

import (
	"github.com/gin-gonic/gin"
	"github.com/mengri/plugins/ignore"
	"github.com/mengri/plugins/pm3"
	"net/http"
)

type plugin struct {
	driverCreate func() pm3.Driver
	name         string
	fronted      *FrontendConfig
}
type FrontendConfig struct {
	Name    string    `json:"name"`
	Remote  string    `json:"remote"`
	Preload bool      `json:"preload"`
	Router  []*Router `json:"router"`
}
type Router struct {
	Path   string `json:"path"`
	Router string `json:"router"`
}

type frontendPlugins []*FrontendConfig

func (f frontendPlugins) Name() string {
	return "_system.config"
}

func (f frontendPlugins) APis() []pm3.Api {
	ignore.IgnorePath("*", http.MethodGet, "/api/v1/system/config")
	return []pm3.Api{
		pm3.CreateApiSimple(http.MethodGet, "/api/v1/system/config", func(ctx *gin.Context) (any, error) {
			return map[string]any{"plugins": f}, nil
		}),
	}
}
