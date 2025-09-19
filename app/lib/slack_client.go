package lib

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func SendDM(channelID, text string) {
	url := "https://slack.com/api/chat.postMessage"
	body := map[string]string{
		"channel": channelID,
		"text":    text,
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SLACK_BOT_TOKEN"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Slack送信エラー:", err)
		return
	}
	var respData map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&respData)

	if ok, exists := respData["ok"].(bool); !exists || !ok {
		log.Printf("Slack API error:%+v\n", respData)
	}

	defer resp.Body.Close()
}
