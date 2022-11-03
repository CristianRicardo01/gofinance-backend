package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)
const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(number int) string {
	var sb strings.Builder
	K := len(alphabet)
	for i := 0; i < number; i++ {
		c := alphabet[rand.Intn(K)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomEmail(number int ) string {
	return fmt.Sprintf("%s@email.com", RandomString(number))
}