package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
)

func GetSession() *mgo.Session {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error occured while loading env")
	}
	connString := os.Getenv("MONGODB_URL")
	s,err := mgo.Dial(connString)
	if err != nil {
		panic(err)
	}
	return s
}