package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

func Configdirect() {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath("conf")

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
	}

	// 读取配置项
	peers := viper.GetStringSlice("peers")
	frontServer := viper.GetString("front-server")
	defaultReplicas := viper.GetInt("http.defaultReplicas")
	defaultBasePath := viper.GetString("http.defaultBasePath")
	api := viper.GetString("api")
	mysqlPassword := viper.GetString("mysql.password")
	mysqlUser := viper.GetString("mysql.user")
	mysqlDatabase := viper.GetString("mysql.database")
	// 打印配置项
	fmt.Println("Peers:", peers)
	fmt.Println("Front Server:", frontServer)
	fmt.Println("Default Replicas:", defaultReplicas)
	fmt.Println("Default Base Path:", defaultBasePath)
	fmt.Println("API:", api)
	fmt.Println("MySQL Password:", mysqlPassword)
	fmt.Println("MySQL User:", mysqlUser)
	fmt.Println("MySQL Database:", mysqlDatabase)
}
