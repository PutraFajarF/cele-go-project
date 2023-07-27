package master_author

import (
	"errors"
	"project-go/entities"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	GetMasterAuthor() ([]entities.MasterAuthor, error)
	FindMasterAuthorByID(ID string) (entities.MasterAuthor, error)
	StoreMasterAuthor(masterAuthor entities.MasterAuthor) (entities.MasterAuthor, error)
	UpdateMasterAuthor(ID string, input MasterAuthorInput) (entities.MasterAuthor, error)
	// DeleteMasterBook(masterBook MasterBook, ID int) (MasterBook, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetMasterAuthor() ([]entities.MasterAuthor, error) {
	var ma []entities.MasterAuthor

	if err := r.db.Preload("MasterBook").Find(&ma).Error; err != nil {
		return nil, err
	}

	return ma, nil
}

func (r *repository) FindMasterAuthorByID(ID string) (entities.MasterAuthor, error) {
	var masterAuthor entities.MasterAuthor

	err := r.db.Preload("MasterBook").Where("id = ?", ID).Find(&masterAuthor).Error
	if err != nil {
		return masterAuthor, err
	}

	return masterAuthor, nil
}

func (r *repository) StoreMasterAuthor(masterAuthor entities.MasterAuthor) (entities.MasterAuthor, error) {
	err := r.db.Create(&masterAuthor).Error

	if err != nil {
		return masterAuthor, err
	}

	return masterAuthor, nil
}

func (r *repository) UpdateMasterAuthor(ID string, input MasterAuthorInput) (entities.MasterAuthor, error) {
	var masterAuthor entities.MasterAuthor

	err := r.db.Preload("MasterBook").Where("id = ?", ID).First(masterAuthor).Error

	if err != nil {
		return masterAuthor, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) == true {
		return masterAuthor, errors.New("Data author tidak ditemukan")
	}

	masterAuthor.Name = input.Name
	masterAuthor.UpdatedAt = time.Now()

	err = r.db.Updates(&masterAuthor).Error

	if err != nil {
		return masterAuthor, err
	}

	return masterAuthor, nil
}

// func (r *repository) DeleteMasterBook(masterBook MasterBook, ID int) error {
// 	err := r.db.Delete(&masterBook, ID).Error

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
