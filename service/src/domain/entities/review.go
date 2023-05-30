package entities

import "time"

type Review struct {
	Id          string    `json:"id"`
	OrderId     string    `json:"order_id"`
	Description string    `json:"description"`
	Rating      float32   `json:"rating"`
	Date        time.Time `json:"date"`
}
