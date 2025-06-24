package bookService

import (
	"fmt"
	"time"

	bookDto "pedro21ribeiro.com/dtos/book"
	bookRepository "pedro21ribeiro.com/repositories/book"
)

func CreateBook(name, isbn, dateString, author, publisher string) *bookDto.FormData {
	var shouldSave bool
	formData := bookDto.NewFormData()
	formData.Populate(name, isbn, dateString, author, publisher)

	// Validate form data
	shouldSave = validateFormData(formData, name, isbn, dateString, author, publisher)

	if shouldSave {
		date, _ := time.Parse("2006-01-02", dateString) // Error already handled, ignoring.
		bookRepository.CreateBook(name, isbn, date, author, publisher)
	}
	return &formData
}

// validateFormData validates the form data and populates errors in the FormData.
// It returns true if the data is valid, false otherwise.
func validateFormData(formData bookDto.FormData, name, isbn, dateString, author, publisher string) bool {
	shouldSave := true

	// Check if ISBN already exists
	book := bookRepository.FindBookByIsbn(isbn)
	if book != nil {
		formData.Errors["Isbn"] = "ISBN already exists"
		formData.Values["Isbn"] = "" // Clear the ISBN field if it already exists
		shouldSave = false
	}

	// Validate date
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		fmt.Printf("\n\ndate %s is invalid\n", dateString)
		formData.Errors["Date"] = "Date is invalid"
		formData.Values["Date"] = "" // Clear the date field if it's invalid
		shouldSave = false
	} else {
		if date.After(time.Now()) { // Use date.After for clarity, its the same thing.
			formData.Errors["Date"] = "Date cannot be in the future"
			formData.Values["Date"] = "" // Clear the date field if it's invalid
			shouldSave = false
		}
	}

	// Validate name
	if name == "" {
		formData.Errors["Name"] = "Name is empty"
		formData.Values["Name"] = name
		shouldSave = false
	}

	// Validate author
	if author == "" {
		formData.Errors["Author"] = "Author is empty"
		formData.Values["Author"] = author
		shouldSave = false
	}

	// Validate publisher
	if publisher == "" {
		formData.Errors["Publisher"] = "Publisher is empty"
		formData.Values["Publisher"] = publisher
		shouldSave = false
	}

	// Validate isbn
	if isbn == "" {
		formData.Errors["Isbn"] = "Isbn is empty"
		formData.Values["Isbn"] = isbn
		shouldSave = false
	}

	return shouldSave
}
