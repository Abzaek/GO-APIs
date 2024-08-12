package data

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Abzaek/GO-APIs/task-manager-with-mongo/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskCollection struct {
	clientOptions *options.ClientOptions
	ctx           context.Context
	Task          *mongo.Collection
}

// Sets the client options
func (c *TaskCollection) SetOptions() {
	c.clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
}

// sets the context
func (c *TaskCollection) SetContext() {

	c.ctx = context.TODO()
}

// Initializes the database
func (c *TaskCollection) Init(name string) error {

	var err error
	var dbClient *mongo.Client

	dbClient, err = mongo.Connect(c.ctx, c.clientOptions)

	c.Task = dbClient.Database("manager").Collection(name)
	return err
}

//Update the database

func (c *TaskCollection) Update(obj *models.Task) error {

	filter := bson.M{"id": obj.ID}

	newDocument, err2 := bson.Marshal(*obj)

	if err2 != nil {
		log.Fatal("srvc 1")
		return err2
	}

	// update := bson.M{"$set": newDocument}

	_, err := c.Task.ReplaceOne(c.ctx, filter, newDocument)

	return err
}

// Delete by Id
func (c *TaskCollection) Delete(id string) error {

	filter := bson.M{"id": id}

	_, err := c.Task.DeleteOne(c.ctx, filter)

	return err
}

// Post a document
func (c *TaskCollection) Post(obj *models.Task) error {
	newDocument, err1 := bson.Marshal(obj)

	if err1 != nil {
		return err1
	}

	_, err := c.Task.InsertOne(c.ctx, newDocument)

	return err
}

//Get document or documents

func (c *TaskCollection) Get(ctx *gin.Context) ([]models.Task, error) {
	fmt.Println("I am called")
	id := ctx.Param("id")
	var result []models.Task

	if id == "" {
		cur, err := c.Task.Find(c.ctx, bson.D{{}}, options.Find())

		if err != nil {

			return []models.Task{}, errors.New("404")
		}

		for cur.Next(c.ctx) {
			var elem models.Task
			err = bson.Unmarshal(cur.Current, &elem)

			if err != nil {
				return []models.Task{}, errors.New("500")
			}

			result = append(result, elem)
		}
		cur.Close(c.ctx)

		return result, nil
	}

	filter := bson.M{"id": id}

	cur, err := c.Task.Find(c.ctx, filter)

	if err != nil {

		return []models.Task{}, errors.New("404")
	}

	for cur.Next(c.ctx) {
		var elem models.Task

		err := bson.Unmarshal(cur.Current, &elem)

		if err != nil {

			return []models.Task{}, errors.New("500")
		}

		result = append(result, elem)

	}
	cur.Close(c.ctx)
	if len(result) == 0 {
		return result, errors.New("404")
	}
	return result, nil
}
