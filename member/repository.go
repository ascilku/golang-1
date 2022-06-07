package member

import (
	"gorm.io/gorm"
)

type Repository interface {
	SaveRepositoryMember(member Member) (Member, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) SaveRepositoryMember(member Member) (Member, error) {
	err := r.db.Create(member).Error

	if err != nil {
		return member, err
	} else {
		return member, nil
	}
}
