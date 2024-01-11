package conf

type configData struct {
	path string
}

func NewConfigData(path string) *configData {
	return &configData{path: path}
}

func (c *configData) GetPeers() []string {
	config := getConfig(c.path)
	return config.Peers
}
func (c *configData) GetFrontServer() string {
	config := getConfig(c.path)
	return config.FrontServer
}
func (c *configData) GetApi() string {
	config := getConfig(c.path)
	return config.Api
}

//func GetOnlineServers() []string {
//	config := getConfig()
//
//	for i := range config.Servers {
//
//	}
//}
