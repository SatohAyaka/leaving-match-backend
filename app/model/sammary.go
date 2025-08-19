package model

import "time"

type Sammary struct {
	SammaryId    int64     `gorm:"column:sammary_id;primaryKey;autoIncrement"`
	PredictionId int64     `gorm:"column:prediction_id"`
	BusTimeId    int64     `gorm:"column:bustime_id"`
	ResultId     int64     `gorm:"column:result_id"`
	CreatedDate  time.Time `gorm:"column:created_date;type:datetime;autoCreateTime"`
}

func (Sammary) TableName() string { return "Sammary_Data" }
