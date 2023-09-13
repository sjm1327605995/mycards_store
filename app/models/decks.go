package models

import "time"

type Decks struct {
	Id        int64     `json:"id,string" gorm:"autoIncrement:false"`       //卡组id （这里数据库不采用主键自增，使用雪花算法产生不重复的64位主键。前端使用string传输避免精度丢失问题）
	Name      string    `json:"name"      gorm:"type:varchar(255)"`         //牌组名称
	Cards     Cards     `json:"cards"     gorm:"type:text;serializer:json"` //卡组卡牌内容
	UserId    int64     `json:"userId"`                                     //用户id
	CreatedAt time.Time `json:"createdAt"`                                  //创建时间
	UpdatedAt time.Time `json:"updatedAt"`                                  //更新时间
	LastUseAt time.Time `json:"lastUseAt"`                                  //最后使用时间 (这里还为做记录处理，待讨论)
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
