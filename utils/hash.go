package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	
	if err != nil {
		return "", err
	}

	return string(bcryptedPassword), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}	