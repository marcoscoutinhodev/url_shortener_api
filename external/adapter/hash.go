package adapter

import (
	"golang.org/x/crypto/bcrypt"
)

type HashAdapter struct{}

func NewHashAdapter() *HashAdapter {
	return &HashAdapter{}
}

func (h HashAdapter) Generate(plaintext string) string {
	hashByte, err := bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashByte)
}

// Returns true on compare successful
func (h HashAdapter) Compare(hash, plaintext string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plaintext))
	return err == nil
}
