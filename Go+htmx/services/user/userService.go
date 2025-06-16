package userService

import (
	"errors"
	"net/http"

	"pedro21ribeiro.com/auth/bcrypt"
	"pedro21ribeiro.com/auth/jwt"
	"pedro21ribeiro.com/dtos/user"
	"pedro21ribeiro.com/repositories/user"
)

func GetAllUsers() (userDto.UsersDTO, error, int) {
	users, err := userRepository.GetAllUsers(); if err!=nil {
		return nil, err, http.StatusNotFound
	}

	var usersDTO userDto.UsersDTO

	for _, user := range *users {
		usersDTO = append(usersDTO, userDto.NewUserDTO(user.Name, user.Email))
	}

	return usersDTO, nil, http.StatusOK
}

func Login(email, password string) (string, error){
	returnError := errors.New("password or email is invalid")
	
	user, err := userRepository.GetUserByEmail(email)
	if err != nil {
		return "", returnError
	}

	IsPasswordCorrect := bcrypt.IsPasswordCorrect(password, user.Password)
	if(!IsPasswordCorrect){
		return "", returnError
	}

	token, err := jwt.GenerateToken(user.ID); if err != nil {
		return "", err
	}

	//TODO: Logica de Token JWT
	return token, nil
}