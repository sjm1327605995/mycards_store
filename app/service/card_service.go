package service

import (
	"github.com/sjm1327605995/mycards_store/app/common/snow"
	"github.com/sjm1327605995/mycards_store/app/database"
	"github.com/sjm1327605995/mycards_store/app/models"
)

var CardsService = new(Cards)

type Cards struct {
}

func (service *Cards) GetDesksById(decksId uint64) (*models.Decks, error) {
	var decks models.Decks
	err := database.GetDB().Model(&decks).Where("id = ?", decksId).First(&decks).Error
	if err != nil {
		return nil, err
	}

	return &decks, nil
}

func (service *Cards) PutDesk(decks *models.Decks) (err error) {
	if decks.Id == 0 {
		decks.Id = snow.GenID()
	}
	err = database.GetDB().Save(&decks).Error
	if err != nil {
		return err
	}

	return nil
}
func (service *Cards) GetDesksList(userId int64) (list []models.DecksNames, err error) {
	list = make([]models.DecksNames, 0)
	err = database.GetDB().Select("id", "name").Where("user_id = ?", userId).Find(&list).Error
	if err != nil {
		return
	}
	return
}
