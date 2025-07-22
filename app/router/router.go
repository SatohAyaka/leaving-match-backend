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
		versionEngine.POST("/bustime", controller.CreateBusTimeHandler)
		versionEngine.GET("/bustime", controller.GetBusTimeHandler)

		versionEngine.GET("/result", controller.GetResultHandler)
	}

	r.Run(":8085")

}
