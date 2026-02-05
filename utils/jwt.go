package utils

import (
	"errors"
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

func VerifyJWT(token string) error{
	parseToken , err := jwt.Parse(token , func(TokenJw *jwt.Token) (interface{} , error){
	_, ok := TokenJw.Method.(*jwt.SigningMethodHMAC) // this is to check the signing method the signing method used in generating and verifying must be same
		if !ok {
			return nil , errors.New("Unexpected signing method")
		}

		return []byte(jwtSecret) , nil
	})

	if err != nil { 
		return errors.New("Could not parse token")
	}

	tokenIsValid := parseToken.Valid  // to check if the token is valid or not
		if !tokenIsValid {
			return errors.New("Invalid token")
		}


	// Extracting the email and password from the token claims	

		// claims , ok := parseToken.Claims.(jwt.MapClaims)	// this is type assertion to check if the claims are of type MapClaims
		// 	if !ok {
		// 			return errors.New("Could not extract claims")
		// 	}

		// id, ok:= claims["userId"].(int64) // this is to extract the userId from the claims and type assertion to check if it is of type int64	
		// email , ok := claims["email"].(string) // this is to extract the email from the claims map  and type assertion to check if it is of type string

	return nil
}