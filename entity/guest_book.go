package entity

import "time"

type GuestBook struct {
	ID              int
	IdUndangan      string
	NamaTamu        string
	StatusKehadiran string
	Pesan           string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
