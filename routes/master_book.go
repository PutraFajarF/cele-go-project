package routes

import (
	"project-go/auth"
	"project-go/handler"
	"project-go/middleware"
	"project-go/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MasterBookRoutes(api *gin.RouterGroup, handler *handler.MasterBookHandler, db *gorm.DB, authService auth.Service, userService user.Service) {
	api.GET("/book", handler.Get)
	api.GET("/book/:id", handler.Find)
	api.POST("/book/store", middleware.AuthMiddleware(authService, userService), handler.Create)
	api.PATCH("/book/update/:id", middleware.AuthMiddleware(authService, userService), handler.Edit)
}
