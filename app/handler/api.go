package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sjm1327605995/mycards_store/app/common/custom_error"
	"github.com/sjm1327605995/mycards_store/app/handler/resp"
	"github.com/sjm1327605995/mycards_store/app/models"
	"github.com/sjm1327605995/mycards_store/app/service"
	"github.com/spf13/cast"
)

func PutDesks(c *gin.Context) {
	var desks models.Decks
	err := c.ShouldBindJSON(&desks)
	if err != nil {
		resp.Fail(c, "数据接收错误")
		return
	}
	if desks.Name == "" {
		resp.Fail(c, "卡牌名称不能为空")
		return
	}
	if desks.UserId == 0 {
		resp.Fail(c, custom_error.LostUserIdErr.Error())
		return
	}
	err = service.CardsService.PutDesk(&desks)
	if err != nil {
		resp.Fail(c, err.Error())
		return
	}
	resp.Success(c, nil)
}

func GetDesksById(c *gin.Context) {
	desksId := c.Query("id")
	if desksId == "" {
		resp.Fail(c, "卡组id不存在")
		return
	}

	desks, err := service.CardsService.GetDesksById(cast.ToUint64(desksId))
	if err != nil {
		resp.Fail(c, err.Error())
	}
	resp.Success(c, desks)
}

func GetDesksList(c *gin.Context) {
	userId := c.Query("userId")
	if userId == "" {
		resp.Fail(c, "卡组id不存在")
		return
	}

	desks, err := service.CardsService.GetDesksList(cast.ToInt64(userId))
	if err != nil {
		resp.Fail(c, err.Error())
		return
	}
	resp.Success(c, desks)
}
