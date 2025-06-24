package bookModels

import "time"

type Tabler interface {
	TableName() string
}

type Book struct {
	ID          uint
	Name        string `gorm:"column:nome_livro"`
	Isbn        string
	ReleaseDate time.Time `gorm:"column:data_lancamento"`
	Author      string    `gorm:"column:nome_autor"`
	Publisher   string    `gorm:"column:nome_editora"`
}

func (Book) TableName() string {
	return "books"
}

type Books = []Book

func NewBook(name, isbn string, date time.Time, author, publisher string) Book {
	return Book{
		Name:        name,
		Isbn:        isbn,
		ReleaseDate: date,
		Author:      author,
		Publisher:   publisher,
	}
}
