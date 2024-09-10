package models

import "time"

type Notification struct {
	ID                    int       `json:"id" db:"id"`
	CurrentPrice          float64   `json:"current_price" db:"current_price"`
	DailyChangePercentage float64   `json:"daily_change_percentage" db:"daily_change_percentage"`
	TradingVolume         int       `json:"trading_volume" db:"trading_volume"`
	Emails                []string  `json:"emails" db:"emails"`
	Status                string    `json:"status" db:"status"`
	CreatedAt             time.Time `json:"created_at" db:"created_at"`
	UpdatedAt             time.Time `json:"updated_at" db:"updated_at"`
}
