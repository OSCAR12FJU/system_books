package repository

import (
	"database/sql"
	"homework-apirest/model"
)

type BookRepository struct {
	DB *sql.DB
}

func (repo *BookRepository) SaveBook(books *model.Books) (int, error) {

	query := `INSERT INTO books (name, author, pages, description, published, image) VALUES ($1,$2,$3,$4,$5,$6) RETURNING book_id`
	var lastInsertID int

	err := repo.DB.QueryRow(query, books.Name, books.Author, books.Pages, books.Description, books.Published, books.Image).Scan(&lastInsertID)
	if err != nil {
		return 0, nil
	}
	return int(lastInsertID), nil
	// result, err := repo.DB.Exec(`INSERT INTO books (name, author, pages, description, published, image) VALUES (?,?,?,?,?,?)`, books.Name, books.Author, books.Pages, books.Description, books.Published, books.Image)
	// if err != nil {
	// 	return 0, nil
	// }
	// lastInsertID, err := result.LastInsertId()
	// if err != nil {
	// 	return 0, nil
	// }
	// return int(lastInsertID), nil
}

func (repo *BookRepository) GetAllBooks() ([]model.Books, error) {
	rows, err := repo.DB.Query("SELECT book_id, name, author,image, status FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []model.Books

	for rows.Next() {
		var book model.Books
		err := rows.Scan(&book.ID, &book.Name, &book.Author, &book.Image, &book.Status)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (repo *BookRepository) UpdateBook(book model.Books) error {
	query := "UPDATE books SET name= $1, author = $2, image= $3 WHERE book_id = $4"
	_, err := repo.DB.Exec(query, book.Name, book.Author, book.Published, book.ID)
	return err
}

func (repo *BookRepository) GetBookByID(id int) (*model.Books, error) {
	var book model.Books
	err := repo.DB.QueryRow("SELECT book_id, name, author, image, status FROM books WHERE book_id = $1", id).Scan(&book.ID, &book.Name, &book.Author, &book.Published, &book.Status)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (repo *BookRepository) GetBookByName(name string) (*model.Books, error) {
	var book model.Books
	err := repo.DB.QueryRow("SELECT book_id, name, author, description, pages, published,image, status FROM books WHERE name = $1", name).Scan(&book.ID, &book.Name, &book.Author, &book.Description, &book.Pages, &book.Published, &book.Image, &book.Status)
	if err != nil {
		return nil, err
	}
	return &book, nil
}
