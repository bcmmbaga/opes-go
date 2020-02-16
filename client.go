package opes

// Service domain implementation for opes SMS API.
type Service struct {
	auth *auth
}

// NewService ...
func NewService() *Service {
	return &Service{}
}

// Send ...
func (s *Service) Send(msgs ...Message) *SMSResponse {
	return &SMSResponse{}
}
