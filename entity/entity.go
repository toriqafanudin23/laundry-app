package entity

import "time"

type Customer struct {
	Customer_id int
	Name        string
	Phone       string
	Address     string
	Created_at  time.Time
	Updated_at  time.Time
}
