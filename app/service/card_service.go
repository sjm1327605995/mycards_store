package service

import (
	"mycard_store/app/database"
	"mycard_store/app/models"
)

var CardsService = new(Cards)

type Cards struct {
}

func (service *Cards) GetCardsById(decksId uint64) ([]string, error) {
	var decks models.Decks
	err := database.GetDB().Model(&decks).Where("id = ?", decksId).First(&decks).Error
	if err != nil {
		return nil, err
	}

	return
}
