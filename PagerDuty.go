package main

import (
	"encoding/json"
	"net/http"
	"log"
	"github.com/namsral/flag"
	"strings"
	"os"
	"fmt"

	webhook "github.com/haroldsphinx/pd-webhook-service"
)

const (
	IncidentTrigger       = "incident.trigger"       // IncidentTrigger is sent when an incident is newly created/triggered
	IncidentAcknowledge   = "incident.acknowledge"   // IncidentAcknowledge is sent when an incident is acknowledged by a user
	IncidentUnAcknowledge = "incident.unacknowledge" // IncidentUnacknowledge is sent when an incident is unacknowledged due to its acknowledgement timing out
	IncidentResolve       = "incident.resolve"       // IncidentResolve is sent when an incident has been resolved
	IncidentAssign        = "incident.assign"        // IncidentAssign is sent when an incident has been assigned to another user. Often occurs in concert with an acknowledge
	IncidentEscalate      = "incident.escalate"      // IncidentEscalate is sent when an incident has been escalated to another user in the same escalation chain
	IncidentDelegate      = "incident.delegate"      // IncidentDelegate is sent when an incident has been reassigned to another escalation policy
	IncidentAnnotate      = "incident.annotate"      // IncidentAnnotate is sent when a note is created on an incident.
)

var (
	flagSet       = flag.NewFlagSetWithEnvPrefix(os.Args[0], "GWP", 0)
	listenAddress = flagSet.String("listen", ":8080", "Listening port for the webhook service")
	secret        = flagSet.String("secret", "", "Secret of the webhook API")
	provider      = flagSet.String("provider", "pagerduty", "Provider Name")
	allowedPaths  = flagSet.String("allowedPaths", "", "Comma separated list of the paths we want to accept from the client")
)

func validateRequiredFlags() {
	isValid := true
	if len(strings.TrimSpace(*listenAddress)) == 0 {
		log.Println("Required flag 'listen address' not specified")
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
		log.Println("Payload", payload)
	}
}



func main() {
	http.HandleFunc("/", incidentHandler())
	http.ListenAndServe(":8080", nil)
}

