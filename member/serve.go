package member

import "golang.org/x/crypto/bcrypt"

type Serve interface {
	SaveServeMember(inputMember InputMember) (Member, error)
}

type serve struct {
	repository Repository
}

func NewServeMember(repository Repository) *serve {
	return &serve{repository}
}

func (s *serve) SaveServeMember(inputMember InputMember) (Member, error) {
	serveMember := Member{}
	serveMember.Nri = inputMember.Nri
	serveMember.Nama = inputMember.Nama
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(inputMember.Angkatan), bcrypt.MinCost)

	if err != nil {
		return serveMember, err
	} else {
		serveMember.Angkatan = string(passwordHash)

		newServeMember, err := s.repository.SaveRepositoryMember(serveMember)
		if err != nil {
			return newServeMember, err
		} else {
			return newServeMember, nil
		}
	}
}
