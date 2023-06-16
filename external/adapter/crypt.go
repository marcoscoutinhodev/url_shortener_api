package adapter

import (
	"crypto/rand"
	"fmt"
)

type CryptoAdapter struct{}

func NewCryptoAdapter() *CryptoAdapter {
	return &CryptoAdapter{}
}

func (c CryptoAdapter) GenerateRandomBytes() string {
	buf := make([]byte, 6)
	if _, err := rand.Read(buf); err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x", buf)
}
