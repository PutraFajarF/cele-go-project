package master_book

import (
	"errors"
	"project-go/master_author"

	"gorm.io/gorm"
)

type Repository interface {
	GetMasterBook() ([]MasterBook, error)
	FindMasterBookByID(ID string) (MasterBook, error)
	StoreMasterBook(masterBook MasterBook) (MasterBook, error)
	UpdateMasterBook(ID string, input MasterBookInput) (MasterBook, error)
	// DeleteMasterBook(masterBook MasterBook, ID int) (MasterBook, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetMasterBook() ([]MasterBook, error) {
	var mb []MasterBook

	if err := r.db.Find(&mb).Error; err != nil {
		return nil, err
	}

	return mb, nil
}

func (r *repository) FindMasterBookByID(ID string) (MasterBook, error) {
	var masterBook MasterBook

	err := r.db.Where("id = ?", ID).Find(&masterBook).Error
	if err != nil {
		return masterBook, err
	}

	return masterBook, nil
}

func (r *repository) StoreMasterBook(masterBook MasterBook) (MasterBook, error) {
	var masterAuthor master_author.MasterAuthor
	err := r.db.Table("master_authors").Where("id = ?", masterBook.AuthorID).First(&masterAuthor).Error

	if err != nil {
		return masterBook, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) == true {
		return masterBook, gorm.ErrRecordNotFound
	}

	err = r.db.Create(&masterBook).Error

	if err != nil {
		return masterBook, err
	}

	return masterBook, nil
}

func (r *repository) UpdateMasterBook(ID string, input MasterBookInput) (MasterBook, error) {
	var masterBook MasterBook
	var masterAuthor master_author.MasterAuthor

	err := r.db.Where("id = ?", ID).First(&masterBook).Error

	if err != nil {
		return masterBook, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) == true {
		return masterBook, errors.New("Data buku tidak ditemukan")
	}

	err = r.db.Table("master_authors").Where("id = ?", masterBook.ID).First(&masterAuthor).Error

	if err != nil {
		return masterBook, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) == true {
		return masterBook, errors.New("Data author tidak ditemukan")
	}

	masterBook.Name = input.Name
	masterBook.Amount = input.Amount
	masterBook.Price = input.Price
	masterBook.AuthorID = input.AuthorID

	err = r.db.Updates(&masterBook).Error

	if err != nil {
		return masterBook, err
	}

	return masterBook, nil
}

// func (r *repository) DeleteMasterBook(masterBook MasterBook, ID int) error {
// 	err := r.db.Delete(&masterBook, ID).Error

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
