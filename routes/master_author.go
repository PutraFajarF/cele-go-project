package routes

import (
	"project-go/auth"
	"project-go/handler"
	"project-go/middleware"
	"project-go/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MasterAuthorRoutes(api *gin.RouterGroup, handler *handler.MasterAuthorHandler, db *gorm.DB, authService auth.Service, userService user.Service) {
	api.GET("/author", handler.Get)
	api.GET("/author/:id", handler.Find)
	api.POST("/author/store", middleware.AuthMiddleware(authService, userService), handler.Create)
	api.PATCH("/author/update/:id", middleware.AuthMiddleware(authService, userService), handler.Edit)
}
