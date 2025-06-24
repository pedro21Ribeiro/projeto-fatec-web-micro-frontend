package bookDto

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

func NewFormData() FormData {
	return FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

func (f *FormData) Populate(name, isbn, date, author, publisher string) {
	f.Values["Name"] = name
    f.Values["Isbn"] = isbn
    f.Values["Date"] = date
    f.Values["Author"] = author
    f.Values["Publisher"] = publisher
}