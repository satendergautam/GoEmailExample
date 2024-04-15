package mappings

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func CreateUrlMappings() {
	gin.SetMode(gin.ReleaseMode)
	Router = gin.Default()

	Router.Use(controllers.Cors())
	// v1 of the API
	v1 := Router.Group("/api/v1")
	{
		v1.POST("/email", controllers.EmailAPI)

	}
}
