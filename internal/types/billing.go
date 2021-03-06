package types

import "time"

// VehicleDetails -
type VehicleDetails struct {
	ID         uint      `json:"id" gorm:"primaryKey,autoIncrement"`
	ArrivedAt  time.Time `json:"arrived_at" gorm:"autoCreateTime"`
	DepartedAt time.Time `json:"departed_at"`
	RegNo      string    `json:"reg_no"`
	Type       string    `json:"type"`
}
