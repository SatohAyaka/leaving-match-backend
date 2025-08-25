package service

import (
	"SatohAyaka/leaving-match-backend/lib"
	"SatohAyaka/leaving-match-backend/model"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type UserService struct{}

func (UserService) CreateUser(staywatchUserId int64, slackUserId int64, userName string) (int64, error) {
	user := model.User{
		StayWatchUserId: staywatchUserId,
		SlackUserId:     slackUserId,
		UserName:        userName,
	}

	if err := lib.DB.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) || strings.Contains(err.Error(), "Duplicate entry") {
			return 0, errors.New("user already exists")
		}
		return 0, err
	}
	return user.BackendUserId, nil
}

func (UserService) GetUser(backendUserId int64, staywatchUserId int64, slackUserId int64, userName string) ([]model.User, error) {
	var users []model.User
	db := lib.DB

	// クエリパラメータがある場合のみ絞り込み
	if backendUserId != 0 {
		db = db.Where("backend_user_id = ?", backendUserId)
	}
	if staywatchUserId != 0 {
		db = db.Where("staywatch_user_id = ?", staywatchUserId)
	}
	if slackUserId != 0 {
		db = db.Where("slack_user_id = ?", slackUserId)
	}
	if userName != "" {
		db = db.Where("user_name = ?", userName)
	}

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
