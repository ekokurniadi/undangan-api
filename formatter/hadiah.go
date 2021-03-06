package formatter

import "undangan_online_api/entity"

type HadiahFormatter struct {
	ID             int    `json:"id"`
	IdUndangan     string `json:"id_undangan"`
	NamaBank       string `json:"nama_bank"`
	NoRekening     string `json:"no_rekening"`
	NamaPenerima   string `json:"nama_penerima"`
	AlamatPenerima string `json:"alamat_penerima"`
}

func FormatHadiah(hadiah entity.Hadiah) HadiahFormatter {
	hadiahFormatter := HadiahFormatter{}
	hadiahFormatter.ID = hadiah.ID
	hadiahFormatter.IdUndangan = hadiah.IdUndangan
	hadiahFormatter.NamaBank = hadiah.NamaBank
	hadiahFormatter.NoRekening = hadiah.NoRekening
	hadiahFormatter.NamaPenerima = hadiah.NamaPenerima
	hadiahFormatter.AlamatPenerima = hadiah.AlamatPenerima
	return hadiahFormatter
}
func FormatHadiahs(hadiahs []entity.Hadiah) []HadiahFormatter {
	hadiahsFormatter := []HadiahFormatter{}
	for _, hadiah := range hadiahs {
		hadiahFormatter := FormatHadiah(hadiah)
		hadiahsFormatter = append(hadiahsFormatter, hadiahFormatter)
	}
	return hadiahsFormatter
}

//Generated by Micagen at 05 Juli 2022
