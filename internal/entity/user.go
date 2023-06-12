package entity

import (
	"errors"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	missingNameError     = "name must be provided"
	invalidNameError     = "name requires at least 5 characters"
	missingEmailError    = "email must be provided"
	invalidEmailError    = "invalid email format"
	missingPasswordError = "password must be provided"
)

type UserDetailEntity struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	AllAccesses int64              `json:"all_accesses" bson:"all_accesses"`
	AllReports  int64              `json:"all_reports" bson:"all_reports"`
	UserID      primitive.ObjectID `json:"-" bson:"user_id"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty" bson:"updated_at"`
}

type UserEntity struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Name       string             `json:"name" bson:"name"`
	Email      string             `json:"email" bson:"email"`
	Password   string             `json:"-" bson:"password"`
	UserDetail UserDetailEntity   `json:"user_detail,omitempty"`
	CreatedAt  time.Time          `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at,omitempty" bson:"updated_at"`
}

func NewUserEntity(name, email, password string) *UserEntity {
	return &UserEntity{
		Name:     name,
		Email:    email,
		Password: password,
	}
}

func (u *UserEntity) RegistrationValidator() error {
	if u.Name == "" {
		return errors.New(missingNameError)
	}

	if n := strings.ReplaceAll(u.Name, " ", ""); len(n) < 5 {
		return errors.New(invalidNameError)
	}

	u.Name = strings.ToUpper(u.Name)

	if u.Email == "" {
		return errors.New(missingEmailError)
	}

	if !isEmailValid(u.Email) {
		return errors.New(invalidEmailError)
	}

	u.Email = strings.ToLower(u.Email)

	if u.Password == "" {
		return errors.New(missingPasswordError)
	}

	if err := validatePassword(u.Password); err != nil {
		return err
	}

	return nil
}

func (u *UserEntity) AuthenticationValidator() error {
	if u.Email == "" {
		return errors.New(missingEmailError)
	}

	if !isEmailValid(u.Email) {
		return errors.New(invalidEmailError)
	}

	u.Email = strings.ToLower(u.Email)

	if u.Password == "" {
		return errors.New(missingPasswordError)
	}

	if err := validatePassword(u.Password); err != nil {
		return err
	}

	return nil
}
