package shippinglabel

import "time"

type User struct {
	ID            int        `json:"id,omitempty"`
	Email         string     `json:"email,omitempty"`
	Lang          string     `json:"language,omitempty"`
	RegisterDate  *time.Time `json:"registerDate,omitempty"`
	LastSeen      *time.Time `json:"lastSeen,omitempty"`
	Balance       float64    `json:"balance,omitempty"`
	Address       *Address   `json:"address,omitempty"`
	ShipmentStats *Stats     `json:"shipmentStats,omitempty"`
}

type Stats struct {
	Total        float32 `json:"total,omitempty"`
	CurrentYear  float32 `json:"currentYear,omitempty"`
	CurrentMonth float32 `json:"currentMonth,omitempty"`
	Today        float32 `json:"today,omitempty"`
	LastYear     float32 `json:"lastYear,omitempty"`
}
