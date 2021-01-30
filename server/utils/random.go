package utils

import (
	"github.com/google/uuid"
	"strings"
)

func GenerateUUID() string {
	uid, _ := uuid.NewRandom()
	return strings.Join(strings.Split(uid.String(), "-"), "")
}
