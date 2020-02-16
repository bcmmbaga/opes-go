package opes

import "time"

type auth struct {
	token  string
	expire time.Time
}

func (s *Service) refreshToken() error {
	if ok := s.auth.isvalid(); ok {
		return nil
	}

	//refresh new token by getting from endpoint
	return nil
}

func (a *auth) isvalid() bool {
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
