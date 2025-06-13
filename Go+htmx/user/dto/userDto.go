package dto

type UserDTO struct {
	Name string
	Email string
}

type UsersDTO = []UserDTO

func NewUserDTO(name, email string) UserDTO{
	return UserDTO{
		Name: name,
		Email: email,
	}
}