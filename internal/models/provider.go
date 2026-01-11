package models

type Provider struct {
	Link        string `json:"link"`
	Description string `json:"description,omitempty"`
}
