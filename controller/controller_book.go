package controller

import (
	"database/sql"
	"encoding/json"
	"log"

	"fmt"
	"homework-apirest/model"
	"homework-apirest/service"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
)

type BookController struct {
	Service *service.BookService
}

func (c *BookController) CreateBookNew(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obtener los valores del formulario
	name := r.FormValue("name")
	author := r.FormValue("author")
	pages := r.FormValue("pages")
	description := r.FormValue("description")
	published := r.FormValue("published")

	// Manejar el archivo de imagen
	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	uploadDir := "uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, os.ModePerm)
	}

	// Guardar el archivo de imagen en el servidor
	filePath := filepath.Join(uploadDir, handler.Filename)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convertir las páginas a int
	pagesInt, err := strconv.Atoi(pages)
	if err != nil {
		http.Error(w, "Invalid number of pages", http.StatusBadRequest)
		return
	}
	serverURL := "http://localhost:8080"

	imageURL := fmt.Sprintf("%s/%s", serverURL, filePath)

	// Crear el objeto de libro
	book := &model.Books{
		Name:        name,
		Author:      author,
		Pages:       pagesInt,
		Description: description,
		Published:   published,
		Image:       imageURL,
		// Guardar el nombre del archivo
	}
	log.Printf("Book: %+v", book)

	// Guardar el libro en la base de datos
	newBook, err := c.Service.CreateBook(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Responder con el nuevo libro
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newBook)
}

func (c *BookController) SearchBooks(w http.ResponseWriter, r *http.Request) {
	books, err := c.Service.SearchBook()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (c *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var book model.Books
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Información invalida", http.StatusBadRequest)
		return
	}
	bookID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalido ID", http.StatusBadRequest)
		return
	}
	book.ID = bookID

	updatedBook, err := c.Service.UpdateBook(book)
	if err != nil {
		http.Error(w, "Error updating book", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBook)
}

func (c *BookController) GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Id invalido", http.StatusBadRequest)
		return
	}

	book, err := c.Service.GetBookByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Libro no encontrado", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (c *BookController) GetBookByName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	nameBook, ok := params["bookName"]

	if !ok {
		http.Error(w, "Nombre no valido", http.StatusBadRequest)
		return
	}

	book, err := c.Service.GetBookByName(nameBook)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Libro no encontrado", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
