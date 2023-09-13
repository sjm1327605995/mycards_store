package models

import "time"

type Replay struct {
	Id        int64     `json:"id,string"`                     //录像id（这里数据库不采用主键自增，使用雪花算法产生不重复的64位主键。前端使用string传输避免精度丢失问题）
	Name      string    `json:"name" gorm:"type:varchar(255)"` //录像名称
	UserId    int64     `json:"userId"`                        //用户Id     (这里可能涉及一个对战玩家的信息，现在暂时未考虑)
	CreatedAt time.Time `json:"createdAt"`                     //创建时间
}
