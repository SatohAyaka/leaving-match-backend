package service

import (
	"SatohAyaka/leaving-match-backend/lib"
	"SatohAyaka/leaving-match-backend/model"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"gorm.io/gorm"
)

type UserService struct{}

func (UserService) CreateUser(staywatchUserId *int64, slackUserId *string, channelId *string, userName *string) (int64, error) {
	if slackUserId != nil && *slackUserId == "" {
		slackUserId = nil
	}
	if channelId != nil && *channelId == "" {
		channelId = nil
	}
	if userName != nil && *userName == "" {
		userName = nil
	}

	user := model.User{
		StayWatchUserId: staywatchUserId,
		SlackUserId:     slackUserId,
		ChannelId:       channelId,
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

func (UserService) UpdateUser(backendUserId int64, staywatchUserId *int64, slackUserId *string, channelId *string, userName *string) (model.User, error) {
	var user model.User
	if err := lib.DB.Where("backend_user_id = ?", backendUserId).First(&user).Error; err != nil {
		return model.User{}, err
	}
	if staywatchUserId != nil {
		user.StayWatchUserId = staywatchUserId
	}
	if slackUserId != nil && *slackUserId != "" {
		user.SlackUserId = slackUserId
	}
	if channelId != nil && *channelId != "" {
		user.ChannelId = channelId
	}
	if userName != nil && *userName != "" {
		user.UserName = userName
	}
	if err := lib.DB.Save(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (UserService) GetUser(backendUserId int64, staywatchUserId *int64, slackUserId *string, channelId *string, userName *string) ([]model.User, error) {
	var users []model.User
	db := lib.DB

	// クエリパラメータがある場合のみ絞り込み
	if backendUserId != 0 {
		db = db.Where("backend_user_id = ?", backendUserId)
	}
	if staywatchUserId != nil {
		db = db.Where("staywatch_user_id = ?", staywatchUserId)
	}
	if slackUserId != nil {
		db = db.Where("slack_user_id = ?", slackUserId)
	}
	if channelId != nil {
		db = db.Where("channel_id = ?", channelId)
	}
	if userName != nil {
		db = db.Where("user_name = ?", userName)
	}

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (UserService) GetAllUsers() ([]model.StayWatchUser, error) {
	apiURL := os.Getenv("StayWatch_API")
	apiKey := os.Getenv("API_KEY")

	if apiURL == "" || apiKey == "" {
		return nil, fmt.Errorf("STAY_WATCH_API または API_KEY が設定されていません")
	}

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("リクエスト作成失敗: %w", err)
	}

	req.Header.Set("X-API-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("外部API呼び出し失敗: %w", err)
	}
	defer response.Body.Close()
	log.Println("ユーザ登録API呼び出し成功")

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("レスポンス読み込み失敗: %w", err)
	}

	var staywatchResponse []model.StayWatchUser
	if err := json.Unmarshal(body, &staywatchResponse); err != nil {
		return nil, fmt.Errorf("JSONパース失敗: %w", err)
	}

	return staywatchResponse, nil
}
