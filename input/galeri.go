package input

type InputIDGaleri struct {
	ID int `uri:"id" binding:"required"`
}

type GaleriInput struct {
	IdUndangan string `json:"id_undangan"`
	Foto       string `json:"foto"`
}

//Generated by Micagen at 05 Juli 2022