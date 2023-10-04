package routes

import (
	"go-paslon/controllers"

	"github.com/gin-gonic/gin"
)

func PaslonRoutes(router *gin.Engine) {
    routes := router.Group("/paslons")
    {
        routes.POST("", controllers.CreatePaslon)
        routes.GET("", controllers.GetPaslons)
        routes.GET("/:id", controllers.GetPaslonByID)
        routes.PATCH("/:id", controllers.UpdatePaslon)
        routes.DELETE("/:id", controllers.DeletePaslon)
    }
}

