package utils

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func GetID() uuid.UUID {
	return uuid.New()
}

func GetHashkey(key string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(key)))
}

func GetTimestamp() int64 {
	return time.Now().Unix()
}
