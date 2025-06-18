package userService

import (
	"errors"
	"pedro21ribeiro.com/auth/bcrypt"
	"pedro21ribeiro.com/auth/jwt"
	"pedro21ribeiro.com/repositories/user"
)

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