package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const jwtSecret = "this is super secret key fo generatin jwt"

func GenerateJWT(userId int64, email string) (string, error) {
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256 , jwt.MapClaims{
		"userId" : userId ,
		 "email" : email ,
		 "exp" : time.Now().Add(time.Hour * 2).Unix() , //  time must be in unix format
	})
	return token.SignedString([]byte(jwtSecret)) // this will give us the signed toke string , the original token is too more complex so we use signed string , the signed string says takes any input but in reality it must be in byte form 
}