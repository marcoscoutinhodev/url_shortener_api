package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnNilError(t *testing.T) {
	urlEntity := NewURLEntity("https://www.google.com", "")

	err := urlEntity.OriginalURLValidator()
	assert.Nil(t, err)
}

func TestShouldReturnInvalidOriginalURLError(t *testing.T) {
	urlEntity := NewURLEntity("google.com", "")

	err := urlEntity.OriginalURLValidator()
	assert.Equal(t, err.Error(), "original url is invalid")
}
