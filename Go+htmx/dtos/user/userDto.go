package userDto

type FormError struct {
	Error string
}


func NewError() FormError{
	return FormError{}
}