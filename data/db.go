package data

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB is the connection to the database
var DB *gorm.DB

func init() {
	var err error
	// set up DB connection and then attempt to connect 5 times over 25 seconds
	connectionParams := "user=user password=password sslmode=disable host=db"
	for i := 0; i < 5; i++ {
		DB, err = gorm.Open("postgres", connectionParams) // gorm checks Ping on Open
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Println("connection to DB failed, aborting...")
		log.Fatal(err)
	}

	// create table if it does not exist
	if !DB.HasTable(&Exercise{}) {
		DB.CreateTable(&Exercise{})
	}
	if !DB.HasTable(&Section{}) {
		DB.CreateTable(&Section{})
	}
	if !DB.HasTable(&Tag{}) {
		DB.CreateTable(&Tag{})
	}
	if !DB.HasTable(&Trial{}) {
		DB.CreateTable(&Trial{})
	}
}
