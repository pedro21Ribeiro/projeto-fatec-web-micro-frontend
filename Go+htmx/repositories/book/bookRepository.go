package bookRepository

import (
	"time"

	"pedro21ribeiro.com/dbConector"
	bookModels "pedro21ribeiro.com/models/book"
)

func CreateBook(name, isbn string, date time.Time, author, publisher string) (uint, error) {
	db, err := dbConector.GetDatabase()
	if(err!=nil){
		return 0, err
	}
	book := bookModels.NewBook(name, isbn, date, author, publisher)

	err = db.Create(&book).Error; if err != nil {
		return 0, err 
	}

	return book.ID, err
}

func FindBookByIsbn(isbn string) *bookModels.Book{
	db, err := dbConector.GetDatabase()
	if(err!=nil){
		return nil
	}

	book := bookModels.Book{}

	err = db.Where("isbn = ?", isbn).First(&book).Error;  if err != nil {
		return nil	
	}

	return &book 
}