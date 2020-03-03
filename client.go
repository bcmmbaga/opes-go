package opes

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	configPath = fmt.Sprintf("%s/%s", path, "config.toml")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0700)
		if err != nil {
			log.Fatal(err)
		}
	}

	viper.AddConfigPath(path)
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

	svc := &Service{}
	svc.Client = c

	// generate new token if config file does not exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		svc.Auth = generateToken(c)
	}

	svc.Auth = generateTokenFromFile()

	return svc
}

// Send message
func (s Service) Send(msgs ...Message) *SMSResponses {
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

	jsonContent, _ := ioutil.ReadAll(resp.Body)
	smsResponses := new(SMSResponses)
	json.Unmarshal(jsonContent, smsResponses)

	return smsResponses
}
