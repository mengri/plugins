package plugins

import "encoding/json"

var (
	installList []string
)

func SetInstall(pluginName ...string) {
	installList = pluginName
}

func ReadFrontendConfig(data []byte) *FrontendConfig {
	if len(data) == 0 {
		return nil
	}
	c := &FrontendConfig{}
	err := json.Unmarshal(data, c)
	if err != nil {
		return nil
	}
	return c

}
