package model

type User struct {
	BackendUserId   int64  `gorm:"column:backend_user_id;primaryKey;autoIncrement"`
	StayWatchUserId int64  `gorm:"column:staywatch_user_id;uniqueIndex"`
	SlackUserId     string `gorm:"column:slack_user_id;uniqueIndex"`
	UserName        string `gorm:"column:user_name;uniqueIndex"`
}

func (User) TableName() string { return "User_Data" }
