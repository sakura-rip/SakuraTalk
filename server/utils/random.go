package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
)

func GenerateUUID() string {
	uid, _ := uuid.NewRandom()
	return fmt.Sprintf("%s", md5.Sum([]byte(uid.String())))
}
