package main

import (
	"encoding/json"
	"net/http"
	"log"

	webhook "github.com/haroldsphinx/pagerduty-webhook-service"
)

func main() {
	http.HandleFunc("/", incidentHandler())
	http.ListenAndServe(":8080", nil)
}

func incidentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var payload webhook.Payload
		err := decoder.Decode(&payload)
		if err != nil {
			return
		}
		log.Println(payload)
	}
}
