package optional

import "golang.org/x/crypto/bcrypt"

func HashPass(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password + "СОЛЬ"), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckHash(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass + "СОЛЬ"))
	return err == nil
}