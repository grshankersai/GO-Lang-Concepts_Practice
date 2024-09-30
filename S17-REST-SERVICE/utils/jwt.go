package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secrectKey = "supersecret"

func GenerateToken(email string,userId int64) (string,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"email":email,
		"userId":userId,
		"exp":time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secrectKey))
}

func VerifyToken(token string) (int64,error){

	parsedToken , err := jwt.Parse(token, func(token *jwt.Token) (any,error){
		_,ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok{
			return nil,errors.New("Unexpected Sigining method")
		}
		return []byte(secrectKey),nil
	})

	if err != nil{
		return 0,errors.New("Couldnt parse token")
	}

	tokenIsValid := parsedToken.Valid

	if(! tokenIsValid){
		return 0,errors.New("Invalid Token")
	}

	claims,ok := parsedToken.Claims.(jwt.MapClaims)

	if(!ok){
		return 0,errors.New("Invalid Token claims.")
	}

	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return userId,nil

}