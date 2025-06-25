package entity

import "time"

type Customer struct {
	Customer_id int       `json:"customer_id"`
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

type UploadedFile struct {
	Filename   string `json:"filename"`
	UploadedAt time.Time `json:"uploaded_at"` // atau time.Time jika kamu pakai time parsing
}
