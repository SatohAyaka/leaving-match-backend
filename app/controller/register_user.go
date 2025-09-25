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

func ConnectUserData() error {
	users, err := service.GetAllSlackUsers()
	if err != nil {
		return err
	}

	userService := service.UserService{}

	for _, u := range users {
		slackID := u.ID
		userName := u.Name

		// DB から backendUserId を取得
		dbUsers, err := userService.GetUser(0, nil, nil, nil, &userName)
		if err != nil || len(dbUsers) == 0 {
			log.Printf("DBにユーザが存在しない: %s", userName)
			continue
		}
		backendUserID := dbUsers[0].BackendUserId

		// DM チャンネルを開設（通知はされない）
		channelID, err := service.OpenDM(slackID)
		if err != nil {
			log.Printf("DM開設失敗: %s, err: %v", userName, err)
			continue
		}

		// DB 更新
		_, err = userService.UpdateUser(backendUserID, nil, &slackID, &channelID, nil)
		if err != nil {
			log.Printf("ユーザ更新失敗: %s, err: %v", userName, err)
			continue
		}

		log.Printf("SlackID と ChannelID を紐づけ: %s", userName)
	}

	return nil
}
