package mocks

import "github.com/stretchr/testify/mock"

type EncryptAdapterMock struct {
	mock.Mock
}

func (e *EncryptAdapterMock) GenerateToken(payload map[string]interface{}, minutesToExpire uint) string {
	args := e.Called(payload, minutesToExpire)
	return args.String(0)
}
