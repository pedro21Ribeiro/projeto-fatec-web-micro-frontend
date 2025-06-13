package service

import (
	"net/http"

	"pedro21ribeiro.com/user/dto"
	"pedro21ribeiro.com/user/repository"
)

func GetAllUsers() (dto.UsersDTO, error, int) {
	users, err := repository.GetAllUsers(); if err!=nil {
		return nil, err, http.StatusNotFound
	}

	var usersDTO dto.UsersDTO

	for _, user := range *users {
		usersDTO = append(usersDTO, dto.NewUserDTO(user.Name, user.Email))
	}

	return usersDTO, nil, http.StatusOK
}