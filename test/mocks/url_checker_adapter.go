package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type URLCheckerAdapterMock struct {
	mock.Mock
}

func (u *URLCheckerAdapterMock) IsURLSafe(ctx context.Context, urlEncoded string) bool {
	args := u.Called(ctx, urlEncoded)
	return args.Bool(0)
}
