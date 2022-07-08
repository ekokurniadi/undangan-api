package entity

import "time"

type Undangan struct {
	ID                   int
	KodeUndangan         string
	KodeTemaUndangan     string
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
	FotoPria             string
	FotoWanita           string
	ExpiredAt            time.Time
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
