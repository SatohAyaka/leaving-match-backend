package controller

import (
	"SatohAyaka/leaving-match-backend/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(c *gin.Context) {
	var staywatchUserId, slackUserId int64
	var err error

	staywatchUserQuery := c.Query("staywatch")
	if staywatchUserQuery != "" {
		staywatchUserId, err = strconv.ParseInt(staywatchUserQuery, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	slackUserQuery := c.Query("slack")
	if slackUserQuery != "" {
		slackUserId, err = strconv.ParseInt(slackUserQuery, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	userName := c.Query("name")

	if staywatchUserQuery == "" && slackUserQuery == "" && userName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userService := service.UserService{}
	backendUserId, err := userService.CreateUser(staywatchUserId, slackUserId, userName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"backend_user_id": backendUserId})
}

func UpdateUserHandler(c *gin.Context) {}

func GetUserHandler(c *gin.Context) {
	var backendUserId, staywatchUserId, slackUserId int64
	var err error

	backendUserQuery := c.Query("backend")
	if backendUserQuery != "" {
		staywatchUserId, err = strconv.ParseInt(backendUserQuery, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	staywatchUserQuery := c.Query("staywatch")
	if staywatchUserQuery != "" {
		staywatchUserId, err = strconv.ParseInt(staywatchUserQuery, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	slackUserQuery := c.Query("slack")
	if slackUserQuery != "" {
		slackUserId, err = strconv.ParseInt(slackUserQuery, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	userName := c.Query("name")
	userService := service.UserService{}
	userData, err := userService.GetUser(backendUserId, staywatchUserId, slackUserId, userName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get user data"})
		return
	}
	c.JSON(http.StatusOK, userData)
}
