package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sjm1327605995/mycards_store/app/common/custom_error"
	"github.com/sjm1327605995/mycards_store/app/handler/resp"
	"github.com/sjm1327605995/mycards_store/app/models"
	"github.com/sjm1327605995/mycards_store/app/service"
	"github.com/spf13/cast"
)

// PutDesks
// @Summary 保存卡组
// @Schemes
// @Description 保存卡组，如果卡组Id没有则创建。有id则覆盖数据库这条记录
// @Tags 卡组
// @Param data body models.Decks true "卡组信息"
// @Accept json
// @Produce json
// @Success 200 {object} resp.SuccessResp
// @Router /api/putDesks [post]
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
	resp.Success(c, "ok")
}

// GetDesksById
// @Summary 查询卡组
// @Schemes
// @Description 查询卡组id，获取卡组信息
// @Tags 卡组
// @Param data  query string true "id"
// @Accept json
// @Produce json
// @Success 200 {object} resp.SuccessResp{data=models.Decks} "ok"
// @Router /api/getDesksById [get]
func GetDesksById(c *gin.Context) {
	desksId := c.Query("id")
	if desksId == "" {
		resp.Fail(c, "卡组id不存在")
		return
	}

	desks, err := service.CardsService.GetDesksById(cast.ToUint64(desksId))
	if err != nil {
		resp.Fail(c, err.Error())
		return
	}
	resp.Success(c, desks)
}

// DelDesksById
// @Summary 删除卡组
// @Schemes
// @Description 根据卡组id，删除卡组
// @Tags 卡组
// @Param data  query string true "id"
// @Accept json
// @Produce json
// @Success 200 {object} resp.SuccessResp "ok"
// @Router /api/delDesksById [delete]
func DelDesksById(c *gin.Context) {
	desksId := c.Query("id")
	if desksId == "" {
		resp.Fail(c, "卡组id不存在")
		return
	}

	err := service.CardsService.DelDesksById(cast.ToInt64(desksId))
	if err != nil {
		resp.Fail(c, err.Error())
		return
	}
	resp.Success(c, "ok")
}

// GetDesksList
// @Summary 查询卡组列表
// @Schemes
// @Description 查询用户的卡组列表
// @Tags 卡组
// @Param data  query string true "userId"
// @Accept json
// @Produce json
// @Success 200 {object} resp.SuccessResp{data=[]models.Decks}
// @Router /api/getDesksList [get]
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
