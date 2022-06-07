package member

type InputMember struct {
	Nri      string `json:"nri" binding:"required"`
	Nama     string `json:"nama" binding:"required"`
	Angkatan string `json:"angkatan" binding:"required,email"`
}
