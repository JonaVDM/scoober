package scoober

import (
	"net/http"
)

// Scoober is the main scoober account object
type Scoober struct {
	Token   string
	Client  *http.Client
	BaseURL string
	Name    string
}

// NewScoober creates a new scoober object
func NewScoober(token, name string) *Scoober {
	return &Scoober{
		Client:  &http.Client{},
		BaseURL: "https://shiftplanning-api.scoober.com",
		Name:    name,
		Token:   token,
	}
}
