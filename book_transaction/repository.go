package book_transaction

import (
	"errors"
	"project-go/master_book"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	GetBookTransaction() ([]BookTransaction, error)
	FindBookTransactionByID(ID string) (BookTransaction, error)
	StoreBookTransaction(bookTransaction BookTransaction) (BookTransaction, error)
	UpdateBookTransaction(ID string, input BookTransactionInput) (BookTransaction, error)
	// DeleteMasterBook(masterBook MasterBook, ID int) (MasterBook, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetBookTransaction() ([]BookTransaction, error) {
	var bt []BookTransaction

	if err := r.db.Find(&bt).Error; err != nil {
		return nil, err
	}

	return bt, nil
}

func (r *repository) FindBookTransactionByID(ID string) (BookTransaction, error) {
	var bookTransaction BookTransaction

	err := r.db.Where("id = ?", ID).Find(&bookTransaction).Error
	if err != nil {
		return bookTransaction, err
	}

	return bookTransaction, nil
}

func (r *repository) StoreBookTransaction(bookTransaction BookTransaction) (BookTransaction, error) {
	var masterBook master_book.MasterBook
	err := r.db.Table("master_books").Where("id", bookTransaction.BookID).First(masterBook).Error

	if err != nil {
		return bookTransaction, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) != true {
		return bookTransaction, errors.New("Data buku tidak ditemukan")
	}

	if masterBook.Amount >= bookTransaction.TotalBook {
		masterBook.Amount -= bookTransaction.TotalBook
		bookTransaction.TotalPrice = masterBook.Price * bookTransaction.TotalBook
	} else {
		return bookTransaction, errors.New("Pesanan melebihi jumlah stok buku")
	}

	err = r.db.Create(&bookTransaction).Error

	if err != nil {
		return bookTransaction, err
	}

	err = r.db.Updates(&masterBook).Error

	if err != nil {
		return bookTransaction, err
	}

	return bookTransaction, nil
}

func (r *repository) UpdateBookTransaction(ID string, input BookTransactionInput) (BookTransaction, error) {
	var masterBook master_book.MasterBook
	var bookTransaction BookTransaction
	err := r.db.Where("id = ?", ID).First(bookTransaction).Error

	if err != nil {
		return bookTransaction, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) == true {
		return bookTransaction, errors.New("Data transaksi tidak ditemukan")
	}

	err = r.db.Table("master_books").Where("id", bookTransaction.BookID).First(masterBook).Error

	if err != nil {
		return bookTransaction, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) != true {
		return bookTransaction, errors.New("Data buku tidak ditemukan")
	}

	if masterBook.Amount >= input.TotalBook {
		masterBook.Amount -= input.TotalBook
		bookTransaction.TotalPrice = masterBook.Amount * bookTransaction.TotalBook
	} else {
		return BookTransaction{}, errors.New("Pesanan melebihi jumlah stok buku")
	}

	bookTransaction.BookID = masterBook.ID
	bookTransaction.UserID = input.UserID
	bookTransaction.TotalBook = input.TotalBook
	bookTransaction.UpdatedAt = time.Now()

	err = r.db.Updates(&bookTransaction).Error

	if err != nil {
		return bookTransaction, err
	}

	return bookTransaction, nil
}

// func (r *repository) DeleteMasterBook(masterBook MasterBook, ID int) error {
// 	err := r.db.Delete(&masterBook, ID).Error

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
