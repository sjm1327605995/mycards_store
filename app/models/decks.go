package models

import "time"

type Decks struct {
	Id        int64     `json:"id,string" gorm:"autoIncrement:false"`
	Name      string    `json:"name"`
	Cards     Cards     `json:"cards" gorm:"type:text;serializer:json"`
	UserId    int64     `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	LastUseAt time.Time `json:"lastUseAt"`
}
type Cards struct {
	Main  []int64 `json:"main"`
	Side  []int64 `json:"side"`
	Extra []int64 `json:"extra"`
}
type DecksNames struct {
	Id   int64  `json:"id,string" gorm:"autoIncrement:false"`
	Name string `json:"name"`
}

func (d DecksNames) TableName() string {
	return "decks"
}
