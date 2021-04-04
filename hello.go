package main

import (
	"fmt"
	"time"
)

type Payload struct {
	Messages []Message `json:"messages"`
}

// https://developer.pagerduty.com/docs/webhooks/v2-overview/
type Message struct {
	ID         string     `json:"id"`
	Event      string     `json:"event"`
	CreatedOn  string     `json:"created_on"`
	Incident   Incident   `json:"incident"`
	Webhook    Webhook    `json:"webhook"`
	LogEntries []LogEntry `json:"log_entries"`
}

// incident-details https://developer.pagerduty.com/docs/webhooks/v2-overview#incident-details

type Incident struct {
	ID                   string               `json:"id"`
	Alerts               []Alert              `json:"alerts"`
	IncidentNumber       int                  `json:"incident_number"`
	Title                string               `json:"title"`
	CreatedAt            string               `json:"created_at"`
	Status               string               `json:"status"`
	IncidentKey          string               `json:"incident_key"`
	HTMLURL              string               `json:"html_url"`
	PendingActions       []PendingAction      `json:"pending_actions"`
	Service              Service              `json:"service"`
	Assignments          []Assignment         `json:"assignments"`
	Acknowledgements     []Acknowledgement    `json:"acknowledgements"`
	LastStatusChangeAt   string               `json:"last_status_change_at"`
	LastStatusChangeBy   LastStatusChangeBy   `json:"last_status_change_by"`
	FirstTriggerLogEntry FirstTriggerLogEntry `json:"first_trigger_log_entry"`
	EscalationPolicy     EscalationPolicy     `json:"escalation_policy"`
	Privilege            string               `json:"privilege"`
	Teams                []Team               `json:"teams"`
	Priority             []Priority           `json:"priority"`
	Urgency              string               `json:"urgency"`
	ResolveReason        string               `json:"resolve_reason"`
	AlertCounts          AlertCounts          `json:"alert_counts"`
	Metadata             Metadata             `json:"metadata"`
	Type                 string               `json:"type"`
	Summary              string               `json:"summary"`
	Self                 string               `json:"self"`
	Description          string               `json:"description"`
	ImpactedServices     []ImpactedService    `json:"impacted_services"`
	IsMergeable          bool                 `json:"is_mergeable"`
	ExternalReferences   []ExternalReference  `json:"external_references"`
	Importance           string               `json:"importance"`
	BasicAlertGrouping   string               `json:"basic_alert_grouping"`
	IncidentsResponders  []IncidentsResponder `json:"incidents_responders"`
	ResponderRequests    []ResponderRequest   `json:"responder_requests"`
	SubscriberRequests   []SubscriberRequest  `json:"subscriber_requests"`
}

type LogEntry struct {
	ID           string       `json:"id"`
	Type         string       `json:"type"`
	Summary      string       `json:"summary"`
	Self         string       `json:"self"`
	HTMLURL      string       `json:"html_url"`
	CreatedAt    string       `json:"created_at"`
	Agent        Agent        `json:"agent"`
	Channel      Channel      `json:"channel"`
	Note         string       `json:"note"`
	Contexts     []Context    `json:"contexts"`
	Incident     Incident     `json:"incident"`
	Service      Service      `json:"service"`
	Teams        []Team       `json:"teams"`
	EventDetails EventDetails `json:"event_details"`
}

type Agent struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Summary string `json:"summary"`
	Self    string `json:"self"`
	HTMLURL string `json:"html_url"`
}

type Channel struct {
	Type    string `json:"type"`
	Summary string `json:"summary"`
	Subject string `json:"subject"`
	Details string `json:"details"`
}

type IncidentsResponder struct{}

type ResponderRequest struct{}

type SubscriberRequest struct{}

type ExternalReference struct{}

type Webhook struct {
	EndpointURL         string              `json:"endpoint_url"`
	Name                string              `json:"name"`
	Description         string              `json:"description"`
	WebhookObject       WebhookObject       `json:"webhook_object"`
	Config              Config              `json:"config"`
	OutboundIntegration OutboundIntegration `json:"outbound_integration"`
	AccountsAddon       string              `json:"account_addon"`
	ID                  string              `json:"id"`
	Type                string              `json:"type"`
	Summary             string              `json:"summary"`
	Self                string              `json:"self"`
	HTMLURL             string              `json:"html_url"`
}

