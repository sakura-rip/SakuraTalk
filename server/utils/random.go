package utils

import (
	"github.com/google/uuid"
)

func GenerateUUID() string {
	uid, _ := uuid.NewRandom()
	return uid.String()
}
