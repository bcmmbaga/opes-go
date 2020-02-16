package opes

// Message response code
const (
	StatusSuccess              = 200
	StatusAuthenticationFailed = 416
	StatusGeneralError         = 406
)

var showText = map[int]string{
	StatusSuccess:              "success",
	StatusAuthenticationFailed: "Authentication Failed",
	StatusGeneralError:         "General Error",
}
