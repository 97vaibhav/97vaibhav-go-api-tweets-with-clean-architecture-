package controllers

import (
	"encoding/json"
	"fmt"
	"go-api-tweets-with-clean-architecture/pkg/models"
	"go-api-tweets-with-clean-architecture/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewTweet models.Tweet //creating new tweet of type tweet from model file

func GetTweet(w http.ResponseWriter, r *http.Request) {
	newTweets := models.GetAllTweets()
	ans, _ := json.Marshal(newTweets)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ans)
}

func GetTweetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tweetId := vars["tweetId"]
	ID, err := strconv.ParseInt(tweetId, 0, 0)
	if err != nil {
		fmt.Println("You have error check while parsing")
	}
	TweetDetails, _ := models.GetTweetById(ID)
	ans, _ := json.Marshal(TweetDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ans)
}

func CreateTweet(w http.ResponseWriter, r *http.Request) {
	CreateTweet := &models.Tweet{}
	utils.ParseBody(r, CreateTweet)
	b := CreateTweet.CreateTweet()
	ans, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(ans)
}

func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tweetId := vars["tweetId"]
	ID, err := strconv.ParseInt(tweetId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	tweet := models.DeleteTweet(ID)
	ans, _ := json.Marshal(tweet)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ans)
}

func UpdateTweet(w http.ResponseWriter, r *http.Request) {
	var updateTweet = &models.Tweet{}
	utils.ParseBody(r, updateTweet)
	vars := mux.Vars(r)
	tweetId := vars["tweetId"]
	ID, err := strconv.ParseInt(tweetId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	TweetDetails, db := models.GetTweetById(ID)
	if updateTweet.Title != "" {
		TweetDetails.Title = updateTweet.Title
	}
	if updateTweet.Body != "" {
		TweetDetails.Body = updateTweet.Body
	}
	if updateTweet.User != "" {
		TweetDetails.User = updateTweet.User
	}
	db.Save(&TweetDetails)
	ans, _ := json.Marshal(TweetDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ans)
}
