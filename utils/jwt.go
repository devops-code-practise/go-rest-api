package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "your_secret_key"

func GenerateJWT(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userid": userId,
		"exp":  time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}	


func VerifyToken(tokenString string)  error {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return  err
	}

	if !parsedToken.Valid {
		return  jwt.ErrSignatureInvalid
	}

	return nil

	// claims, err := parsedToken.Claims.(jwt.MapClaims)

	// if err != nil {
	// 	return err
	// }	
	// email := claims["email"].(string)
	// userId := int64(claims["userid"].(float64))

	// if email == "" || userId == 0 {
	// 	return  jwt.ErrSignatureInvalid
	// }
}