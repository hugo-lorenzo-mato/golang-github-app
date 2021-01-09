package main

import (
	"fmt"
	"github.com/hugo-lorenzo-mato/golang-github-app/api"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var (
	port = 8080
	path = "/api/v1/webhook"
)

func main() {
	router := api.WebHookRouter(path)
	log.Debugf("Server running at localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", port), router))
}
