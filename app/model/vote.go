package model

type Vote struct {
	VoteId        int64    `gorm:"column:vote_id;primaryKey;autoIncrement"`
	BusTimeId     int64    `gorm:"column:bustime_id"`
	BackendUserId int64    `gorm:"column:backend_user_id"`
	Previous      bool     `gorm:"column:previous"`
	Nearest       bool     `gorm:"column:nearest"`
	Next          bool     `gorm:"column:next"`
	CreatedDate   JSONTime `gorm:"column:created_date;type:datetime;autoCreateTime"`
}

func (Vote) TableName() string { return "Vote_Data" }
