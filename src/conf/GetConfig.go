package conf

func GetPeers() []string {
	config := getConfig()
	return config.Peers
}
func GetFrontServer() string {
	config := getConfig()
	return config.FrontServer
}
func GetApi() string {
	config := getConfig()
	return config.Api
}
