package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error) // string adalah token yang dihasilkan
}

type jwtService struct {
}

var SECRET_KEY = []byte("inirahasia") // buatlah di .env

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signToken, err
	}
	return signToken, nil
}

func NewService() *jwtService {
	return &jwtService{}
}
 