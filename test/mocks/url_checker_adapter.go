package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type URLCheckerAdapterMock struct {
	mock.Mock
}

func (u *URLCheckerAdapterMock) IsURLSafe(ctx context.Context, url string) bool {
	args := u.Called(ctx, url)
	return args.Bool(0)
}
