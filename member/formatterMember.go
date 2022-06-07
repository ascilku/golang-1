package member

import "github.com/go-playground/validator/v10"

type formatterMember struct {
	Nri      string `json:"nri"`
	Nama     string `json:"nama"`
	Angkatan string `json:"angkatan"`
}

func NewFormatterMember(member Member, token string) formatterMember {
	keyFormatterMember := formatterMember{}
	keyFormatterMember.Nri = member.Nri
	keyFormatterMember.Nama = member.Nama
	keyFormatterMember.Angkatan = member.Angkatan

	return keyFormatterMember
}

func NewFormatterError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
