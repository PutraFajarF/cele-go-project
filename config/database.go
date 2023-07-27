package config

import (
	"fmt"
	"log"
	"project-go/entities"
	"project-go/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Connect() (*gorm.DB, error) {
	dataConn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		AppConfig.Host,
		AppConfig.Username,
		AppConfig.Password,
		AppConfig.DatabaseName,
		AppConfig.Port,
	)
	DB, err = gorm.Open(postgres.Open(dataConn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	return DB, nil
}

func Migrate() error {
	err := DB.AutoMigrate(
		&user.User{},
		&entities.MasterAuthor{},
		&entities.MasterBook{},
		&entities.BookTransaction{},
	)

	if err != nil {
		return err
	}

	DB.Migrator().CreateConstraint(&entities.MasterBook{}, "MasterAuthors")
	DB.Migrator().CreateConstraint(&entities.MasterBook{}, "fk_master_books_authors")

	DB.Migrator().CreateConstraint(&entities.BookTransaction{}, "Users")
	DB.Migrator().CreateConstraint(&entities.BookTransaction{}, "fk_book_transactions_users")

	DB.Migrator().CreateConstraint(&entities.BookTransaction{}, "MasterBooks")
	DB.Migrator().CreateConstraint(&entities.BookTransaction{}, "fk_book_transactions_master_books")

	log.Println("Database Migration Completed...")

	return nil
}
