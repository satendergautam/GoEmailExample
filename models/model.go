package models

type EmailRequest struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type Email struct {
	From    string
	To      string
	Cc      string
	Bcc     string
	Subject string
	Body    string
	IsHTML  bool
}
