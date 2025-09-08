package model

type Result struct {
	ResultId    int64    `gorm:"column:result_id;primaryKey;autoIncrement"`
	BusTimeId   int64    `gorm:"column:bustime_id;not null"`
	BusTime     JSONTime `gorm:"column:bus_time;not null;uniqueIndex"`
	Member      int64    `gorm:"column:member"`
	CreatedDate JSONTime `gorm:"column:created_date;type:datetime;autoCreateTime"`
}

func (Result) TableName() string { return "Result_Data" }
