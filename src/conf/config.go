package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Peers       []string `yaml:"peers"`
	FrontServer string   `yaml:"front-server"`
	Api         string   `yaml:"api"`
}

func getConfig() *Config {
	v := viper.New()
	v.SetConfigType("yaml")
	file, err := os.Open("../src/conf/config.yaml")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	var config Config
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println(err)
	}

	return &config
	// 输出数组内容

	//for _, peer := range config.Peers {
	//	fmt.Println(peer)
	//}

}

// todo : viper读取配置
// todo: mysql连接
