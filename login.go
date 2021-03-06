package scoober

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
)

// Login to a scoober account
func (s *FactoryScoober) Login(email string, password string) (string, error) {
	postBody, err := json.Marshal(map[string]string{
		"userName": email,
		"password": password,
	})

	reqBody := bytes.NewBuffer(postBody)

	if err != nil {
		return "", err
	}

	resp, err := s.Client.Post(s.BaseURL+"/login", "application/json", reqBody)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	type loginResponse struct {
		Token     string `json:"accessToken"`
		Firstname string `json:"firstName"`
		Lastname  string `json:"lastName"`
	}

	sb := string(body)
	data := loginResponse{}
	err = json.Unmarshal([]byte(sb), &data)
	if err != nil {
		return "", err
	}

	if data.Token == "" {
		return "", errors.New("failed to sign in")
	}

	s.Token = data.Token

	return "", nil
}
