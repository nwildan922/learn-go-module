package model

import "time"

type Counter struct {
	ID        uint      `gorm:"column:ID;primaryKey"`
	Counter   int32     `gorm:"column:Counter;not null"`
	Timestamp time.Time `gorm:"column:Timestamp;not null"`
	CreatedAt time.Time `gorm:"column:CreatedAt"`
	AppId     string    `gorm:"column:AppID;not null"`
}

func (Counter) TableName() string {
	return "counters"
}
