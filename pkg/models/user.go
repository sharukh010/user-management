package models

import (
	"context"
	"fmt"
	"time"

	"github.com/sharukh010/user-managment/pkg/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var userCollection *mongo.Collection

type User struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name   string             `json:"name" bson:"name"`
	Gender string             `json:"gender" bson:"gender"`
	Age    int                `json:"age" bson:"age"`
}

func init() {
	client := config.GetSession()
	userCollection = client.Database("mongo-golang").Collection("users")
	fmt.Println("✅ Successfully connected to database and initialized collection")
}

func (u *User) CreateUser() error {
	u.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := userCollection.InsertOne(ctx, u)
	return err
}

func GetUserById(userId primitive.ObjectID) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User
	err := userCollection.FindOne(ctx, bson.M{"_id": userId}).Decode(&user)
	return &user, err
}

func DeleteUser(userId primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := userCollection.DeleteOne(ctx, bson.M{"_id": userId})
	return err
}
