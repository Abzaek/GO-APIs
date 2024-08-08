package main

import (
	"fmt"
	"log"

	"github.com/Abzaek/GO-APIs/task-manager-with-mongo/controllers"
	"github.com/Abzaek/GO-APIs/task-manager-with-mongo/router"
)

func main() {

	Controller := controllers.Control{}

	Controller.SetContext()
	Controller.SetOptions()

	err := Controller.Init("tasks")

	if err != nil {
		log.Fatal("error: ", err)
	}
	fmt.Println("db connected successfully")
	router.StartApp(&Controller)
}
