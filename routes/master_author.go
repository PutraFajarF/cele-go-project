package routes

import (
	"project-go/handler"

	"github.com/gin-gonic/gin"
)

func MasterAuthorRoutes(api *gin.RouterGroup, handler *handler.MasterAuthorHandler) {
	api.GET("/author", handler.Get)
	api.GET("/author/:id", handler.Find)
	api.POST("/author/store", handler.Create)
	api.PATCH("/author/update/:id", handler.Edit)
}
