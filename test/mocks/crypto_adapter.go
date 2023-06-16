package mocks

import "github.com/stretchr/testify/mock"

type CryptoAdapterMock struct {
	mock.Mock
}

func (e *CryptoAdapterMock) GenerateRandomBytes() string {
	args := e.Called()
	return args.String(0)
}
