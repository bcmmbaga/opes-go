package opes

// Service domain implementation for opes SMS API.
type Service struct {
	Auth *auth
}

// NewSMSService ...
func NewSMSService() *Service {
	return &Service{
		Auth: generateToken(),
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
