package usecase

import (
	"context"
	"encoding/base64"
	"errors"
	"testing"

	"github.com/marcoscoutinhodev/url_shortener_api/internal/dto"
	"github.com/marcoscoutinhodev/url_shortener_api/internal/entity"
	"github.com/marcoscoutinhodev/url_shortener_api/test/mocks"
	"github.com/stretchr/testify/assert"
)

// CreateShortURL
func TestShouldReturnAnErrorIfOriginalURLValidatorFails(t *testing.T) {
	urlUseCase := NewURLUseCase(
		&mocks.URLRepositoryMock{},
		&mocks.URLCheckerAdapterMock{},
		&mocks.CryptoAdapterMock{},
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
		&mocks.CryptoAdapterMock{},
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

func TestShouldReturnNilErrorToCreateShortURL(t *testing.T) {
	ctx := context.Background()
	ch := make(chan UseCaseResponse)

	urlEnconded := base64.RawURLEncoding.EncodeToString([]byte("https://www.google.com"))

	urlEntity := entity.NewURLEntity("https://www.google.com", "randombyte")
	urlRepositoryMock := mocks.URLRepositoryMock{}
	urlRepositoryMock.On("CreateShortURL", ctx, urlEntity, "any_user_id")

	urlCheckerAdapterMock := mocks.URLCheckerAdapterMock{}
	urlCheckerAdapterMock.On("IsURLSafe", ctx, urlEnconded).Return(true)

	cryptoAdapterMock := mocks.CryptoAdapterMock{}
	cryptoAdapterMock.On("GenerateRandomBytes").Return("randombyte")

	urlUseCase := NewURLUseCase(
		&urlRepositoryMock,
		&urlCheckerAdapterMock,
		&cryptoAdapterMock,
	)
	go urlUseCase.CreateShortURL(ctx, ch, &dto.ShortURLInput{
		OriginalURL: "https://www.google.com",
	}, "any_user_id")

	response := <-ch
	assert.Equal(t, response.Code, 201)
	assert.True(t, response.Success)

	urlCheckerAdapterMock.AssertExpectations(t)
	urlCheckerAdapterMock.AssertExpectations(t)
	urlRepositoryMock.AssertNumberOfCalls(t, "CreateShortURL", 1)
	close(ch)
}

// GetOriginalURL
func TestShouldReturnAnErrorIfShortURLIsNotFound(t *testing.T) {
	ctx := context.Background()
	ch := make(chan UseCaseResponse)

	urlRepositoryMock := mocks.URLRepositoryMock{}
	urlRepositoryMock.On("GetOriginalURL", ctx, "short_url").Return(&entity.URLEntity{}, errors.New("no matching url"))

	urlUseCase := NewURLUseCase(
		&urlRepositoryMock,
		&mocks.URLCheckerAdapterMock{},
		&mocks.CryptoAdapterMock{},
	)

	go urlUseCase.GetOriginalURL(ctx, ch, "short_url")

	response := <-ch
	assert.Equal(t, response.Code, 404)
	assert.False(t, response.Success)
	assert.Equal(t, response.Data, "no matching url")

	urlRepositoryMock.AssertExpectations(t)

	close(ch)
}

func TestShouldReturnSuccessOnGetOriginalURL(t *testing.T) {
	ctx := context.Background()
	ch := make(chan UseCaseResponse)

	urlRepositoryMock := mocks.URLRepositoryMock{}
	urlRepositoryMock.On("GetOriginalURL", ctx, "short_url").Return(&entity.URLEntity{}, nil)

	urlUseCase := NewURLUseCase(
		&urlRepositoryMock,
		&mocks.URLCheckerAdapterMock{},
		&mocks.CryptoAdapterMock{},
	)

	go urlUseCase.GetOriginalURL(ctx, ch, "short_url")

	response := <-ch
	assert.Equal(t, response.Code, 200)
	assert.True(t, response.Success)

	urlRepositoryMock.AssertExpectations(t)

	close(ch)
}

func TestShouldReturnAnErrorIfUrlIDIsNotFound(t *testing.T) {
	ctx := context.Background()
	ch := make(chan UseCaseResponse)

	urlRepositoryMock := mocks.URLRepositoryMock{}
	urlRepositoryMock.On("ReportURL", ctx, "url_id").Return(errors.New("no matching url"))

	urlUseCase := NewURLUseCase(
		&urlRepositoryMock,
		&mocks.URLCheckerAdapterMock{},
		&mocks.CryptoAdapterMock{},
	)

	go urlUseCase.ReportURL(ctx, ch, "url_id")

	response := <-ch
	assert.Equal(t, response.Code, 404)
	assert.False(t, response.Success)
	assert.Equal(t, response.Data, "no matching url")

	urlRepositoryMock.AssertExpectations(t)

	close(ch)
}

func TestShouldReturnSuccessOnReportURL(t *testing.T) {
	ctx := context.Background()
	ch := make(chan UseCaseResponse)

	urlRepositoryMock := mocks.URLRepositoryMock{}
	urlRepositoryMock.On("ReportURL", ctx, "url_id").Return(nil)

	urlUseCase := NewURLUseCase(
		&urlRepositoryMock,
		&mocks.URLCheckerAdapterMock{},
		&mocks.CryptoAdapterMock{},
	)

	go urlUseCase.ReportURL(ctx, ch, "url_id")

	response := <-ch
	assert.Equal(t, response.Code, 200)
	assert.True(t, response.Success)

	urlRepositoryMock.AssertExpectations(t)

	close(ch)
}
