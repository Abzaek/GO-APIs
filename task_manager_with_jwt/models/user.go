package models

type User struct {
	ID       string `json:"id" bson:"id" validate:"required"`
	Role     string `json:"role" bson:"role" validate:"required"`
	Password string `json:"password" bson:"password" validate:"required"`
	Token    string `json:"token" bson:"token"`
}
