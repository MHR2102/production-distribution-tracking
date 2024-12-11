package models

import "time"

type Distribution struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	BatchID      uint      `json:"batch_id"`
	SerialNumber string    `gorm:"unique;not null" json:"serial_number"`
	Location     string    `json:"location"`
	Status       string    `json:"status"` // "dikirim" or "diterima"
	UpdatedAt    time.Time `json:"updated_at"`
}
