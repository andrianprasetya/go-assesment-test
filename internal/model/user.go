package model

import "time"

type User struct {
	ID         string    `gorm:"type:varchar(50);primary_key:true" json:"id"`
	Name       string    `gorm:"type:varchar(100);not null" json:"name"`
	Nik        string    `gorm:"type:varchar(25);not null" json:"nik"`
	NoHp       string    `gorm:"type:varchar(15);not null" json:"no_hp"`
	NoRekening string    `gorm:"type:varchar(10);not null" json:"no_rekening"`
	Balance    int       `gorm:"type:integer;default:0" json:"balance"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
