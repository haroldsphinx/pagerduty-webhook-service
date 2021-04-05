package main

import (
	"encoding/json"
	"net/http"
	"log"
	"github.com/namsral/flag"
	"strings"
	webhook "github.com/pagerduty-webhook-service"
)

var (
	flagSet       = flag.NewFlagSetWithEnvPrefix(os.Args[0], "GWP", 0)
	listenAddress = flagSet.String("listen", ":8080", "Listening port for the webhook service")
	secret        = flagSet.String("secret", "", "Secret of the webhook API")
	provider      = flagSet.String("provider", "pager-duty")
	allowedPaths  = flagSet.String("allowedPaths", "", "Comma separated list of the paths we want to accept from the client")
)

func validateRequiredFlags() {
	isValid := true
	if len(strings.TrimSpace(*provider)) == 0 {
		log.Println("Required flag 'provider' not specified")
		isValid = false
	}

	if !isValid {
		fmt.Println("")
		flagSet.Usage()
		fmt.Println("")
		panic("See Flag Usage")
	}
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


func main() {
	flagSet.Parse(os.Args[1:])
	validateRequiredFlags()
	lowerProvider := strings.ToLower(*provider)

	//split , into an array
	allowedPathsArray := []string{}
	if len(*allowedPaths) > 0 {
		allowedPathsArray = strings.Split(*allowedPaths, ",")
	}

	log.Printf("Dapperlabs Webhook Service Started")
	p, err := proxy.NewProxy(*upstreamURL, allowedPathsArray, lowerProvider, *secret, ignoredUsersArray)
	if err != nil {
		log.Fatal(err)
	}


	http.HandleFunc("/", incidentHandler())
	http.ListenAndServe(":8080", nil)
}

