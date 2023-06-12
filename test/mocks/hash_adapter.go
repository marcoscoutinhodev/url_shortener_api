package mocks

import (
	"github.com/stretchr/testify/mock"
)

type HashAdapterMock struct {
	mock.Mock
}

func (h *HashAdapterMock) Generate(plaintext string) string {
	args := h.Called(plaintext)
	return args.String(0)
}

func (h *HashAdapterMock) Compare(hash, plaintext string) bool {
	args := h.Called(hash, plaintext)
	return args.Bool(0)
}
