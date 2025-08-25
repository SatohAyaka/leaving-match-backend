package model

import "time"

type BusTime struct {
	BusTimeId     int64     `gorm:"column:bustime_id;primaryKey;autoIncrement"`
	RecommendedId int64     `gorm:"column:recommended_id"`
	PreviousTime  time.Time `gorm:"column:previous_time"`
	NearestTime   time.Time `gorm:"column:nearest_time"`
	NextTime      time.Time `gorm:"column:next_time"`
	CreatedDate   time.Time `gorm:"column:created_date;type:datetime;autoCreateTime"`
	EndTime       time.Time `gorm:"column:end_date"`
}

func (BusTime) TableName() string { return "BusTime_Data" }
