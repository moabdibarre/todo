package utils




import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashPass), nil
}

func ComparePassword(hashedpass, plainPass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedpass), []byte(plainPass))
	if err != nil{
		return err
	}
	return nil
}