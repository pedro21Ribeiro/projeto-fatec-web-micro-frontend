package bookService

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	bookDto "pedro21ribeiro.com/dtos/book"
	bookRepository "pedro21ribeiro.com/repositories/book"
)

func CreateBook(name, isbn, dateString, author, publisher, imgUrl string) *bookDto.FormData {
	var shouldSave bool
	formData := bookDto.NewFormData()
	formData.Populate(name, isbn, dateString, author, publisher, imgUrl)

	// Validate form data
	shouldSave = validateFormData(formData, name, isbn, dateString, author, publisher, imgUrl)

	if shouldSave {
		date, _ := time.Parse("2006-01-02", dateString) // Error already handled, ignoring.
		bookRepository.CreateBook(name, isbn, date, author, publisher, imgUrl)
		formData.Values["Succes"] = "Cadastro concluido com sucesso!"
	}
	return &formData
}

// validateFormData validates the form data and populates errors in the FormData.
// It returns true if the data is valid, false otherwise.
func validateFormData(formData bookDto.FormData, name, isbn, dateString, author, publisher, imgUrl string) bool {
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

	//TODO: make verification if it's possible to do a GET request to the provided ImgUrl if the type is .jpg .png .jpeg .webp
	// Validate imgUrl
	if imgUrl != "" {
		validImageTypes := []string{".jpg", ".png", ".jpeg", ".webp", ".gif"}
		isSupported := false

		for _, suffix := range validImageTypes {
			if strings.HasSuffix(strings.ToLower(imgUrl), suffix) {
				isSupported = true
				break
			}
		}

		if isSupported {
			resp, err := http.Get(imgUrl)
			if err != nil {
				formData.Errors["ImgUrl"] = "Invalid Image URL: Could not connect to the URL."
				formData.Values["ImgUrl"] = ""
				shouldSave = false
			} else {
				defer resp.Body.Close()
				if resp.StatusCode < 200 || resp.StatusCode >= 300 {
					formData.Errors["ImgUrl"] = "Invalid Image URL: URL returned status code " + resp.Status + "."
					formData.Values["ImgUrl"] = ""
					shouldSave = false
				}
			}
		} else {
			formData.Errors["ImgUrl"] = "The image must be a .jpg, .png, .jpeg, .webp or .gif"
			formData.Values["ImgUrl"] = ""
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

	// Validate isbn
	if imgUrl == "" {
		formData.Errors["ImgUrl"] = "Img Url is empty"
		formData.Values["ImgUrl"] = imgUrl
		shouldSave = false
	}

	return shouldSave
}
