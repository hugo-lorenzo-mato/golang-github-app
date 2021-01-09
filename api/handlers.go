package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
	"github.com/hugo-lorenzo-mato/golang-github-app/bot"
	"github.com/hugo-lorenzo-mato/golang-github-app/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type WebHookResponse struct {
	Processed bool `json:"processed"`
}

func WebHookRouter(path string) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.NewMiddleware())

	router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		context, _ := bot.FromContext(r.Context())
		// Just some examples
		switch p := context.Payload.(type) {
		case *github.IssuesEvent:
			fmt.Println(p.Issue.HTMLURL)
		case *github.NewPullRequest:
			fmt.Println(p.GetIssue())
		case *github.PullRequestReview:
			fmt.Println(p.PullRequestURL)
		case *github.IssueComment:
			fmt.Println(p.HTMLURL)
		default:
			log.Printf("Unknown event type: %s\n", github.WebHookType(r))
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		sendResponse(w)

	}).Methods("POST")

	return router
}

func sendResponse(w http.ResponseWriter) {
	resp := WebHookResponse{
		Processed: true,
	}
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Error(errors.New(fmt.Sprintf("Error in json.NewEncoder(w).Encode(resp) - %+v", err)))
	}
}
