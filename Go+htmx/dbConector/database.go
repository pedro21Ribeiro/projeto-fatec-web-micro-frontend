package dbConector

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	secretloader "pedro21ribeiro.com/secretLoader"
)


func GetDatabase() (*gorm.DB, error){
	env := secretloader.LoadEnv()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		env.DBHost,
		env.DBUser,
		env.DBPassword,
		env.DBName,
		env.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}