package router

import (
	"github.com/Abzaek/GO-APIs/task-manager/controllers"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	router := gin.Default()

	router.GET("/tasks", controllers.GetTasks)
	router.GET("/tasks/:id", controllers.GetTasks)
	router.PUT("/tasks/:id", controllers.UpdateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)
	router.POST("/tasks", controllers.PostTask)

	router.Run("localhost:3000")
}
