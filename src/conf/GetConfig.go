package conf

type configData struct {
	path string
}

func NewConfigData(path string) *configData {
	return &configData{path: path}
}

func (c *configData) GetFrontServer() string {
	config := getConfig(c.path)
	return config.FrontServer
}
func (c *configData) GetApi() string {
	config := getConfig(c.path)
	return config.Api
}

func (c *configData) GetOnlineServers() []Server {
	config := getConfig(c.path)
	return config.Servers

}
