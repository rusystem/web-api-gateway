package domain

import "time"

type Company struct {
	ID         int       `json:"id" db:"id"`
	NameRU     string    `json:"name_ru" db:"name_ru"`
	NameEN     string    `json:"name_en" db:"name_en"`
	Country    string    `json:"country" db:"country"`
	Address    string    `json:"address" db:"address"`
	Phone      string    `json:"phone" db:"phone"`
	Email      string    `json:"email" db:"email"`
	Website    string    `json:"website" db:"website"`
	IsActive   bool      `json:"is_active" db:"is_active" default:"true"`
	CreatedAt  time.Time `json:"created_at" db:"created_at" default:"CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at" default:"CURRENT_TIMESTAMP"`
	IsApproved bool      `json:"is_approved" db:"is_approved"`
	Timezone   string    `json:"timezone" db:"timezone"`
}
