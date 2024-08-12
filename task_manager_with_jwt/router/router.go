package router

import (
	"github.com/Abzaek/GO-APIs/task-manager-with-mongo/controllers"
	"github.com/Abzaek/GO-APIs/task-manager-with-mongo/middleware"
	"github.com/gin-gonic/gin"
)

func StartApp(control *controllers.Control) {
	router := gin.Default()

	router.GET("/tasks", middleware.AuthMiddleware("user", control), control.GetTasks)
	router.GET("/tasks/:id", middleware.AuthMiddleware("user", control), control.GetTasks)
	router.PUT("/tasks/:id", middleware.AuthMiddleware("admin", control), control.UpdateTask)
	router.DELETE("/tasks/:id", middleware.AuthMiddleware("admin", control), control.DeleteTask)
	router.POST("/tasks", middleware.AuthMiddleware("admin", control), control.PostTask)
	router.POST("/register", control.RegisterUser)
	router.POST("/login", control.LoginUser)
	router.PUT("/promote/:id", middleware.AuthMiddleware("admin", control), control.PromotUser)
	router.DELETE("/users/:id", middleware.AuthMiddleware("admin", control), control.DeleteUser)
	router.GET("/getall", control.GetUsers)
	router.Run("localhost:3000")
}
