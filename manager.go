package plugins

import (
	"github.com/mengri/plugins/pm3"
	"github.com/mengri/plugins/server"
)

var (
	data = untyped.BuildUntyped[string, *plugin]()
)

func Register(name string, driver func() pm3.Driver, frontend ...*FrontendConfig) {
	p := &plugin{
		driverCreate: driver,
		name:         name,
	}
	if len(frontend) > 0 {
		p.fronted = frontend[0]

	}

	data.Set(name, p)
}

func Init() {
	infl := make(frontendPlugins, 0, len(installList))
	for _, n := range installList {
		p, has := data.Get(n)
		if has {
			pm3.Register(n, p.driverCreate())
			if p.fronted != nil {
				infl = append(infl, p.fronted)
			}
		}
	}
	server.AddSystemPlugin(infl)
}
