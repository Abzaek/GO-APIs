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

	Controller.User.SetContext()
	Controller.User.SetOptions()

	err := Controller.Init("tasks")

	if err != nil {
		log.Fatal("error: ", err)
	}

	err = Controller.Init("users")

	if err != nil {
		log.Fatal("error: ", err)
	}
	err = Controller.User.Init("users")

	if err != nil {
		log.Fatal("error: ", err)
	}

	fmt.Println("db connected successfully")
	router.StartApp(&Controller)
}
