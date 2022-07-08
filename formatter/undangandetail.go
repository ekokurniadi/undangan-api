package formatter

import "undangan_online_api/entity"

type UndanganDetailFormatter struct {
	ID           int    `json:"id"`
	IdUndangan   string `json:"id_undangan"`
	NamaAcara    string `json:"nama_acara"`
	TanggalAcara string `json:"tanggal_acara"`
	WaktuAcara   string `json:"waktu_acara"`
	LokasiAcara  string `json:"lokasi_acara"`
	AlamatAcara  string `json:"alamat_acara"`
}

func FormatUndanganDetail(undangandetail entity.UndanganDetail) UndanganDetailFormatter {
	undangandetailFormatter := UndanganDetailFormatter{}
	undangandetailFormatter.ID = undangandetail.ID
	undangandetailFormatter.IdUndangan = undangandetail.IdUndangan
	undangandetailFormatter.NamaAcara = undangandetail.NamaAcara
	undangandetailFormatter.TanggalAcara = undangandetail.TanggalAcara
	undangandetailFormatter.WaktuAcara = undangandetail.WaktuAcara
	undangandetailFormatter.LokasiAcara = undangandetail.LokasiAcara
	undangandetailFormatter.AlamatAcara = undangandetail.AlamatAcara
	return undangandetailFormatter
}
func FormatUndanganDetails(undangandetails []entity.UndanganDetail) []UndanganDetailFormatter {
	undangandetailsFormatter := []UndanganDetailFormatter{}
	for _, undangandetail := range undangandetails {
		undangandetailFormatter := FormatUndanganDetail(undangandetail)
		undangandetailsFormatter = append(undangandetailsFormatter, undangandetailFormatter)
	}
	return undangandetailsFormatter
}

//Generated by Micagen at 05 Juli 2022