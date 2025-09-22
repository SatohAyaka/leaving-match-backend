package controller

import (
	"SatohAyaka/leaving-match-backend/lib"
	"SatohAyaka/leaving-match-backend/model"
	"SatohAyaka/leaving-match-backend/service"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type SlackEvent struct {
	Type      string `json:"type"`
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
	Event     struct {
		Type    string `json:"type"`
		User    string `json:"user"`
		Text    string `json:"text"`
		Channel string `json:"channel"`
	} `json:"event"`
}

func SendDMHandler(c *gin.Context) {
	membersQuery := c.QueryArray("member")
	busTimesQuery := c.QueryArray("bustime")

	busMessage := "乗れそうなバスを選択してください\n"
	for i, minbustime := range busTimesQuery {
		bustime, err := ParseQueryToTime(minbustime, "DM bustime")
		if err != nil {
			log.Println("failed to parse bustime:", minbustime, "err:", err)
			continue
		}
		busMessage += fmt.Sprintf("%d. %s\n", i+1, bustime)
	}

	channels := make([]string, 0, len(membersQuery))
	for _, m := range membersQuery {
		staywatchId, err := strconv.ParseInt(m, 10, 64)
		if err != nil {
			log.Println("invalid member id:", m, err)
			continue
		}

		channelId, err := StayWatchIdToChannelId(staywatchId)
		if err != nil {
			log.Println("failed to map staywatchId:", staywatchId, "err:", err)
			continue
		}

		channels = append(channels, channelId)
	}

	for _, channel := range channels {
		lib.SendDM(channel, "乗れそうなバスを選択してください")
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"channels": channels,
		"message":  busMessage,
	})
}

func SlackEventHandler(c *gin.Context) {
	var payload map[string]interface{}
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	// SlackのURL検証 (最初だけ飛んでくる)
	if t, ok := payload["type"].(string); ok && t == "url_verification" {
		c.JSON(http.StatusOK, gin.H{"challenge": payload["challenge"]})
		return
	}

	event := payload["event"].(map[string]interface{})
	slackUserId := event["user"].(string) // SlackのユーザーID
	text := event["text"].(string)        // 投稿内容 (例 "1")

	// 数字を投票内容にマッピング
	previous, nearest, next := false, false, false

	normalized := strings.ReplaceAll(text, ",", " ")
	for _, token := range strings.Fields(normalized) { // スペース区切りで分割
		switch strings.TrimSpace(token) {
		case "1":
			previous = true
		case "2":
			nearest = true
		case "3":
			next = true
		}
	}

	if !previous && !nearest && !next {
		c.JSON(http.StatusOK, gin.H{"message": "ignored"})
		return
	}

	busTimeService := service.BusTimeService{}
	lastBusTime, err := busTimeService.GetLatestBusTime()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get latest bustime"})
		return
	}

	// 投票済みチェック
	resultService := service.ResultService{}
	lastResult, err := resultService.GetResult(lastBusTime.BusTimeId)
	if err == nil && lastResult != (model.Result{}) {
		c.JSON(http.StatusForbidden, gin.H{"message": "voting closed"})
		return
	}

	backendId, err := SlackIdToBackendId(slackUserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	voteService := service.VoteService{}
	if err := voteService.CreateVote(lastBusTime.BusTimeId, backendId, previous, nearest, next); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create vote"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "vote created"})
}
