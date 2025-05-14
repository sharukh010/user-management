package models

import (
	"github.com/sharukh010/user-managment/pkg/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Collection

type User struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Gender string `json:"gender" bson:"gender"`
	Age int `json:"age" bson:"age"`
}
func init(){
	session := config.GetSession()
	db = session.DB("mongo-golang").C("users")

}
func (u *User)CreateUser() *User{
	u.Id = bson.NewObjectId()
	db.Insert(u)
	return u
}

func GetUserById(userId bson.ObjectId) (*User,error){
	var user User 
	err := db.FindId(userId).One(&user)
	return &user,err
}

func DeleteUser(userId bson.ObjectId) error{
	err := db.Remove(userId)
	return err
}