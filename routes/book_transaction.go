package routes

import (
	"project-go/handler"

	"github.com/gin-gonic/gin"
)

func BookTransactionRoutes(api *gin.RouterGroup, handler *handler.BookTransactionHandler) {
	api.GET("/", handler.Get)
	api.GET("/:id", handler.Find)
	api.POST("/store", handler.Create)
	api.PATCH("/update/:id", handler.Edit)
}
