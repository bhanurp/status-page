package component

import "time"

type Component struct {
	ID                 string    `json:"id,omitempty"`
	PageID             string    `json:"page_id,omitempty"`
	GroupID            string    `json:"group_id,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
	Group              bool      `json:"group,omitempty"`
	Name               string    `json:"name,omitempty"`
	Description        string    `json:"description,omitempty"`
	Position           int       `json:"position,omitempty"`
	Status             string    `json:"status,omitempty"`
	Showcase           bool      `json:"showcase,omitempty"`
	OnlyShowIfDegraded bool      `json:"only_show_if_degraded,omitempty"`
	AutomationEmail    string    `json:"automation_email,omitempty"`
	StartDate          string    `json:"start_date,omitempty"`
}
