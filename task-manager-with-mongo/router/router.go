package router

import (
	"github.com/Abzaek/GO-APIs/task-manager-with-mongo/controllers"
	"github.com/gin-gonic/gin"
)

func StartApp(control *controllers.Control) {
	router := gin.Default()

	router.GET("/tasks", control.GetTasks)
	router.GET("/tasks/:id", control.GetTasks)
	router.PUT("/tasks/:id", control.UpdateTask)
	router.DELETE("/tasks/:id", control.DeleteTask)
	router.POST("/tasks", control.PostTask)

	router.Run("localhost:3000")
}
