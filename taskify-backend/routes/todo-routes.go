package routes

import (
	"taskify/controllers"

	"github.com/gin-gonic/gin"
)

func TodoRoutes(router *gin.Engine) {
	todoRouter := router.Group("/todos")
	todoRouter.GET("", controllers.GetAllTodos())
	todoRouter.POST("", controllers.AddTodo())
	todoRouter.GET("/:id", controllers.GetTodo())
	todoRouter.PUT("/:id", controllers.UpdateTodo())
	todoRouter.DELETE("/:id", controllers.DeleteTodo())
}
