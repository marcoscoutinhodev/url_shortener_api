package mocks

import (
	"context"

	"github.com/marcoscoutinhodev/url_shortener_api/internal/entity"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (u *UserRepositoryMock) IsEmailRegistered(ctx context.Context, email string) bool {
	args := u.Called(ctx, email)
	return args.Bool(0)
}

func (u *UserRepositoryMock) CreateUser(ctx context.Context, user *entity.UserEntity) {
	u.Called(ctx, user)
}

func (u *UserRepositoryMock) LoadUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
	args := u.Called(ctx, email)
	return args.Get(0).(*entity.UserEntity), args.Error(1)
}
