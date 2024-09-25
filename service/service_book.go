package service

import (
	"homework-apirest/model"
	"homework-apirest/repository"
)

type BookService struct {
	Repo *repository.BookRepository
}

func (service *BookService) CreateBook(books *model.Books) (*model.Books, error) {
	id, err := service.Repo.SaveBook(books)
	if err != nil {
		return nil, err
	}
	books.ID = id
	return books, nil
}

func (service *BookService) GetBooks() ([]model.Books, error) {
	books, err := service.Repo.GetAllBooks()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (service *BookService) SearchBook() ([]model.Books, error) {
	books, err := service.Repo.SearchBook()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (service *BookService) UpdateBook(book model.Books) (model.Books, error) {
	err := service.Repo.UpdateBook(book)
	return book, err
}

func (service *BookService) GetBookByID(id int) (*model.Books, error) {
	return service.Repo.GetBookByID(id)
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{Repo: repo}
}

func (service *BookService) GetBookByName(name string) (*model.Books, error) {
	return service.Repo.GetBookByName(name)
}
