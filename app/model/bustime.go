package model

type BusTime struct {
	BusTimeId     int64    `gorm:"column:bustime_id;primaryKey;autoIncrement"`
	RecommendedId int64    `gorm:"column:recommended_id;not null;uniqueIndex"`
	PreviousTime  JSONTime `gorm:"column:previous_time;"`
	NearestTime   JSONTime `gorm:"column:nearest_time"`
	NextTime      JSONTime `gorm:"column:next_time"`
	CreatedDate   JSONTime `gorm:"column:created_date;type:datetime;autoCreateTime"`
	EndTime       JSONTime `gorm:"column:end_date;type:datetime"`
}

func (BusTime) TableName() string { return "BusTime_Data" }
