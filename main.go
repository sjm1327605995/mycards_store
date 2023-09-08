package main

import (
	"github.com/sjm1327605995/mycards_store/app/common/config"
	"github.com/sjm1327605995/mycards_store/app/common/snow"
	"github.com/sjm1327605995/mycards_store/app/log"

	"github.com/sjm1327605995/mycards_store/app/database"
	"github.com/sjm1327605995/mycards_store/app/router"
)

func main() {
	config.InitConfig()
	log.InitLog()
	snow.Init()
	database.InitDB()
	err := router.Router().Run(":8080")
	if err != nil {
		panic(err)
		return
	}
}
