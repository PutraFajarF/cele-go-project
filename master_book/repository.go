package master_book

import (
	"errors"
	"project-go/entities"

	"gorm.io/gorm"
)

type Repository interface {
	GetMasterBook() ([]entities.MasterBook, error)
	FindMasterBookByID(ID string) (entities.MasterBook, error)
	StoreMasterBook(masterBook entities.MasterBook) (entities.MasterBook, error)
	UpdateMasterBook(ID string, input MasterBookInput) (entities.MasterBook, error)
	// DeleteMasterBook(masterBook MasterBook, ID int) (MasterBook, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetMasterBook() ([]entities.MasterBook, error) {
	var mb []entities.MasterBook

	if err := r.db.Preload("BookTransaction").Find(&mb).Error; err != nil {
		return nil, err
	}

	return mb, nil
}

func (r *repository) FindMasterBookByID(ID string) (entities.MasterBook, error) {
	var masterBook entities.MasterBook

	err := r.db.Preload("BookTransaction").Where("id = ?", ID).Find(&masterBook).Error
	if err != nil {
		return masterBook, err
	}

	return masterBook, nil
}

func (r *repository) StoreMasterBook(masterBook entities.MasterBook) (entities.MasterBook, error) {
	var masterAuthor entities.MasterAuthor
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

func (r *repository) UpdateMasterBook(ID string, input MasterBookInput) (entities.MasterBook, error) {
	var masterBook entities.MasterBook
	var masterAuthor entities.MasterAuthor

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
