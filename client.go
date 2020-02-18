package opes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var smsEndpoint = "https://sms.opestechnologies.co.tz/api/messages/send"

// Service domain implementation for opes SMS API.
type Service struct {
	username string
	password string
	auth     *auth
}

// NewService ...
func NewService(username, password string) *Service {
	return &Service{
		username: username,
		password: password,
		auth:     generateToken(username, password),
	}
}

// Send ...
func (s *Service) Send(msgs ...Message) (smsResponse *SMSResponse) {
	if ok := s.auth.isvalid(); !ok {
		if err := s.refreshToken(); err != nil {
			log.Fatalf(err.Error())
		}
	}

	content, _ := json.Marshal(&msgs)
	req, _ := http.NewRequest("POST", smsEndpoint, bytes.NewBuffer(content))
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.auth.token))
	req.Header.Set("Content-Type", "application/json")

	
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(smsResponse); err != {
		return 
	}

	return 
}
