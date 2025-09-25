package service

import (
	"SatohAyaka/leaving-match-backend/model"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func GetAllSlackUsers() ([]model.SlackUser, error) {
	apiURL := os.Getenv("SLACK_GET_USERS_API")
	if apiURL == "" {
		return nil, fmt.Errorf("SLACK_GET_USERS_API が設定されていません")
	}
	token := os.Getenv("SLACK_BOT_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("SLACK_BOT_TOKEN が設定されていません")
	}

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("リクエスト作成失敗: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("slack API呼び出し失敗: %w", err)
	}
	defer resp.Body.Close()

	var data struct {
		OK      bool `json:"ok"`
		Members []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Deleted bool   `json:"deleted"`
			IsBot   bool   `json:"is_bot"`
		} `json:"members"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("JSONパース失敗: %w", err)
	}

	if !data.OK {
		return nil, fmt.Errorf("slack API returned not ok")
	}

	users := []model.SlackUser{}
	for _, m := range data.Members {
		if m.Deleted || m.IsBot {
			continue
		}
		users = append(users, model.SlackUser{ID: m.ID, Name: m.Name})
	}

	return users, nil

}

func OpenDM(slackUserID string) (string, error) {
	apiURL := os.Getenv("SLACK_OPEN_DM_API")
	if apiURL == "" {
		return "", fmt.Errorf("SLACK_OPEN_DM_API が設定されていません")
	}
	token := os.Getenv("SLACK_BOT_TOKEN")

	payload := map[string]interface{}{"users": slackUserID}
	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("リクエスト作成失敗: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("slack API呼び出し失敗: %w", err)
	}
	defer resp.Body.Close()

	var data struct {
		OK      bool `json:"ok"`
		Channel struct {
			ID string `json:"id"`
		} `json:"channel"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", fmt.Errorf("レスポンスパース失敗: %w", err)
	}
	if !data.OK {
		return "", fmt.Errorf("slack API エラー")
	}

	return data.Channel.ID, nil
}
