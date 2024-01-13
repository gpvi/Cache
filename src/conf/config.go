package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
)

var default_port = 8080
var default_api = false
var default_ip = "localhost"

type Server struct {
	Port int    `yaml:"port" default:"80"`
	API  bool   `yaml:"api" default:"false"`
	IP   string `yaml:"ip" default:"localhost"`
}

type Config struct {
	FrontServer string   `yaml:"front-server" default:"http://localhost:9999"`
	Api         string   `yaml:"api"`
	Servers     []Server `yaml:"online-servers"`
}

func (c *Config) setDefaults() {
	servers := c.Servers
	var count_num int = 0
	for i := range servers {
		// 如果未提供API值，则设置为默认值false
		if servers[i].API == false {
			servers[i].API = default_api
		}
		if servers[i].IP == "" {
			servers[i].IP = default_ip
		}
		if servers[i].Port == 0 {
			servers[i].Port = default_port + count_num
		}
		count_num++
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
	config.setDefaults()
	return &config

}
