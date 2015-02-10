package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Spider struct {
	SpiderList []string
}

func Spiders(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	now := time.Now().Format("2006-01-02 15:04:05")
	welcome := WelcomeInfo{
		"for crawl",
		"do not panic",
		now,
	}
	encoder, err := json.Marshal(welcome)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(encoder)
}