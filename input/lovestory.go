package input

type InputIDLoveStory struct {
	ID int `uri:"id" binding:"required"`
}

type LoveStoryInput struct {
	IdUndangan string `json:"id_undangan"`
	Judul      string `json:"judul"`
	Tanggal    string `json:"tanggal"`
	Tempat     string `json:"tempat"`
	Keterangan string `json:"keterangan"`
}

//Generated by Micagen at 05 Juli 2022