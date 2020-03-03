package opes

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var configPath string

func init() {
	path, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	path = filepath.Join(path, ".config/opes")
	configPath = path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0700)
		if err != nil {
			log.Fatal(err)
		}
	}

	viper.SetConfigFile(path)
	viper.SetConfigName("config")

}

// Service domain implementation for opes SMS API.
type Service struct {
	Auth   *auth
	Client *http.Client
}

// NewSMSService ...
func NewSMSService() *Service {
	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	return &Service{
		Auth:   generateToken(c),
		Client: c,
	}
}

// // Send ...
// func (s *Service) Send(msgs ...Message) (smsResponse *SMSResponse) {
// 	if ok := s.auth.isvalid(); !ok {
// 		if err := s.refreshToken(); err != nil {
// 			log.Fatalf(err.Error())
// 		}
// 	}

// 	content, _ := json.Marshal(&msgs)
// 	req, _ := http.NewRequest("POST", "", bytes.NewBuffer(content))
// 	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.auth.token))
// 	req.Header.Set("Content-Type", "application/json")

// 	resp, _ := http.DefaultClient.Do(req)
// 	defer resp.Body.Close()

// 	if err := json.NewDecoder(resp.Body).Decode(smsResponse); err != nil {
// 		return
// 	}

// 	return
// }

// Send message
func (s Service) Send(msgs ...Message) *SMSResponse {
	url := "https://sms.opestechnologies.co.tz/api/messages/send"
	if ok := s.Auth.isvalid(); !ok {
		if err := s.refreshToken(); err != nil {
			log.Fatal(err.Error())
		}
	}

	sendReq := smsRequest{}
	sendReq.Messages = append(sendReq.Messages, msgs...)

	content, _ := json.Marshal(&sendReq)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(content))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.Auth.token))

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	smsResponse := &SMSResponse{}
	if err := json.NewDecoder(resp.Body).Decode(smsResponse); err != nil {
		return nil
	}

	fmt.Println(smsResponse)
	return smsResponse
}
