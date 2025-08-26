package router

import (
	"SatohAyaka/leaving-match-backend/controller"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		// c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://leaving-match.vercel.app")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Next()
	})

	versionEngine := r.Group("/api")
	{
		versionEngine.POST("/user", controller.CreateUserHandler)
		versionEngine.PUT("/user/:backendUserId", controller.UpdateUserHandler)
		versionEngine.GET("/user", controller.GetUserHandler)

		versionEngine.POST("/recommended", controller.CreateRecommendedHandler)
		versionEngine.GET("/recommended/latest/status", controller.GetLatestRecommendedStatusHandler)
		versionEngine.GET("/recommended/latest/members", controller.GetLatestRecommendedMembersHandler)

		versionEngine.POST("/bustime/:recommendedId", controller.CreateBusTimeHandler)
		versionEngine.GET("/bustime/:bustimeId", controller.GetBusTimeByIdHandler)
		versionEngine.GET("/bustime/latest", controller.GetLatestBusTimeHandler)

		versionEngine.POST("/vote/:bustimeId/:staywatchUserId", controller.CreateVoteHandler)
		versionEngine.GET("/vote/:bustimeId", controller.GetVoteHandler)

		versionEngine.POST("/result/:bustimeId", controller.CreateResultHandler)
		versionEngine.GET("/result/:bustimeId", controller.GetResultHandler)
		versionEngine.GET("/result/latest", controller.GetLatestResultHandler)
	}

	r.Run(":8085")

}
