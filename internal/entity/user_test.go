package entity

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnNilErrorOnRegistrationValidator(t *testing.T) {
	name := "Lorem Ipsum"
	email := "LOreM@ipsum.com"
	password := "Abc1234!"

	user := NewUserEntity(name, email, password)
	err := user.RegistrationValidator()

	assert.NotNil(t, user)
	assert.Nil(t, err)

	assert.Equal(t, strings.ToUpper(name), user.Name)
	assert.Equal(t, strings.ToLower(email), user.Email)
	assert.Equal(t, password, user.Password)
}

func TestShouldReturnNameErrorOnRegistrationValidator(t *testing.T) {
	name := "Lore"
	email := "LOreM@ipsum.com"
	password := "Abc1234!"

	user := NewUserEntity(name, email, password)
	err := user.RegistrationValidator()
	assert.NotNil(t, err)

	user = NewUserEntity("", email, password)
	err = user.RegistrationValidator()
	assert.NotNil(t, err)
}

func TestShouldReturnEmailErrorOnRegistrationValidator(t *testing.T) {
	name := "Lorem Ipsum"
	email := "lorem@ipsumcom"
	password := "Abc1234!"

	user := NewUserEntity(name, email, password)
	err := user.RegistrationValidator()
	assert.NotNil(t, err)

	user = NewUserEntity(name, "", password)
	err = user.RegistrationValidator()
	assert.NotNil(t, err)
}

func TestShouldReturnPasswordErrorOnRegistrationValidator(t *testing.T) {
	name := "Lorem Ipsum"
	email := "lorem@ipsum.com"
	password := "Abc1234"

	user := NewUserEntity(name, email, password)
	err := user.RegistrationValidator()
	assert.NotNil(t, err)

	user = NewUserEntity(name, email, "")
	err = user.RegistrationValidator()
	assert.NotNil(t, err)
}

func TestShouldReturnNilErrorOnAuthenticationValidator(t *testing.T) {
	email := "LOreM@ipsum.com"
	password := "Abc1234!"

	user := NewUserEntity("", email, password)
	err := user.AuthenticationValidator()

	assert.NotNil(t, user)
	assert.Nil(t, err)

	assert.Equal(t, strings.ToUpper(""), user.Name)
	assert.Equal(t, strings.ToLower(email), user.Email)
	assert.Equal(t, password, user.Password)
}

func TestShouldReturnEmailErrorOnAuthenticationValidator(t *testing.T) {
	email := "lorem@ipsumcom"
	password := "Abc1234!"

	user := NewUserEntity("", email, password)
	err := user.AuthenticationValidator()
	assert.NotNil(t, err)

	user = NewUserEntity("", "", password)
	err = user.AuthenticationValidator()
	assert.NotNil(t, err)
}

func TestShouldReturnPasswordErrorOnAuthenticationValidator(t *testing.T) {
	email := "lorem@ipsum.com"
	password := "Abc1234"

	user := NewUserEntity("", email, password)
	err := user.AuthenticationValidator()
	assert.NotNil(t, err)

	user = NewUserEntity("", email, "")
	err = user.AuthenticationValidator()
	assert.NotNil(t, err)
}
