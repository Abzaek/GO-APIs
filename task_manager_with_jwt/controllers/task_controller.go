package controllers

import (
	"log"
	"net/http"

	"github.com/Abzaek/GO-APIs/task-manager-with-mongo/data"
	// "github.com/Abzaek/GO-APIs/task-manager-with-mongo/middleware"
	// "github.com/Abzaek/GO-APIs/task-manager-with-mongo/middleware"
	"github.com/Abzaek/GO-APIs/task-manager-with-mongo/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JwtKey = []byte("abzaeko")

type Control struct {
	data.TaskCollection
	User data.UserCollection
}

func (ctrl *Control) UpdateTask(ctx *gin.Context) {
	var obj models.Task

	if err := ctx.ShouldBindBodyWithJSON(&obj); err != nil {
		log.Fatal("ctrl 1")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return

	}

	err := ctrl.Update(&obj)

	if err != nil {

		ctx.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	ctx.String(http.StatusNoContent, "Successfully Updated")
}

func (ctrl *Control) DeleteTask(ctx *gin.Context) {
	var id string = ctx.Param("id")

	err := ctrl.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"message": "successfully deleted"})

}

func (ctrl *Control) PostTask(ctx *gin.Context) {
	var obj models.Task

	if err := ctx.ShouldBindJSON(&obj); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}

	err := ctrl.Post(&obj)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"message": "POST request successful"})
}

func (ctrl *Control) GetTasks(ctx *gin.Context) {
	var obj []models.Task
	var err error

	obj, err = ctrl.Get(ctx)

	if err != nil {

		switch err.Error() {
		case "404":
			ctx.JSON(http.StatusNotFound, gin.H{"message": "RESOURCE NOT FOUND"})
		case "500":
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		}

		return
	}

	ctx.JSON(http.StatusCreated, obj)
}

func (ctrl *Control) PromotToAdmin(id *int) {

}

func (ctrl *Control) PromotUser(c *gin.Context) {
	id := c.Param("id")

	err := ctrl.User.Promote(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user promoted"})
}

func (ctrl *Control) LoginUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.User.Login(&user)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	// claims := middleware.CustomClaim{
	// 	ID:   user.ID,
	// 	Role: user.Role,
	// }
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
	})

	stringToken, err1 := token.SignedString(JwtKey)

	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err1})
		return
	}

	user.Token = stringToken

	c.JSON(http.StatusCreated, user)
}

func (ctrl *Control) RegisterUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, num, err := ctrl.User.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if num == 0 {
		user.Role = "admin"
	}

	// claims := middleware.CustomClaim{
	// 	ID:   user.ID,
	// 	Role: user.Role,
	// }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
	})

	stringToken, err1 := token.SignedString(JwtKey)

	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err1.Error()})
		return
	}

	err = ctrl.User.Register(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Token = stringToken

	c.JSON(http.StatusCreated, user)

}

func (ctrl *Control) GetSingleUser(id string) *models.User {

	user, err := ctrl.User.Get(id)
	if err != nil {
		return &models.User{}
	}
	return user
}

func (ctrl *Control) GetUsers(c *gin.Context) {
	res, _, err := ctrl.User.GetAll()
	if err != nil {
		return
	}
	var arr []models.User

	for _, r := range res {
		arr = append(arr, *r)
	}
	c.JSON(http.StatusOK, arr)

}

func (ctrl *Control) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := ctrl.User.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusNoContent, "deleted successfully")
}
