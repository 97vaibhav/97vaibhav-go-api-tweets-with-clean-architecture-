// this file helps to set up our struct tweet ,

package models

import (
	"go-api-tweets-with-clean-architecture/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Tweet struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
	User  string `json:"user"`
}

// database connection steps
// make struct
// do initialize ---> connect,automigrate

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Tweet{})
}

//now all function below talk to database
func (b *Tweet) CreateTweet() *Tweet {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllTweets() []Tweet {
	var Tweets []Tweet
	db.Find(&Tweets)
	return Tweets
}

func GetTweetById(Id int64) (*Tweet, *gorm.DB) {
	var getTweet Tweet
	db := db.Where("ID=?", Id).Find(&getTweet)
	return &getTweet, db
}

func DeleteTweet(ID int64) Tweet {
	var tweet Tweet
	db.Where("ID=?", ID).Delete(tweet)
	return tweet
}
