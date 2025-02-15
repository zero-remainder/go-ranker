package models

import "time"

type AnalyzeRequest struct {
	URL       string    `json:"url"`
	IP        string    `json:"ip"`
	UserAgent string    `json:"userAgent"`
	Host      string    `json:"host"`
	CreatedAt time.Time `json:"createdAt"`
	Status    bool      `json:"status"`
}
