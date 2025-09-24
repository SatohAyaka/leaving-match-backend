package model

type User struct {
	BackendUserId   int64   `gorm:"column:backend_user_id;primaryKey;autoIncrement"`
	StayWatchUserId *int64  `gorm:"column:staywatch_user_id;unique"`
	SlackUserId     *string `gorm:"column:slack_user_id;unique"`
	ChannelId       *string `gorm:"column:channel_id;unique"`
	UserName        *string `gorm:"column:user_name;unique"`
}

type StayWatchUser struct {
	StayWatchUserId int64  `json:"id"`
	Name            string `json:"name"`
	Tags            []struct {
		TagID   int    `json:"id"`
		TagName string `json:"name"`
	} `json:"tags"`
}

func (User) TableName() string { return "User_Data" }
