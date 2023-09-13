package main

import (
	"github.com/sjm1327605995/mycards_store/app/common/config"
	"github.com/sjm1327605995/mycards_store/app/common/snow"
	"github.com/sjm1327605995/mycards_store/app/common/storage"
	"github.com/sjm1327605995/mycards_store/app/database"
	"github.com/sjm1327605995/mycards_store/app/log"
	"github.com/sjm1327605995/mycards_store/app/router"
	"net/http"
)

func main() {
	config.InitConfig()
	log.InitLog()
	snow.Init()
	database.InitDB()
	storage.InitStorageMedia()

	err := http.ListenAndServe(":8080", router.Router().Handler())
	if err != nil {
		panic(err)
		return
	}

}
