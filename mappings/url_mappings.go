package mappings

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func CreateUrlMappings() {
	gin.SetMode(gin.ReleaseMode)
	Router = gin.Default()

	// Middleware to enable CORS
	Router.Use(func(c *gin.Context) {
		// List of allowed origins
		allowedOrigins := []string{
			"http://localhost:3000",
			"https://marketing-leaders.de/",
			"http://marketing-leaders.de/",
			"*",
		}

		origin := c.GetHeader("Origin")
		for _, allowedOrigin := range allowedOrigins {
			if allowedOrigin == origin {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
				c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
				c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

				// Handle preflight requests
				if c.Request.Method == "OPTIONS" {
					c.AbortWithStatus(204)
					return
				}
			}
		}

		c.Next()
	})

	//	Router.Use(controllers.Cors())
	// v1 of the API
	v1 := Router.Group("/api/v1")
	{
		v1.POST("/email", controllers.EmailAPI)

	}
}
