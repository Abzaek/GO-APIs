package controllers

import (
	"log"
	"net/http"

	"github.com/Abzaek/GO-APIs/task-manager-with-mongo/data"
	"github.com/Abzaek/GO-APIs/task-manager-with-mongo/models"
	"github.com/gin-gonic/gin"
)

type Control struct {
	data.Collection
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
