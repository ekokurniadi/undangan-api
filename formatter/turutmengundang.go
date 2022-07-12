package formatter

import "undangan_online_api/entity"

type TurutMengundangFormatter struct {
	ID         int    `json:"id"`
	IdUndangan string `json:"id_undangan"`
	Pihak      int    `json:"pihak"`
	Nama       string `json:"nama"`
	Hubungan   string `json:"hubungan"`
}

func FormatTurutMengundang(turutmengundang entity.TurutMengundang) TurutMengundangFormatter {
	turutmengundangFormatter := TurutMengundangFormatter{}
	turutmengundangFormatter.ID = turutmengundang.ID
	turutmengundangFormatter.IdUndangan = turutmengundang.IdUndangan
	turutmengundangFormatter.Pihak = turutmengundang.Pihak
	turutmengundangFormatter.Nama = turutmengundang.Nama
	turutmengundangFormatter.Hubungan = turutmengundang.Hubungan
	return turutmengundangFormatter
}
func FormatTurutMengundangs(turutmengundangs []entity.TurutMengundang) []TurutMengundangFormatter {
	turutmengundangsFormatter := []TurutMengundangFormatter{}
	for _, turutmengundang := range turutmengundangs {
		turutmengundangFormatter := FormatTurutMengundang(turutmengundang)
		turutmengundangsFormatter = append(turutmengundangsFormatter, turutmengundangFormatter)
	}
	return turutmengundangsFormatter
}

//Generated by Micagen at 12 Juli 2022
