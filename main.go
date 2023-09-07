package main

import (
	"fmt"
	"github.com/sjm1327605995/mycards_store/app/common/snow"

	"github.com/sjm1327605995/mycards_store/app/database"
	"github.com/sjm1327605995/mycards_store/app/router"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	snow.Init()
	database.InitDB()
	err = router.Router().Run(":8080")
	if err != nil {
		panic(err)
		return
	}
}
