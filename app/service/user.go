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
