package master_author

import (
	"errors"
	"time"
)

type Service interface {
	GetAuthors() ([]MasterAuthor, error)
	GetAuthorById(ID string) (MasterAuthor, error)
	CreateMasterAuthor(input MasterAuthorInput) (MasterAuthor, error)
	EditMasterAuthor(ID string, input MasterAuthorInput) (MasterAuthor, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAuthors() ([]MasterAuthor, error) {
	authors, err := s.repository.GetMasterAuthor()

	if err != nil {
		return []MasterAuthor{}, err
	}

	return authors, err
}

func (s *service) GetAuthorById(ID string) (MasterAuthor, error) {
	author, err := s.repository.FindMasterAuthorByID(ID)

	if err != nil {
		return author, err
	}

	if author.ID == 0 {
		return author, errors.New("data author tidak ditemukan berdasarkan ID")
	}

	return author, nil
}

func (s *service) CreateMasterAuthor(input MasterAuthorInput) (MasterAuthor, error) {
	var masterAuthor MasterAuthor
	masterAuthor.Name = input.Name
	masterAuthor.CreatedAt = time.Now()

	newAuthor, err := s.repository.StoreMasterAuthor(masterAuthor)
	if err != nil {
		return newAuthor, err
	}

	return newAuthor, nil
}

func (s *service) EditMasterAuthor(ID string, input MasterAuthorInput) (MasterAuthor, error) {
	updateAuthor, err := s.repository.UpdateMasterAuthor(ID, input)

	if err != nil {
		return updateAuthor, err
	}

	return updateAuthor, nil
}
