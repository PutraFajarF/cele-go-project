package master_author

import (
	"errors"
	"project-go/entities"
	"time"
)

type Service interface {
	GetAuthors() ([]entities.MasterAuthor, error)
	GetAuthorById(ID string) (entities.MasterAuthor, error)
	CreateMasterAuthor(input MasterAuthorInput) (entities.MasterAuthor, error)
	EditMasterAuthor(ID string, input MasterAuthorInput) (entities.MasterAuthor, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAuthors() ([]entities.MasterAuthor, error) {
	authors, err := s.repository.GetMasterAuthor()

	if err != nil {
		return []entities.MasterAuthor{}, err
	}

	return authors, err
}

func (s *service) GetAuthorById(ID string) (entities.MasterAuthor, error) {
	author, err := s.repository.FindMasterAuthorByID(ID)

	if err != nil {
		return author, err
	}

	if author.ID == 0 {
		return author, errors.New("data author tidak ditemukan berdasarkan ID")
	}

	return author, nil
}

func (s *service) CreateMasterAuthor(input MasterAuthorInput) (entities.MasterAuthor, error) {
	var masterAuthor entities.MasterAuthor
	masterAuthor.Name = input.Name
	masterAuthor.CreatedAt = time.Now()

	newAuthor, err := s.repository.StoreMasterAuthor(masterAuthor)
	if err != nil {
		return newAuthor, err
	}

	return newAuthor, nil
}

func (s *service) EditMasterAuthor(ID string, input MasterAuthorInput) (entities.MasterAuthor, error) {
	updateAuthor, err := s.repository.UpdateMasterAuthor(ID, input)

	if err != nil {
		return updateAuthor, err
	}

	return updateAuthor, nil
}