type Alert struct {
	AlertKey string `json:"alert_key"`
}

// PendingAction describes the list of pending_actions on the incident.
// A pending_action object contains a type of action which can be escalate, unacknowledge, resolve or urgency_change.
// A pending_action object contains at, the time at which the action will take place.
// An urgency_change pending_action will contain to, the urgency that the incident will change to.
type PendingAction struct {
	Type string `json:"type"`
	At   string `json:"at"`
}

type Service struct {
	ID                     string              `json:"id"`
	Name                   string              `json:"name"`
	Description            string              `json:"description"`
	AutoResolveTimeout     int                 `json:"auto_resolve_timeout"`
	AcknowledgementTimeout int                 `json:"acknowledgement_timeout"`
	CreatedAt              string              `json:"created_at"`
	Status                 string              `json:"status"`
	SupportHours           string              `json:"support_hours"`
	Addons                 []Addon             `json:"addons"`
	Privilege              string              `json:"priviledge"`
	AlertCreation          string              `json:"alert_creation"`
	Integrations           []Integration       `json:"integrations"`
	ScheduledActions       []ScheduledAction   `json:"scheduled_actions"`
	LastIncidentTimestamp  string              `json:"last_incident_timestamp"`
	IncidentUrgencyRule    IncidentUrgencyRule `json:"incident_urgency_rule"`
	EscalationPolicy       EscalationPolicy    `json:"escalation_policy"`
	Teams                  []Team              `json:"teams"`
	Type                   string              `json:"type"`
	Summary                string              `json:"summary"`
	Self                   string              `json:"self"`
	HTMLURL                string              `json:"html_url"`
	Metadata               Metadata            `json:"metadata"`
}

type ImpactedService struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Summary string `json:"summary"`
	Self    string `json:"self"`
	HTMLURL string `json:"html_url"`
}

type Assignment struct {
	At       string   `json:"at"`
	Assignee Assignee `json:"assignee"`
}

type Assignee struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Summary string `json:"summary"`
	Self    string `json:"self"`
	HTMLURL string `json:"html_url"`
}

type Acknowledgement struct {
	At           string       `json:"at"`
	Acknowledger Acknowledger `json:"acknowledger"`
}

type Acknowledger struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Summary string `json:"summary"`
	Self    string `json:"self"`
	HTMLURL string `json:"html_url"`
}

type LastStatusChangeBy struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Summary string `json:"summary"`
	Self    string `json:"self"`
	HTMLURL string `json:"html_url"`
}

type FirstTriggerLogEntry struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Summary string `json:"summary"`
	Self    string `json:"self"`
	HTMLURL string `json:"html_url"`
}

type EscalationPolicy struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Summary string `json:"summary"`
	Self    string `json:"self"`
	HTMLURL string `json:"html_url"`
}

type Team struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Summary string `json:"summary"`
	Self    string `json:"self"`
	HTMLURL string `json:"html_url"`
}

type Priority struct {
}

type AlertCounts struct {
	All       int `json:"all"`
	Triggered int `json:"triggered"`
	Resolved  int `json:"resolved"`
}

type Metadata struct{}

type WebhookObject struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Summary string `json:"summary"`
	Self    string `json:"self"`
	HTMLURL string `json:"html_url"`
}

type Config struct{}

type OutboundIntegration struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Summary string `json:"summary"`
	Self    string `json:"self"`
	HTMLURL string `json:"html_url"`
}

type Context struct{}

type EventDetails struct {
	Description string `json:"description"`
}

type Addon struct {
}

type Integration struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Summary string `json:"summary"`
	Self    string `json:"self"`
	HTMLURL string `json:"html_url"`
}

type IncidentUrgencyRule struct {
	Type    string `json:"type"`
	Urgency string `json:"urgency"`
}

type ScheduledAction struct {
}




//getCreatedOn
func (msg *Message) GetCreatedOn() (time.Time, error) {
	t, err := time.Parse(time.RFC3339, msg.CreatedOn)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

//getCreatedAt
func (i *Incident) GetCreatedAt() (time.Time, error) {
	t, err := time.Parse(time.RFC3339, i.CreatedAt)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

