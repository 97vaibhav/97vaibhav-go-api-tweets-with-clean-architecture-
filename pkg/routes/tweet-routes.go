//THIS PACKAGE IS USED BECAUSE IT WILL RETURN A ROUTER OR YOU CAN SAY DECLARE A ROUTER

package routes

import (
	"go-api-tweets-with-clean-architecture/pkg/controllers"

	"github.com/gorilla/mux"
)

var TweetRoutes = func(r *mux.Router) {
	r.HandleFunc("/tweet/", controllers.CreateTweet).Methods("POST")
	r.HandleFunc("/tweet/", controllers.GetTweet).Methods("GET")
	r.HandleFunc("/tweet/{tweetId}", controllers.GetTweetById).Methods("GET")
	r.HandleFunc("/tweet/{tweetId}", controllers.UpdateTweet).Methods("PUT")
	r.HandleFunc("/tweet/{tweetId}", controllers.DeleteTweet).Methods("DELETE")
}
