package repository

import (
	"errors"

	"pedro21ribeiro.com/conector"
	"pedro21ribeiro.com/models"
)


func GetFirstUser() (*models.User, error){
	db, err := conector.GetDatabase()
	if(err!=nil){
		return nil, err
	}

	var user models.User

	res := db.First(&user).Error

	if(res!=nil){
		return nil, errors.New("nenhum usuário encontrado")
	}

	return &user, nil
}