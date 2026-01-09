package routes

import (
	"bioskop-api/controllers"

	"github.com/gin-gonic/gin"
)

func BioskopRoutes(router *gin.Engine) {
	router.GET("/bioskops", controllers.GetBioskopAll)
	router.GET("/bioskop/:id", controllers.GetBioskopByID)
	router.PUT("/bioskop/:id", controllers.UpdateBioskop)
	router.DELETE("/bioskop/:id/delete", controllers.DeleteBioskop)
	router.POST("/bioskop", controllers.CreateBioskop)
}
