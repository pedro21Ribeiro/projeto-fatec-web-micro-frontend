package repository

import (
	"errors"

	"pedro21ribeiro.com/dbConector"
	"pedro21ribeiro.com/user/models"
)


func GetFirstUser() (*models.User, error){
	db, err := dbConector.GetDatabase()
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

func GetAllUsers() (*models.Users, error){
	db, err := dbConector.GetDatabase()
	if(err!=nil){
		return nil, err
	}
	
	var users models.Users

	res := db.Find(&users).Error; if (res != nil){
		return nil, errors.New("error when queryng the database")
	}

	return &users, nil
}