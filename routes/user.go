package routes

import (
	"project-go/handler"

	"github.com/gin-gonic/gin"
)

func UserRoutes(api *gin.RouterGroup, handler *handler.UserHandler) {
	api.POST("/register", handler.RegisterUser)
	api.POST("/login", handler.Login)
	api.POST("/check-email", handler.CheckAvailabilityEmail)
}
