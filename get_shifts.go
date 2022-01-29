package scoober

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

// GetShifts will make a request and return the shift for
// the period in time, with great detail!
func (s *FactoryScoober) GetShifts(startTime, endTime time.Time) ([]Shift, error) {
	if s.Token == "" {
		return nil, errors.New("client is not logged in")
	}

	start := startTime.Format("2006-01-02")
	end := endTime.Format("2006-01-02")

	req, err := http.NewRequest("GET", s.BaseURL+"/api/users/plannings", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("fromDate", start)
	q.Add("toDate", end)
	req.Header.Add("accessToken", s.Token)
	req.URL.RawQuery = q.Encode()

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := []Shift{}
	sb := string(body)
	err = json.Unmarshal([]byte(sb), &data)

	return data, err
}
