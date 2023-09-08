package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

//type Config struct {
//	Database Database `yaml:"database"`
//}
//type Database struct {
//	Type     string `yaml:"type"`
//	Port     int    `yaml:"port"`
//	Host     string `yaml:"host"`
//	User     string `yaml:"user"`
//	Dbname   string `yaml:"dbname"`
//	Password string `yaml:"password"`
//	Init     bool   `yaml:"init"`
//}
//
//var Conf *Config

func InitConfig() {
	viper.SetConfigFile("config.yaml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.SetDefault("db.type", "sqlite")

}
