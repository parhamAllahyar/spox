package authenticator

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"

)

// TODO: get  secret key from config map
type JWT struct {
	ID uuid.UUID
}

func DecodeUserToken(userToken string) (JWT, error) {
	log.Println(userToken)
	token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
		//TODO:  read key from k8s secret
		return []byte("your key"), nil
	})

	if err == nil && token.Valid {
		claims, _ := token.Claims.(jwt.MapClaims)
		return JWT{
			ID: uuid.MustParse(claims["id"].(string)),
		}, nil
	} else {
		return JWT{}, err
	}
}

func DecodeAdminToken(userToken string) (JWT, error) {
	log.Println(userToken)
	token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
		//TODO:  read key from k8s secret
		return []byte("your key"), nil
	})

	if err == nil && token.Valid {
		claims, _ := token.Claims.(jwt.MapClaims)
		return JWT{
			ID: uuid.MustParse(claims["id"].(string)),
		}, nil
	} else {
		return JWT{}, err
	}
}
