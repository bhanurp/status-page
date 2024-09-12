package incident

import (
	"time"
)

type CreateIncident struct {
	HostName       string
	APIKey         string
	IncidentName   string
	IncidentStatus string
	IncidentBody   string
	ComponentID    string
	PageID         string
	IncidentHeader string
	Metadata       Metadata
}

type Metadata struct {
	Data Data `json:"data"`
}

type Data struct {
}

type UpdateIncident struct {
	HostName       string
	APIKey         string
	IncidentName   string
	IncidentStatus string
	IncidentBody   string
	ComponentID    string
	PageID         string
	IncidentHeader string
	Metadata       Metadata
}

type Incident struct {
	ID         string `json:"id"`
	Components []struct {
		ID                 string    `json:"id"`
		PageID             string    `json:"page_id"`
		GroupID            string    `json:"group_id"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
		Group              bool      `json:"group"`
		Name               string    `json:"name"`
		Description        any       `json:"description"`
		Position           int       `json:"position"`
		Status             string    `json:"status"`
		Showcase           bool      `json:"showcase"`
		OnlyShowIfDegraded bool      `json:"only_show_if_degraded"`
		AutomationEmail    string    `json:"automation_email"`
		StartDate          string    `json:"start_date"`
	} `json:"components"`
	CreatedAt                                 time.Time `json:"created_at"`
	Impact                                    string    `json:"impact"`
	Name                                      string    `json:"name"`
	PageID                                    string    `json:"page_id"`
	PostmortemBody                            any       `json:"postmortem_body"`
	PostmortemBodyLastUpdatedAt               any       `json:"postmortem_body_last_updated_at"`
	PostmortemIgnored                         bool      `json:"postmortem_ignored"`
	PostmortemNotifiedSubscribers             bool      `json:"postmortem_notified_subscribers"`
	PostmortemNotifiedTwitter                 bool      `json:"postmortem_notified_twitter"`
	PostmortemPublishedAt                     any       `json:"postmortem_published_at"`
	ResolvedAt                                any       `json:"resolved_at"`
	ScheduledAutoCompleted                    bool      `json:"scheduled_auto_completed"`
	ScheduledAutoInProgress                   bool      `json:"scheduled_auto_in_progress"`
	ScheduledFor                              any       `json:"scheduled_for"`
	AutoTransitionDeliverNotificationsAtEnd   any       `json:"auto_transition_deliver_notifications_at_end"`
	AutoTransitionDeliverNotificationsAtStart any       `json:"auto_transition_deliver_notifications_at_start"`
	AutoTransitionToMaintenanceState          any       `json:"auto_transition_to_maintenance_state"`
	AutoTransitionToOperationalState          any       `json:"auto_transition_to_operational_state"`
	ScheduledRemindPrior                      bool      `json:"scheduled_remind_prior"`
	ScheduledRemindedAt                       any       `json:"scheduled_reminded_at"`
	ScheduledUntil                            any       `json:"scheduled_until"`
	Shortlink                                 string    `json:"shortlink"`
	Status                                    string    `json:"status"`
	UpdatedAt                                 time.Time `json:"updated_at"`
	Metadata                                  Metadata  `json:"metadata"`
}

type Payload struct {
	Incident IncidentData `json:"incident"`
}

type IncidentData struct {
	Name         string            `json:"name"`
	Body         string            `json:"body"`
	Status       string            `json:"status"`
	ComponentIds []string          `json:"component_ids"`
	Components   map[string]string `json:"components"`
	Metadata     Metadata          `json:"metadata"`
}
