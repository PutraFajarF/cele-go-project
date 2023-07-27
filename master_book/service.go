package master_book

import (
	"errors"
	"project-go/entities"
)

type Service interface {
	GetBooks() ([]entities.MasterBook, error)
	GetBookById(ID string) (entities.MasterBook, error)
	CreateMasterBook(input MasterBookInput) (entities.MasterBook, error)
	EditMasterBook(ID string, input MasterBookInput) (entities.MasterBook, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetBooks() ([]entities.MasterBook, error) {
	books, err := s.repository.GetMasterBook()

	if err != nil {
		return []entities.MasterBook{}, err
	}

	return books, err
}

func (s *service) GetBookById(ID string) (entities.MasterBook, error) {
	user, err := s.repository.FindMasterBookByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("data user tidak ditemukan berdasarkan ID")
	}

	return user, nil
}

func (s *service) CreateMasterBook(input MasterBookInput) (entities.MasterBook, error) {
	var masterBook entities.MasterBook
	masterBook.Name = input.Name
	masterBook.Amount = input.Amount
	masterBook.Price = input.Price
	masterBook.AuthorID = input.AuthorID

	newBook, err := s.repository.StoreMasterBook(masterBook)

	if err != nil {
		return newBook, err
	}

	return newBook, nil
}

func (s *service) EditMasterBook(ID string, input MasterBookInput) (entities.MasterBook, error) {
	updateBook, err := s.repository.UpdateMasterBook(ID, input)

	if err != nil {
		return updateBook, err
	}

	return updateBook, nil
}
