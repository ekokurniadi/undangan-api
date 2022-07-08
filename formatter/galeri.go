package formatter

import "undangan_online_api/entity"

type GaleriFormatter struct {
	ID         int    `json:"id"`
	IdUndangan string `json:"id_undangan"`
	Foto       string `json:"foto"`
}

func FormatGaleri(galeri entity.Galeri) GaleriFormatter {
	galeriFormatter := GaleriFormatter{}
	galeriFormatter.ID = galeri.ID
	galeriFormatter.IdUndangan = galeri.IdUndangan
	galeriFormatter.Foto = galeri.Foto
	return galeriFormatter
}
func FormatGaleris(galeris []entity.Galeri) []GaleriFormatter {
	galerisFormatter := []GaleriFormatter{}
	for _, galeri := range galeris {
		galeriFormatter := FormatGaleri(galeri)
		galerisFormatter = append(galerisFormatter, galeriFormatter)
	}
	return galerisFormatter
}

//Generated by Micagen at 05 Juli 2022