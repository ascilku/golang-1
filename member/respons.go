package member

type respons struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Pesan  string
	Kode   int
	Status string
}

func NewRespons(pesan string, kode int, status string, data interface{}) respons {
	keyMeta := meta{
		Pesan:  pesan,
		Kode:   kode,
		Status: status,
	}

	jsonRespons := respons{
		Meta: keyMeta,
		Data: data,
	}

	return jsonRespons
}
