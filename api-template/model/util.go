package model

import (
	gonanoid "github.com/matoous/go-nanoid"
)

const letters = "123456789abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

func GenerateUniqueNanoId(length int) (string, error) {
	return gonanoid.Generate(letters, length)
}
