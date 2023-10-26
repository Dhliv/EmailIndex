package models

type EmailsResponse struct {
	Hits Emails `json:"hits"`
}

type Emails struct {
	Hits []EmailResponse `json:"hits"`
}