package usecase

import (
	"context"
	"encoding/base64"
	"testing"

	"github.com/marcoscoutinhodev/url_shortener_api/internal/dto"
	"github.com/marcoscoutinhodev/url_shortener_api/internal/entity"
	"github.com/marcoscoutinhodev/url_shortener_api/test/mocks"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnAnErrorIfOriginalURLValidatorFails(t *testing.T) {
	urlUseCase := NewURLUseCase(
		&mocks.URLRepositoryMock{},
		&mocks.URLCheckerAdapterMock{},
	)

	ch := make(chan UseCaseResponse)

	go urlUseCase.CreateShortURL(context.Background(), ch, &dto.ShortURLInput{
		OriginalURL: "invalid_url",
	}, "any_user_id")

	response := <-ch
	assert.Equal(t, response.Code, 400)
	assert.False(t, response.Success)
	assert.Equal(t, response.Data, "original url is invalid")
	close(ch)
}

func TestShouldReturnAnErrorIfOriginalURLIsNotSafe(t *testing.T) {
	ctx := context.Background()
	ch := make(chan UseCaseResponse)

	urlEnconded := base64.RawURLEncoding.EncodeToString([]byte("https://www.google.com"))
	urlCheckerAdapterMock := mocks.URLCheckerAdapterMock{}
	urlCheckerAdapterMock.On("IsURLSafe", ctx, urlEnconded).Return(false)

	urlUseCase := NewURLUseCase(
		&mocks.URLRepositoryMock{},
		&urlCheckerAdapterMock,
	)

	go urlUseCase.CreateShortURL(ctx, ch, &dto.ShortURLInput{
		OriginalURL: "https://www.google.com",
	}, "any_user_id")

	response := <-ch
	assert.Equal(t, response.Code, 400)
	assert.False(t, response.Success)
	assert.Equal(t, response.Data, "the provided URL is not secure")

	urlCheckerAdapterMock.AssertExpectations(t)

	close(ch)
}

func TestShouldReturnNilErrorToShortURL(t *testing.T) {
	ctx := context.Background()
	ch := make(chan UseCaseResponse)

	urlEnconded := base64.RawURLEncoding.EncodeToString([]byte("https://www.google.com"))

	urlEntity := entity.NewURLEntity("https://www.google.com", urlEnconded)
	urlRepositoryMock := mocks.URLRepositoryMock{}
	urlRepositoryMock.On("CreateShortURL", ctx, urlEntity, "any_user_id")

	urlCheckerAdapterMock := mocks.URLCheckerAdapterMock{}
	urlCheckerAdapterMock.On("IsURLSafe", ctx, urlEnconded).Return(true)

	urlUseCase := NewURLUseCase(
		&urlRepositoryMock,
		&urlCheckerAdapterMock,
	)
	go urlUseCase.CreateShortURL(ctx, ch, &dto.ShortURLInput{
		OriginalURL: "https://www.google.com",
	}, "any_user_id")

	response := <-ch
	assert.Equal(t, response.Code, 201)
	assert.True(t, response.Success)

	urlCheckerAdapterMock.AssertExpectations(t)
	urlCheckerAdapterMock.AssertExpectations(t)
	// urlRepositoryMock.AssertNumberOfCalls(t, "CreateShortURL", 1)
	close(ch)
}
