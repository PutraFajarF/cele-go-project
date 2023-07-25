package routes

import (
	"project-go/handler"

	"github.com/gin-gonic/gin"
)

func MasterBookRoutes(api *gin.RouterGroup, handler *handler.MasterBookHandler) {
	api.GET("/book", handler.Get)
	api.GET("/book/:id", handler.Find)
	api.POST("/book/store", handler.Create)
	api.PATCH("/book/update/:id", handler.Edit)
}
