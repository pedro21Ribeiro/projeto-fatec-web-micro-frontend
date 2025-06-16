package dbConector

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	secretloader "pedro21ribeiro.com/secretLoader"
)

var db *gorm.DB

func InstanciateDB() (error) {
	var err error
	env := secretloader.GetDBEnv()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		env.DBHost,
		env.DBUser,
		env.DBPassword,
		env.DBName,
		env.DBPort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return err
}

func GetDatabase() (*gorm.DB, error){
	if(db == nil){
		fmt.Printf("Instanciating DB\n")
		err := InstanciateDB(); if err != nil {
			return nil, err
		}
	}

	return db, nil
}