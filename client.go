package scoober

import (
	"net/http"
)

// Scoober is the interface defining all the methods that the struct
// implementing this needs to have
type Scoober interface {
	// GetShifts gets all the shifts between the dates start and end.
	GetShifts(start, end string) ([]Shift, error)

	// Login into your scoober account with email and password. Returns the
	// access token and maybe an error. Fun fact, did you know that the token
	// being returned never changes.
	Login(email, password string) (string, error)
}

// FactoryScoober is the main scoober account object
type FactoryScoober struct {
	Token   string
	Client  *http.Client
	BaseURL string
}

// NewScoober creates a new scoober object
func NewScoober(token string) Scoober {
	return &FactoryScoober{
		Client:  &http.Client{},
		BaseURL: "https://shiftplanning-api.scoober.com",
		Token:   token,
	}
}
