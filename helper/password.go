package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (*string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	hashedStr := string(hashedPassword)

	if err != nil {
		return nil, err
	}

	return &hashedStr, nil
}

func CheckHashedPassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}
