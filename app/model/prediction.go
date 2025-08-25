package model

import "time"

type Prediction struct {
	PredictionId   int64     `gorm:"column:prediction_id"`
	UserId         int64     `gorm:"column:user_id"`
	PredictionTime time.Time `gorm:"column:prediction_time"`
	CreatedDate    time.Time `gorm:"column:created_date;type:datetime;autoCreateTime"`
}

func (Prediction) TableName() string { return "Prediction_Data" }
