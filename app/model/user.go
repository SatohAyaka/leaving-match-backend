package model

type User struct {
	BackendUserId   int64   `gorm:"column:backend_user_id;primaryKey;autoIncrement"`
	StayWatchUserId *int64  `gorm:"column:staywatch_user_id;unique"`
	SlackUserId     *string `gorm:"column:slack_user_id;unique"`
	ChannelId       *string `gorm:"column:channel_id;unique"`
	UserName        *string `gorm:"column:user_name;unique"`
}

func (User) TableName() string { return "User_Data" }
