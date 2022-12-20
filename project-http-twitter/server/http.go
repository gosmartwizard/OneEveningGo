package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Server struct {
	Repository UserPayloadRepository
}

func (s Server) AddTweet(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	up := UserPayload{}

	if err := json.Unmarshal(body, &up); err != nil {
		log.Println("Failed to unmarshal payload:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("Tweet: `%s` from %s\n", up.Message, up.Location)

	id, err := s.Repository.AddUserPayload(up)
	if err != nil {
		log.Println("Failed to add user:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ti := TweetId{
		Id: id,
	}

	b, err := json.Marshal(ti)

	if err != nil {
		log.Println("Failed to unmarshal payload:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, string(b))
}

type tweetsList struct {
	Tweets []UserPayload `json:"tweets"`
}

func (s Server) ListTweets(w http.ResponseWriter, r *http.Request) {

	userPayloads, _ := s.Repository.GetUserPayloads()

	tweets := tweetsList{}

	tweets.Tweets = append(tweets.Tweets, userPayloads...)

	b, err := json.Marshal(tweets)

	if err != nil {
		log.Println("Failed to unmarshal payload:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, string(b))

}
