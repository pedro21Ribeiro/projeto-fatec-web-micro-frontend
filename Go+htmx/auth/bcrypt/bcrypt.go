package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
	"pedro21ribeiro.com/secretLoader"
)

func IsPasswordCorrect(password, chalenge string) bool{
	authVars := secretloader.GetAuthEnv()	
	password = password + authVars.SALT
	err := bcrypt.CompareHashAndPassword([]byte(chalenge), []byte(password))

	return err == nil
}