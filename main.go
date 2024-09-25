package main

import (
	// "database/sql"

	"homework-apirest/controller"
	"homework-apirest/repository"
	"homework-apirest/service"
	"log"

	"homework-apirest/util"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	db, err := util.CreateConnection()
	if err != nil {
		log.Fatalf("Error al establecer conexión a la base de datos: %v", err)
	}
	defer db.Close()

	//Books
	bookRepo := &repository.BookRepository{DB: db}
	bookService := &service.BookService{Repo: bookRepo}
	bookController := &controller.BookController{Service: bookService}

	//Contact
	contactRepo := &repository.ContactRepository{DB: db}
	contactService := &service.ContactService{Repo: contactRepo}
	contactController := &controller.ContactController{Service: contactService}

	//User New
	userRepo := &repository.UserRepository{DB: db}
	userService := &service.UserService{Repo: userRepo}
	userController := &controller.UserController{Service: userService}

	//User Login
	userLoginRepo := &repository.UserLoginRepository{DB: db}
	userLoginService := &service.UserLoginService{Repo: userLoginRepo}
	userLoginController := &controller.UserLoginController{Service: userLoginService}

	router := mux.NewRouter()

	//Routes Peticiones Book
	router.HandleFunc("/create-book", bookController.CreateBookNew).Methods("POST")
	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))
	// router.HandleFunc("/get-books", bookController.GetBooks).Methods("GET")
	// router.HandleFunc("/get-book", bookController.GetBooks).Methods("GET")
	router.HandleFunc("/search-book", bookController.SearchBooks).Methods("GET")
	router.HandleFunc("/get-books/{id}", bookController.GetBookByID).Methods("GET")
	//Routes Book Page
	router.HandleFunc("/books/{bookName}", bookController.GetBookByName).Methods("GET")
	router.HandleFunc("/put-books/{id}", bookController.UpdateBook).Methods("PUT")

	//Routes User Login
	// router.HandleFunc("/login", userLoginController.Login).Methods("POST")
	router.HandleFunc("/login", userLoginController.Login).Methods("POST")
	router.HandleFunc("/logout", userLoginController.Logout).Methods("GET")

	//New Msj contact
	router.HandleFunc("/create-contact", contactController.CreateContactNew).Methods("POST")
	//Routes New User
	router.HandleFunc("/create-user", userController.CreateUserNew).Methods("POST")
	router.HandleFunc("/get-users", userController.GetUsers).Methods("GET")

	handler := cors.Default().Handler(router)
	headersOK := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	orginsOK := handlers.AllowedOrigins([]string{"*"})
	methodsOK := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	http.ListenAndServe(":8081", handler)
	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(orginsOK, headersOK, methodsOK)(router)))

}
