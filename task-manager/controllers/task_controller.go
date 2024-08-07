package controllers

import (
	"net/http"

	"github.com/Abzaek/GO-APIs/task-manager/data"
	"github.com/Abzaek/GO-APIs/task-manager/models"
	"github.com/gin-gonic/gin"
)

var Task []models.Task

func UpdateTask(c *gin.Context) {

	id := c.Param("id")

	if err := c.ShouldBindJSON(data.Update(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.String(http.StatusCreated, "Updated successfully")

}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	isDeleted := data.Delete(id)

	if isDeleted {
		c.String(http.StatusOK, "Successfully deleted")
	} else {
		c.String(http.StatusBadRequest, "Input is not available")
	}
}

func PostTask(c *gin.Context) {
	var newTask models.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		data.Post(&newTask)
		c.JSON(http.StatusAccepted, gin.H{
			"message": "Successfully created",
		})
	}

}

func GetTasks(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(http.StatusOK, data.Tasks)
	}

	for _, task := range data.Tasks {
		if task.ID == id {
			c.IndentedJSON(http.StatusOK, task)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, "Resource Not found")
}
