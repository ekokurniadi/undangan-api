package formatter

import "undangan_online_api/entity"

type UndanganFormatter struct {
	ID                   int     `json:"id"`
	KodeUndangan         string  `json:"kode_undangan"`
	KodeTemaUndangan     string  `json:"kode_tema_undangan"`
	NameTemaUndangan     string  `json:"name_tema_undangan"`
	NamaPria             string  `json:"nama_pria"`
	NamaPanggilanPria    string  `json:"nama_panggilan_pria"`
	NamaWanita           string  `json:"nama_wanita"`
	NamaPanggilanWanita  string  `json:"nama_panggilan_wanita"`
	TanggalResepsi       string  `json:"tanggal_resepsi"`
	NamaOrangTuaPria     string  `json:"nama_orang_tua_pria"`
	NamaOrangTuaWanita   string  `json:"nama_orang_tua_wanita"`
	AlamatOrangTuaPria   string  `json:"alamat_orang_tua_pria"`
	AlamatOrangTuaWanita string  `json:"alamat_orang_tua_wanita"`
	LokasiResepsi        string  `json:"lokasi_resepsi"`
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	LinkVideoPreWed      string  `json:"link_video_pre_wed"`
	Status               int     `json:"status"`
	FotoPria             string  `json:"foto_pria"`
	FotoWanita           string  `json:"foto_wanita"`
	PathUrl              string  `json:"path_url"`
}

func FormatUndangan(undangan entity.Undangan) UndanganFormatter {
	undanganFormatter := UndanganFormatter{}
	undanganFormatter.ID = undangan.ID
	undanganFormatter.KodeUndangan = undangan.KodeUndangan
	undanganFormatter.KodeTemaUndangan = undangan.KodeTemaUndangan
	undanganFormatter.NameTemaUndangan = undangan.NameTemaUndangan
	undanganFormatter.NamaPria = undangan.NamaPria
	undanganFormatter.NamaWanita = undangan.NamaWanita
	undanganFormatter.TanggalResepsi = undangan.TanggalResepsi
	undanganFormatter.NamaOrangTuaPria = undangan.NamaOrangTuaPria
	undanganFormatter.NamaOrangTuaWanita = undangan.NamaOrangTuaWanita
	undanganFormatter.AlamatOrangTuaPria = undangan.AlamatOrangTuaPria
	undanganFormatter.AlamatOrangTuaWanita = undangan.AlamatOrangTuaWanita
	undanganFormatter.LokasiResepsi = undangan.LokasiResepsi
	undanganFormatter.Latitude = undangan.Latitude
	undanganFormatter.Longitude = undangan.Longitude
	undanganFormatter.LinkVideoPreWed = undangan.LinkVideoPreWed
	undanganFormatter.Status = undangan.Status
	undanganFormatter.FotoPria = undangan.FotoPria
	undanganFormatter.FotoWanita = undangan.FotoWanita
	undanganFormatter.NamaPanggilanPria = undangan.NamaPanggilanPria
	undanganFormatter.NamaPanggilanWanita = undangan.NamaPanggilanWanita
	undanganFormatter.PathUrl = undangan.PathUrl
	return undanganFormatter
}
func FormatUndangans(undangans []entity.Undangan) []UndanganFormatter {
	undangansFormatter := []UndanganFormatter{}
	for _, undangan := range undangans {
		undanganFormatter := FormatUndangan(undangan)
		undangansFormatter = append(undangansFormatter, undanganFormatter)
	}
	return undangansFormatter
}

//Generated by Micagen at 05 Juli 2022
