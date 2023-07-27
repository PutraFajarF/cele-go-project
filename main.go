package main

import (
	"log"
	"project-go/auth"
	"project-go/book_transaction"
	"project-go/config"
	"project-go/handler"
	"project-go/master_author"
	"project-go/master_book"
	"project-go/routes"
	"project-go/user"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadAppConfig()
	db, err := config.Connect()

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Print("Database sukses terkoneksi")

	err = config.Migrate()

	if err != nil {
		log.Fatal(err.Error())
	}

	// call repository
	userRepository := user.NewRepository(db)
	masterBookRepository := master_book.NewRepository(db)
	masterAuthorRepository := master_author.NewRepository(db)
	bookTransactionRepository := book_transaction.NewRepository(db)

	// call service
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	masterBookService := master_book.NewService(masterBookRepository)
	masterAuthorService := master_author.NewService(masterAuthorRepository)
	bookTransactionService := book_transaction.NewService(bookTransactionRepository)

	// call handler
	userHandler := handler.NewUserHandler(userService, authService)
	masterBookHandler := handler.NewMasterBookHandler(masterBookService)
	masterAuthorHandler := handler.NewMasterAuthorHandler(masterAuthorService)
	bookTransactionhandler := handler.NewBookTransactionHandler(bookTransactionService)

	// gin router
	router := gin.Default()

	// api versioning
	userApi := router.Group("/api/v1/user")
	masterApi := router.Group("/api/v1/master")
	transactionApi := router.Group("/api/v1/transaction")

	routes.UserRoutes(userApi, userHandler)
	routes.MasterBookRoutes(masterApi, masterBookHandler, db, authService, userService)
	routes.MasterAuthorRoutes(masterApi, masterAuthorHandler, db, authService, userService)
	routes.BookTransactionRoutes(transactionApi, bookTransactionhandler, db, authService, userService)

	router.Run()
}
