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

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	fmt.Println("****", min, max)
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	fmt.Println("****", sb)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		fmt.Println("****", sb, c)

		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOnwner generate a random owner name
func RandomOwner() string {
	return RandomString(7)
}

// RandomAmount generate a random amount of money
func RandomBalance() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generate a random currency mode
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}