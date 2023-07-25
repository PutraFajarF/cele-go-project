package master_book

import (
	"errors"
)

type Service interface {
	GetBooks() ([]MasterBook, error)
	GetBookById(ID string) (MasterBook, error)
	CreateMasterBook(input MasterBookInput) (MasterBook, error)
	EditMasterBook(ID string, input MasterBookInput) (MasterBook, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetBooks() ([]MasterBook, error) {
	books, err := s.repository.GetMasterBook()

	if err != nil {
		return []MasterBook{}, err
	}

	return books, err
}

func (s *service) GetBookById(ID string) (MasterBook, error) {
	user, err := s.repository.FindMasterBookByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("data user tidak ditemukan berdasarkan ID")
	}

	return user, nil
}

func (s *service) CreateMasterBook(input MasterBookInput) (MasterBook, error) {
	var masterBook MasterBook
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

func (s *service) EditMasterBook(ID string, input MasterBookInput) (MasterBook, error) {
	updateBook, err := s.repository.UpdateMasterBook(ID, input)

	if err != nil {
		return updateBook, err
	}

	return updateBook, nil
}
