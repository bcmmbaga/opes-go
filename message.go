package opes

// Message ...
type Message struct {
	ID      int    `json:"msg_id"`
	Sender  string `json:"sender"`
	Channel int    `json:"channel"`
	Text    string `json:"text"`
	MSISDN  string `json:"msisdn"`
}

type smsRequest struct {
	Messages []Message `json:"messages"`
}

// SMSResponses response
type SMSResponses []struct {
	ResultCode int    `json:"result_code"`
	Result     string `json:"result"`
	Reference  int    `json:"reference"`
	Message    string `json:"message"`
}
