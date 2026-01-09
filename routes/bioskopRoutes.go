package routes

import (
	"bioskop-api/controllers"

	"github.com/gin-gonic/gin"
)

func BioskopRoutes(router *gin.Engine) {
	router.POST("/bioskop", controllers.CreateBioskop)
}
