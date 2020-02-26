package opes

import (
	"crypto/tls"
	"net/http"
)

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
