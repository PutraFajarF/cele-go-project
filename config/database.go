package config

import (
	"fmt"
	"log"
	"project-go/book_transaction"
	"project-go/master_author"
	"project-go/master_book"
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
		&master_author.MasterAuthor{},
		&master_book.MasterBook{},
		&book_transaction.BookTransaction{},
	)

	if err != nil {
		return err
	}

	DB.Migrator().CreateConstraint(&master_book.MasterBook{}, "MasterAuthors")
	DB.Migrator().CreateConstraint(&master_book.MasterBook{}, "fk_book_authors")

	DB.Migrator().CreateConstraint(&book_transaction.BookTransaction{}, "Users")
	DB.Migrator().CreateConstraint(&book_transaction.BookTransaction{}, "fk_transaction_users")

	DB.Migrator().CreateConstraint(&book_transaction.BookTransaction{}, "MasterBooks")
	DB.Migrator().CreateConstraint(&book_transaction.BookTransaction{}, "fk_transaction_books")

	log.Println("Database Migration Completed...")

	return nil
}
