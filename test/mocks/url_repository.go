package mocks

import (
	"context"

	"github.com/marcoscoutinhodev/url_shortener_api/internal/entity"
	"github.com/stretchr/testify/mock"
)

type URLRepositoryMock struct {
	mock.Mock
}

func (u *URLRepositoryMock) CreateShortURL(ctx context.Context, url *entity.URLEntity, userId string) {
	u.Called(ctx, url, userId)
}

func (u *URLRepositoryMock) GetOriginalURL(ctx context.Context, shortURL string) (*entity.URLEntity, error) {
	args := u.Called(ctx, shortURL)
	return args.Get(0).(*entity.URLEntity), args.Error(1)
}

func (u *URLRepositoryMock) ReportURL(ctx context.Context, urlID string) error {
	args := u.Called(ctx, urlID)
	return args.Error(0)
}

func (u *URLRepositoryMock) ActiveURL(ctx context.Context, userID, urlID string) error {
	args := u.Called(ctx, userID, urlID)
	return args.Error(0)
}
