package controller

import (
	"SatohAyaka/leaving-match-backend/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRecommendedHandler(c *gin.Context) {
	recommendedTime, err := ParseQueryToTime(c.Query("time"), "recommended")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get recommendedTime"})
		return
	}

	membersQuery := c.QueryArray("member")
	backendIds := make([]int64, 0, len(membersQuery))

	for _, m := range membersQuery {
		staywatchId, err := strconv.ParseInt(m, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid member id"})
			return
		}

		backendId, err := StayWatchIdToBackendId(staywatchId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to map staywatchId to backendId"})
			return
		}

		backendIds = append(backendIds, backendId)
	}

	recommendedService := service.RecommendedService{}
	response, err := recommendedService.CreateRecommended(recommendedTime, backendIds)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get recommended data"})
		return
	}

	c.JSON(http.StatusOK, response)

}

func GetLatestRecommendedStatusHandler(c *gin.Context) {
	recommendedService := service.RecommendedService{}
	recommended, err := recommendedService.GetLatestRecommended()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get recommended data"})
		return
	}
	c.JSON(http.StatusOK, recommended.Status)
}

func GetLatestRecommendedMembersHandler(c *gin.Context) {
	recommendedService := service.RecommendedService{}
	recommended, err := recommendedService.GetLatestRecommended()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get recommended data"})
		return
	}
	c.JSON(http.StatusOK, recommended.MemberIds)
}
