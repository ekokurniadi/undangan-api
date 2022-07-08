package datatables

type Request struct {
	Draw    string    `json:"draw"`
	Columns []Columns `json:"columns"`
	Order   Order     `json:"order"`
	Start   string    `json:"start"`
	Length  string    `json:"length"`
	Search  Search    `json:"search"`
}

type Columns struct {
	Data       string `json:"data"`
	Name       string `json:"name"`
	SearchAble string `json:"searchable"`
	OrderAble  string `json:"orderable"`
	Search     Search `json:"search"`
}

type Search struct {
	Value string `json:"value"`
	Regex string `json:"regex"`
}

type Order struct {
	Column string `json:"column"`
	Dir    string `json:"dir"`
}
