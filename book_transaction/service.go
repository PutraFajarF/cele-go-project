package book_transaction

import (
	"errors"
	"time"
)

type Service interface {
	GetBookTransactions() ([]BookTransaction, error)
	GetTransactionById(ID string) (BookTransaction, error)
	CreateBookTransaction(input BookTransactionInput) (BookTransaction, error)
	UpdateBookTransaction(ID string, input BookTransactionInput) (BookTransaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetBookTransactions() ([]BookTransaction, error) {
	bookTransactions, err := s.repository.GetBookTransaction()

	if err != nil {
		return []BookTransaction{}, err
	}

	return bookTransactions, err
}

func (s *service) GetTransactionById(ID string) (BookTransaction, error) {
	bookTransaction, err := s.repository.FindBookTransactionByID(ID)
	if err != nil {
		return bookTransaction, err
	}

	if bookTransaction.ID == 0 {
		return bookTransaction, errors.New("data transaksi tidak ditemukan berdasarkan ID")
	}

	return bookTransaction, nil
}

func (s *service) CreateBookTransaction(input BookTransactionInput) (BookTransaction, error) {
	var bookTransaction BookTransaction
	bookTransaction.TotalBook = input.TotalBook
	bookTransaction.TotalPrice = input.TotalPrice
	bookTransaction.BookID = input.BookID
	bookTransaction.UserID = input.UserID
	bookTransaction.CreatedAt = time.Now()

	newTransaction, err := s.repository.StoreBookTransaction(bookTransaction)

	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}

func (s *service) UpdateBookTransaction(ID string, input BookTransactionInput) (BookTransaction, error) {
	newUser, err := s.repository.UpdateBookTransaction(ID, input)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
