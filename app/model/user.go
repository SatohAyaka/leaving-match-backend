package model

type User struct {
	BackendUserId   int64 `gorm:"column:backend_user_id"`
	StayWatchUserId int64 `gorm:"column:staywatch_user_id"`
	SlackUserId     int64 `gorm:"column:slack_user_id"`
}

func (User) TableName() string { return "User_Data" }
