package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	secretloader "pedro21ribeiro.com/secretLoader"
)

var key []byte

func getKey() []byte{
	if (key == nil){
		secret := secretloader.GetAuthEnv().SECRET
		key = []byte(secret)
	}

	return key
}

func GenerateToken(id uint) (string, error) {

	tokenCreator := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": id,
			"exp": time.Now().Add(time.Hour * 8).Unix(),
		},
	)

	token, err := tokenCreator.SignedString(getKey()); if err != nil  {
		return "", err	
	}

	return token, nil
}

func GetTokenData(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString,
		func(t *jwt.Token) (any, error) {
			return getKey(), nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid Token from string")
	}

	return token, nil

}

func GetTokenClaims(t *jwt.Token) (jwt.MapClaims, error) {
	if claims, ok := t.Claims.(jwt.MapClaims); ok {
		return claims, nil
	} else {
		return nil, fmt.Errorf("error in converting the claims")
	}
}
