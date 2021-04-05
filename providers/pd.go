package providers

import (
	"strings"
)

const (
	XPDToken = "X-PagerDuty-Token"
	XPDEvent = "X-PagerDuty-Event"
	PagerDutyName   = "pagerduty"
)

type PagerDutyProvider struct {
	secret string
}

func NewPagerDutyProvider(secret string) (*PagerDutyProvider, error) {
	return &PagerDutyProvider{
		secret: secret,
	}, nil
}

func (p *PagerDutyProvider) GetHeaderKeys() []string {
	if len(strings.TrimSpace(p.secret)) > 0 {
		return []string{
			XPDEvent, XPDToken, ContentTypeHeader,
		}
	}

	return []string{
		XPDEvent, ContentTypeHeader,
	}
}

func (p *PagerDutyProvider) GetProviderName() string {
	return PagerDutyName
}

func (p *PagerDutyProvider) Validate(hook Hook) bool {
	token := hook.Headers[XPDToken]
	// validation fails if secret is configured but didnot receive payload from pagerduty
	if len(token) <= 0 {
		return false
	}

	return strings.TrimSpace(token) == strings.TrimSpace(p.secret)
}

