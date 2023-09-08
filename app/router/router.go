package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sjm1327605995/mycards_store/app/handler"
	"github.com/sjm1327605995/mycards_store/app/handler/middleware"
	_ "github.com/sjm1327605995/mycards_store/docs"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() (r *gin.Engine) {
	e := gin.New()
	e.Use(middleware.Cors())
	if viper.GetBool("API") {
		url := ginSwagger.URL("/swagger/doc.json")
		e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}

	api := e.Group("api")
	{
		api.POST("putDesks", handler.PutDesks)
		api.GET("getDesksById", handler.GetDesksById)
		api.GET("getDesksList", handler.GetDesksList)
		api.DELETE("delDesksById", handler.DelDesksById)

		//replay:=api.Group("replay")
		//{
		//	replay.POST("upload",)
		//	replay.GET()
		//}
	}

	return e
}
