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
