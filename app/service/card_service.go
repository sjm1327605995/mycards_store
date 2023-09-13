package service

import (
	"github.com/sjm1327605995/mycards_store/app/common/custom_error"
	"github.com/sjm1327605995/mycards_store/app/common/snow"
	"github.com/sjm1327605995/mycards_store/app/database"
	"github.com/sjm1327605995/mycards_store/app/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var CardsService = new(Cards)

type Cards struct {
}

func (service *Cards) GetDesksById(decksId uint64) (*models.Decks, error) {
	var decks models.Decks
	err := database.GetDB().Model(&decks).Where("id = ?", decksId).First(&decks).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, custom_error.NotFoundErr
		}
		zap.S().Error(err)
		return nil, custom_error.FoundErr
	}

	return &decks, nil
}

func (service *Cards) PutDesk(decks *models.Decks) (err error) {
	now := time.Now()
	cols := []string{"name", "cards", "user_id"}
	if decks.Id == 0 {
		decks.Id = snow.GenID()
		cols = append(cols, "created_at")
		decks.CreatedAt = now
	}
	cols = append(cols, "updated_at")
	decks.UpdatedAt = now
	err = database.GetDB().Select(cols).Save(&decks).Error
	if err != nil {
		zap.S().Error(err)
		return custom_error.SaveErr
	}

	return nil
}
func (service *Cards) DelDesksById(decksId int64) error {
	err := database.GetDB().Delete(&models.Decks{Id: decksId}).Error
	if err != nil {
		zap.S().Error(err)
		return custom_error.DeleteErr
	}
	return nil
}
func (service *Cards) GetDesksList(userId int64) (list []models.DecksNames, err error) {
	list = make([]models.DecksNames, 0)
	err = database.GetDB().Select("id", "name").Where("user_id = ?", userId).Find(&list).Error
	if err != nil {
		zap.S().Error(err)
		err = custom_error.FoundErr
		return
	}
	return
}
