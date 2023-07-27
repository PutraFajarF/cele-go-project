package routes

import (
	"project-go/auth"
	"project-go/handler"
	"project-go/middleware"
	"project-go/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BookTransactionRoutes(api *gin.RouterGroup, handler *handler.BookTransactionHandler, db *gorm.DB, authService auth.Service, userService user.Service) {
	api.GET("/", handler.Get)
	api.GET("/:id", handler.Find)
	api.POST("/store", middleware.AuthMiddleware(authService, userService), handler.Create)
	api.PATCH("/update/:id", middleware.AuthMiddleware(authService, userService), handler.Edit)
}
