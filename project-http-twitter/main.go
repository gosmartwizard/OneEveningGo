package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"twitter/server"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

func main() {
	// Your solution goes here. Good luck!

	s := server.Server{
		Repository: &server.UserInMemoryRepository{},
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Use()

	r.With(httprate.LimitByIP(10, 1*time.Minute)).Post("/tweets", s.AddTweet)

	r.Get("/tweets", s.ListTweets)

	go spamTweets()

	http.ListenAndServe(":8080", r)
}

func spamTweets() {
	for {
		// Prepare payload
		addTweetPayload := server.UserPayload{
			Message:  "ass",
			Location: "ass",
		}
		// Marshal payload
		marshaledPayload, err := json.Marshal(addTweetPayload)

		// Send HTTP POST request
		url := "http://localhost:8080/tweets"

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(marshaledPayload))
		if err != nil {
			return
		}
		// Close response body
		defer resp.Body.Close()

		// (Optionally read and print the response)

	}
}
