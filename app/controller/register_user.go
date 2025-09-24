package controller

import (
	"SatohAyaka/leaving-match-backend/lib"
	"SatohAyaka/leaving-match-backend/service"
	"log"
	"os"
	"sync"
)

var (
	onceRegister sync.Once
)

func RegisterUserOnce() error {
	var registerErr error
	onceRegister.Do(func() {
		userService := service.UserService{}
		response, err := userService.GetAllUsers()
		if err != nil {
			registerErr = err
			return
		}
		for _, v := range response {
			staywatchUserId := v.StayWatchUserId
			userName := v.Name
			slackUserId := ""

			_, err := userService.CreateUser(&staywatchUserId, &slackUserId, &userName)
			if err != nil {
				if err.Error() == "user already exists" {
					continue
				}
				registerErr = err
				return
			}
		}
		channelID := os.Getenv("ADMIN_CHANNEL_ID")
		lib.SendDM(channelID, "初回ユーザ登録処理を実行しました")
		log.Println("ユーザ登録APIを一度だけ実行")
	})
	return registerErr
}
