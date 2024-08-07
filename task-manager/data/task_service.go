package data

import (
	"time"

	"github.com/Abzaek/GO-APIs/task-manager/models"
)

var Tasks = []models.Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func Update(id string) *models.Task {

	for i := 0; i < len(Tasks); i++ {
		if Tasks[i].ID == id {
			return &Tasks[i]
		}
	}
	newTask := &models.Task{}

	return newTask
}

func Delete(id string) bool {

	for i := 0; i < len(Tasks); i++ {
		if Tasks[i].ID == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			return true
		}
	}
	return false
}

func Post(task *models.Task) {
	Tasks = append(Tasks, *task)
}
