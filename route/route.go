package route

import (
	"monit/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/fetch", controllers.Fetch)
		api.POST("/update", controllers.Update)
	}

	return r
}
