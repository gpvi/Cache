package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
)

var default_port = 80
var default_api = false
var default_ip = "localhost"

type Server struct {
	Port int    `yaml:"port,omitempty" default:"80"`
	API  bool   `yaml:"api,omitempty" default:"false"`
	IP   string `yaml:"ip,omitempty" default:"localhost"`
}

type Config struct {
	Peers       []string `yaml:"peers"`
	FrontServer string   `yaml:"front-server"`
	Api         string   `yaml:"api"`
	Servers     []Server `yaml:"online-servers"`
}

func setDefaults(servers []Server) {
	for i := range servers {
		// 如果未提供API值，则设置为默认值false
		if servers[i].API == false {
			servers[i].API = default_api
		}
		if servers[i].IP == "" {
			servers[i].IP = default_ip
		}
		if servers[i].Port == 0 {
			servers[i].Port = default_port
		}
	}
}
func getConfig(configPath string) *Config {
	// 读取配置文件
	println(os.Getwd())
	v := viper.New()
	v.SetConfigType("yaml")

	file, err := os.Open(configPath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	//解码
	decoder := yaml.NewDecoder(file)
	var config Config
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println(err)
	}
	// 设置默认值
	setDefaults(config.Servers)

	return &config

}
