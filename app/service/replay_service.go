package service

import (
	"errors"
	"github.com/sjm1327605995/mycards_store/app/common/custom_error"
	"github.com/sjm1327605995/mycards_store/app/common/snow"
	"github.com/sjm1327605995/mycards_store/app/common/storage"
	"github.com/sjm1327605995/mycards_store/app/database"
	"github.com/sjm1327605995/mycards_store/app/models"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"io"
	"time"
)

var ReplayService = new(Replay)

type Replay struct {
}

func (r Replay) PutReply(userId int64, name string, reader io.Reader) error {
	id := snow.GenID()
	strId := cast.ToString(id)
	err := storage.Upload(strId+".yrp", reader)
	if err != nil {
		zap.S().Info("上传失败: ", err.Error())
		return custom_error.UploadReplayErr
	}
	data := models.Replay{
		Id:        id,
		Name:      name,
		UserId:    userId,
		CreatedAt: time.Now(),
	}
	err = database.GetDB().Create(&data).Error
	if err != nil {
		zap.S().Info("保存录像失败: ", err.Error())
		go func() {
			path := strId + ".yrp"
			err := storage.Delete(path)
			if err != nil {
				zap.S().Error("录像存储删除失败: ", path)
			}
		}()
		return custom_error.UploadReplayErr
	}

	return nil
}

func (r Replay) DeleteReply(id string) error {
	session := database.GetDB().Delete(models.Replay{Id: cast.ToInt64(id)})
	err := session.Error
	if err != nil {
		zap.S().Error("录像存储删除失败: ", err.Error())
		return custom_error.DeleteErr
	}
	if session.RowsAffected > 0 {
		go func() {
			path := id + ".yrp"
			err := storage.Delete(path)
			if err != nil {
				zap.S().Error("录像存储删除失败: ", path, err)
			}
		}()
	}
	return nil
}
func (r Replay) GetReply(id string, writer io.Writer) error {

	err := storage.Download(id+".yrp", writer)
	if err != nil {
		zap.S().Error("获取录像失败: ", err.Error())
		return errors.New("获取录像失败")
	}
	return nil
}
func (r Replay) ReplyList(userId int64) ([]models.Replay, int64, error) {
	var (
		list  = make([]models.Replay, 0, 10)
		total int64
	)
	err := database.GetDB().Model(&models.Replay{}).Where("user_id = ?", userId).Count(&total).Order("created_at DESC").Find(&list).Error
	if err != nil {
		zap.S().Error("获取用户录像失败: ", err.Error())
		return nil, 0, errors.New("获取录像列表失败")
	}
	return list, total, nil
}
