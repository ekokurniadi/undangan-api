package input

type InputIDUndangan struct {
	ID int `uri:"id" binding:"required"`
}

type UndanganInput struct {
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

//Generated by Micagen at 05 Juli 2022
