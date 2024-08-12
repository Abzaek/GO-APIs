package data

import "github.com/Abzaek/GO-APIs/task-manager-with-mongo/models"

type Collections interface {
	Init(name string)
	Update(id string) *models.Task
	Delete(id string) bool
	Post(task *models.Task) *models.Task
	Register(user *models.User) error
	Login(user *models.User) error
	Promote(id int) error
}
