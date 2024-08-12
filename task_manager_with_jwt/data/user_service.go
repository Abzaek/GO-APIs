package data

import (
	"context"
	"errors"
	"fmt"

	"github.com/Abzaek/GO-APIs/task-manager-with-mongo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type UserCollection struct {
	clientOptions *options.ClientOptions
	ctx           context.Context
	User          *mongo.Collection
}

func (c *UserCollection) SetOptions() {
	c.clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
}

// sets the context
func (c *UserCollection) SetContext() {

	c.ctx = context.TODO()
}

// Initializes the database
func (c *UserCollection) Init(name string) error {

	var err error
	var dbClient *mongo.Client

	dbClient, err = mongo.Connect(c.ctx, c.clientOptions)

	c.User = dbClient.Database("manager").Collection(name)
	return err
}

func (c *UserCollection) Promote(userId string) error {
	filter := bson.D{{
		Key:   "id",
		Value: userId,
	}}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "role", Value: "admin"}}}}

	singleResult, err := c.User.UpdateOne(c.ctx, filter, update)

	if err != nil {
		return err
	}

	if singleResult.MatchedCount == 0 {
		return errors.New("the user doesn't exist")
	}

	return nil
}

func (c *UserCollection) Register(user *models.User) error {

	filter := bson.D{{
		Key:   "id",
		Value: user.ID,
	}}
	// fmt.Println("1")
	singleResult := c.User.FindOne(c.ctx, filter)

	if singleResult.Err() == nil {
		return errors.New("user already exists")
	}
	// fmt.Println(2)
	hashedPass, err1 := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err1 != nil {
		return err1
	}
	// fmt.Println(3)
	user.Password = string(hashedPass)

	byteUser, err := bson.Marshal(&user)

	if err != nil {
		return err
	}
	// fmt.Println(4)
	inserOneResult, err2 := c.User.InsertOne(c.ctx, byteUser)
	if err2 != nil {
		return err2
	}

	fmt.Println(inserOneResult.InsertedID)
	return nil
}

func (c *UserCollection) Login(user *models.User) error {
	filter := bson.D{{Key: "id", Value: user.ID}}

	singleResult := c.User.FindOne(c.ctx, filter)

	var dbUser *models.User

	err := singleResult.Decode(&dbUser)

	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))

	if dbUser.Role == "user" && user.Role == "admin" {
		return errors.New("cannot login as an admin")
	}

	if err != nil {
		return err
	}

	user.Role = dbUser.Role

	return nil
}

func (c *UserCollection) Get(id string) (*models.User, error) {
	var user models.User

	filter := bson.D{{Key: "id", Value: id}}

	singleResult := c.User.FindOne(c.ctx, filter)

	if err := singleResult.Decode(&user); err != nil {
		return &user, err
	}

	return &user, nil
}

func (c *UserCollection) GetAll() ([]*models.User, int, error) {
	filter := bson.D{{}}

	cur, err := c.User.Find(c.ctx, filter)

	if err != nil {
		return []*models.User{}, 0, err
	}

	var result []*models.User

	for cur.Next(c.ctx) {
		var elem models.User

		err := bson.Unmarshal(cur.Current, &elem)

		if err != nil {
			fmt.Println("here")
			return []*models.User{}, 0, err
		}

		result = append(result, &elem)
	}
	cur.Close(c.ctx)

	return result, len(result), nil
}

func (c *UserCollection) Delete(id string) error {
	filter := bson.D{{Key: "id", Value: id}}

	_, err := c.User.DeleteMany(c.ctx, filter)

	if err != nil {
		return err
	}

	return nil

}
