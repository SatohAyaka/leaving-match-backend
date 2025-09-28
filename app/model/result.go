package model

import "time"

type Result struct {
	ResultId    int64     `gorm:"column:result_id;primaryKey;autoIncrement;type:BIGINT"`
	BusTimeId   int64     `gorm:"column:bustime_id;not null;type:BIGINT"`
	BusTime     time.Time `gorm:"column:bus_time;not null"`
	Member      int64     `gorm:"column:member"`
	CreatedDate time.Time `gorm:"column:created_date;type:datetime;autoCreateTime"`
}

func (Result) TableName() string { return "Result_Data" }
