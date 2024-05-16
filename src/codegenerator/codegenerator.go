package codegenerator

import "crypto/rand"

func GenerateCode() (string, error) {
	posibleCaracters := "abcdefghijklmnopqrstuvwxyz1234567890"

	bytes := make([]byte, 6)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = posibleCaracters[b%byte(len(posibleCaracters))]
	}

	return string(bytes), nil
}
