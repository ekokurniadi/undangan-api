package entity

import "time"

type Undangan struct {
	ID                   int
	KodeUndangan         string
	KodeTemaUndangan     int
	NameTemaUndangan     string
	NamaPria             string
	NamaWanita           string
	TanggalResepsi       string
	NamaOrangTuaPria     string
	NamaOrangTuaWanita   string
	AlamatOrangTuaPria   string
	AlamatOrangTuaWanita string
	LokasiResepsi        string
	Latitude             float64
	Longitude            float64
	LinkVideoPreWed      string
	Status               int
	ExpiredAt            time.Time
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
