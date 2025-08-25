package model

import "time"

type Recommended struct {
	RecommendedId   int64     `gorm:"column:recommended_id;primaryKey;autoIncrement"`
	RecommendedTime time.Time `gorm:"column:recommended_time"`
	CreatedDate     time.Time `gorm:"column:created_date;type:datetime;autoCreateTime"`
}

func (Recommended) TableName() string { return "Recommended_Data" }
