package controller

import (
	"SatohAyaka/leaving-match-backend/lib"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
