package opes

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.Set("USERNAME", os.Getenv("OPES_USERNAME"))
	viper.Set("PASSWORD", os.Getenv("OPES_PASSWORD"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("%s", err)
	}

}

type config struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type auth struct {
	token  string
	expire time.Time
}

// generateToken request authorization token by sending a sync request
// with username and password from https://sms.opestechnologies.co.tz/api/get-api-key
func generateToken() (auth *auth) {
	req := &config{
		Username: viper.GetString("USERNAME"),
		Password: viper.GetString("PASSWORD"),
	}

	content, _ := json.Marshal(req)
	resp, err := http.DefaultClient.Post(viper.GetString("Endpoints.Token"), "application/json", bytes.NewBuffer(content))
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
	auth := generateToken()
	if auth == nil {
		return errors.New("failed to refresh token")
	}

	s.Auth = auth
	return nil
}

func (a *auth) isvalid() bool {

	// token is valid if number of days remaining is lessor equal to one.
	if remainingDays := int(a.expire.Sub(time.Now()).Hours()) / 24; remainingDays <= 1 {
		return false
	}

	return true
}

type response struct {
	Success struct {
		Token  string `json:"token"`
		Client string `json:"client"`
	} `json:"success,omitempty"`
	Error string `json:"error,omitempty"`
}
