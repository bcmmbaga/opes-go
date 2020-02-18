package opes

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type auth struct {
	token  string
	expire time.Time
}

var tokenEndpoint = "https://sms.opestechnologies.co.tz/api/get-api-key"

func generateToken(username, password string) (auth *auth) {
	req := &request{
		username: username,
		Password: password,
	}

	content, _ := json.Marshal(req)
	resp, err := http.DefaultClient.Post(tokenEndpoint, "application/json", bytes.NewBuffer(content))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	tokenResp := &response{}
	if err := json.NewDecoder(resp.Body).Decode(tokenResp); err != nil {
		return
	}

	if tokenResp.Error != "" {
		return
	}

	auth.token = tokenResp.Success.Token
	// expires after 15 days starting from day token was generated.
	auth.expire = time.Now().Add(time.Hour * 24 * 15)

	return
}

func (s *Service) refreshToken() error {
	auth := generateToken(s.username, s.password)
	if auth == nil {
		return errors.New("failed to refresh token")
	}

	s.auth = auth
	return nil
}

func (a *auth) isvalid() bool {

	// token is valid if number of days remaining is lessor equal to one.
	if remainingDays := int(a.expire.Sub(time.Now()).Hours()) / 24; remainingDays <= 1 {
		return false
	}

	return true
}

type request struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type response struct {
	Success struct {
		Token  string `json:"token"`
		Client string `json:"client"`
	} `json:"success,omitempty"`
	Error string `json:"error,omitempty"`
}
