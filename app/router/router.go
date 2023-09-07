package router

import (
	"github.com/gin-gonic/gin"
	"mycard_store/app/handler"
)

func Router() (r *gin.Engine) {
	e := gin.New()

	api := e.Group("api")
	{
		api.GET("getCardsById", handler.GetCardsById)
	}

	return e
}
