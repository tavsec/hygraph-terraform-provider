package models

type webhook struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type webhookResponse struct {
	Webhooks []webhook `json:"webhooks"`
}
