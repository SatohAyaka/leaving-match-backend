package model

import "time"

type Result struct {
	ResultId    int64     `gorm:"column:result_id;primaryKey;autoIncrement"`
	BusTimeId   int64     `gorm:"column:bustime_id"`
	BusTime     time.Time `gorm:"column:bus_time"`
	Member      int64     `gorm:"column:member"`
	CreatedDate time.Time `gorm:"column:created_date;type:datetime;autoCreateTime"`
}

func (Result) TableName() string { return "Result_Data" }
