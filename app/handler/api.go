package handler

import (
	"github.com/gin-gonic/gin"
	"mycard_store/app/handler/resp"
)

func GetCardsById(c *gin.Context) {
	userId := c.Query("userId")
	if userId == "" {
		resp.Fail(c, "用户id不存在")
		return
	}

}
