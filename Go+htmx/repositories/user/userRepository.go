package userRepository

import (
	"errors"
	"fmt"
	"time"

	"pedro21ribeiro.com/dbConector"
	"pedro21ribeiro.com/models/user"
)


func GetFirstUser() (*userModels.User, error){
	db, err := dbConector.GetDatabase()
	if(err!=nil){
		return nil, err
	}

	var user userModels.User

	res := db.First(&user).Error

	if(res!=nil){
		return nil, errors.New("nenhum usuário encontrado")
	}

	return &user, nil
}

func GetAllUsers() (*userModels.Users, error){
	db, err := dbConector.GetDatabase()
	if(err!=nil){
		return nil, err
	}
	
	var users userModels.Users

	res := db.Find(&users).Error; if (res != nil){
		return nil, errors.New("error when queryng the database")
	}

	return &users, nil
}

func GetUserByEmail(email string) (*userModels.User, error){
	db, err := dbConector.GetDatabase()

	if(err!=nil){
		return nil, err
	}

	var user userModels.User

	querry:= time.Now()
	res := db.First(&user ,"email = ?", email).Error; if res != nil {
		return nil, err
	}
	timeToQuerry:= time.Since(querry)
	fmt.Printf("\n\nTime to querry db was: %s\n", timeToQuerry)

	return &user, nil
}