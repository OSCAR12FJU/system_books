/*
package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	conectionString := "root:Fuerzaabasto1@@tcp(127.0.0.1:3306)/system_books"

	var err error

	db, err = sql.Open("mysql", conectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func InsertBook(name string, author string, status bool) error {
	query := "INSERT INTO books (name, author, status) VALUES (?,?)"
	_, err := db.Exec(query, name, author, status)
	return err
}
*/

// Petición bien echa de busqueda por ID

// func getBookByID(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	id := params["id"]

// 	var book Books

// 	err := db.QueryRow("SELECT book_id, name, author, status FROM books WHERE book_id = ?", id).Scan(&book.ID, &book.Name, &book.Author, &book.Image, &book.Status)

// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			http.Error(w, "Libro no encontrado", http.StatusNotFound)
// 		} else {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(book)
// }

//COODIGO UTIL PARA TRABJAR CON FORMS EL CUAL SOLAMENTE NOS TOMA TEXTO PARA SUBIR

// func (c *BookController) CreateBookNew(w http.ResponseWriter, r *http.Request) {
// 	var book model.Books

// 	err := json.NewDecoder(r.Body).Decode(&book)
// 	if err != nil {
// 		http.Error(w, "Datos invalidos", http.StatusBadRequest)
// 		return
// 	}

// 	newBook, err := c.Service.CreateBook(&book)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(newBook)
// }
