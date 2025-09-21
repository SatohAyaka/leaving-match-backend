package controller

import (
	"SatohAyaka/leaving-match-backend/lib"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SendDMHandler(c *gin.Context) {
	membersQuery := c.QueryArray("member")
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
	})
}
