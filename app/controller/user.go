package controller

import (
	"SatohAyaka/leaving-match-backend/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(c *gin.Context) {
	var staywatchUserId int64
	var staywatchIdPtr *int64
	var err error

	staywatchUserQuery := c.Query("staywatch")
	if staywatchUserQuery != "" {
		staywatchUserId, err = strconv.ParseInt(staywatchUserQuery, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get staywatchUserId"})
			return
		}
		staywatchIdPtr = &staywatchUserId
	}

	slackUserQuery := c.Query("slack")
	var slackIdPtr *string
	if slackUserQuery != "" {
		slackIdPtr = &slackUserQuery
	}

	userNameQuery := c.Query("name")
	var userNamePtr *string
	if userNameQuery != "" {
		userNamePtr = &userNameQuery
	}

	if staywatchUserQuery == "" && slackUserQuery == "" && userNameQuery == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "no query"})
		return
	}

	userService := service.UserService{}
	backendUserId, err := userService.CreateUser(staywatchIdPtr, slackIdPtr, userNamePtr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get backendUserId"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"backend_user_id": backendUserId})
}

func UpdateUserHandler(c *gin.Context) {
	backendUserPass := c.Param("backendUserId")
	backendUserId, err := strconv.ParseInt(backendUserPass, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get BackendUserId"})
		return
	}

	var staywatchUserId int64
	var staywatchIdPtr *int64
	staywatchUserQuery := c.Query("staywatch")
	if staywatchUserQuery != "" {
		staywatchUserId, err = strconv.ParseInt(staywatchUserQuery, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get staywatchUserId"})
			return
		}
		staywatchIdPtr = &staywatchUserId
	}

	slackUserQuery := c.Query("slack")
	var slackIdPtr *string
	if slackUserQuery != "" {
		slackIdPtr = &slackUserQuery
	}

	userNameQuery := c.Query("name")
	var userNamePtr *string
	if userNameQuery != "" {
		userNamePtr = &userNameQuery
	}

	if staywatchUserQuery == "" && slackUserQuery == "" && userNameQuery == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "no query"})
		return
	}

	userService := service.UserService{}
	userData, err := userService.UpdateUser(backendUserId, staywatchIdPtr, slackIdPtr, userNamePtr)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to update user data"})
		return
	}
	c.JSON(http.StatusOK, userData)
}

func GetUserHandler(c *gin.Context) {
	var backendUserId, staywatchUserId int64
	var err error

	backendUserQuery := c.Query("backend")
	if backendUserQuery != "" {
		staywatchUserId, err = strconv.ParseInt(backendUserQuery, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get staywatchUserId"})
			return
		}
	}

	staywatchUserQuery := c.Query("staywatch")
	var staywatchIdPtr *int64
	if staywatchUserQuery != "" {
		staywatchUserId, err = strconv.ParseInt(staywatchUserQuery, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get staywatchUserId"})
			return
		}
		staywatchIdPtr = &staywatchUserId
	}

	slackUserQuery := c.Query("slack")
	var slackIdPtr *string
	if slackUserQuery != "" {
		slackIdPtr = &slackUserQuery
	}

	userNameQuery := c.Query("name")
	var userNamePtr *string
	if userNameQuery != "" {
		userNamePtr = &userNameQuery
	}
	if staywatchUserQuery == "" && slackUserQuery == "" && userNameQuery == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "no query"})
		return
	}

	userService := service.UserService{}
	userData, err := userService.GetUser(backendUserId, staywatchIdPtr, slackIdPtr, userNamePtr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get user data"})
		return
	}
	c.JSON(http.StatusOK, userData)
}

func StayWatchIdToBackendId(staywatchId int64) (int64, error) {
	var staywatchIdPtr = &staywatchId

	userService := service.UserService{}
	userData, err := userService.GetUser(0, staywatchIdPtr, nil, nil)
	if err != nil {
		return 0, err
	}
	if len(userData) == 0 {
		return 0, fmt.Errorf("no backendId found for staywatchId=%d", staywatchId)
	}
	return userData[0].BackendUserId, nil
}

func SlackIdToBackendId(slackId string) (int64, error) {
	var slackIdPtr = &slackId
	userService := service.UserService{}
	userData, err := userService.GetUser(0, nil, slackIdPtr, nil)
	if err != nil {
		return 0, err
	}
	if len(userData) == 0 {
		return 0, fmt.Errorf("no backendId found for slackId=%s", slackId)
	}
	return userData[0].BackendUserId, nil
}
