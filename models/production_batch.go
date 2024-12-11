package models

import "time"

type ProductionBatch struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	BatchNumber string    `gorm:"unique;not null" json:"batch_number"`
	TotalUnits  int       `json:"total_units"`
	Status      string    `json:"status"` // "proses", "selesai", or "gagal"
	CreatedAt   time.Time `json:"created_at"`
}
