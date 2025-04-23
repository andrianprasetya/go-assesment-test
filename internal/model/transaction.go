package model

import "time"

type Transaction struct {
	ID              string    `gorm:"type:varchar(50);primary_key:true" json:"id"`
	UserId          string    `gorm:"type:varchar(50);not null" json:"user_id"`
	TypeTransaction string    `gorm:"type:varchar(5);not null;comment: C => Credit | D => Debit" json:"type_transaction"`
	Amount          int       `gorm:"type:integer;not null" json:"no_rekening"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
