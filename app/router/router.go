package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sjm1327605995/mycards_store/app/handler"
	"github.com/sjm1327605995/mycards_store/app/handler/middleware"
)

func Router() (r *gin.Engine) {
	e := gin.New()
	e.Use(middleware.Cors())
	api := e.Group("api")
	{
		api.POST("putDesks", handler.PutDesks)
		api.GET("getDesksById", handler.GetDesksById)
		api.GET("getDesksList", handler.GetDesksList)
	}

	return e
}
