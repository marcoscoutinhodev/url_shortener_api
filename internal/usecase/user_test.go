package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/marcoscoutinhodev/url_shortener_api/internal/dto"
	"github.com/marcoscoutinhodev/url_shortener_api/internal/entity"
	"github.com/marcoscoutinhodev/url_shortener_api/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Create User
func TestShouldReturnAnErrorOnUserEntityRegistraionValidator(t *testing.T) {
	ctx := context.Background()
	ch := make(chan UseCaseResponse)

	usecase := NewUserUseCase(
		&mocks.UserRepositoryMock{},
		&mocks.HashAdapterMock{},
		&mocks.EncryptAdapterMock{},
	)
	go usecase.CreateUser(ctx, ch, dto.UserInput{
		Name:     "Lorem",
		Email:    "lorem@ipsum.com",
		Password: "lorem",
	})

	response := <-ch
	assert.Equal(t, response.Code, 400)
	assert.False(t, response.Success)
	assert.Equal(t, response.Data, "password should be of 8 characters long")
	close(ch)
}

func TestShouldReturnAnErroIfProvidedEmailIsAlreadyRegistered(t *testing.T) {
	ctx := context.Background()
	ch := make(chan UseCaseResponse)

	userRepositoryMock := mocks.UserRepositoryMock{}
	userRepositoryMock.On("IsEmailRegistered", ctx, "lorem@ipsum.com").Return(true)

	usecase := NewUserUseCase(
		&userRepositoryMock,
		&mocks.HashAdapterMock{},
		&mocks.EncryptAdapterMock{},
	)
	go usecase.CreateUser(ctx, ch, dto.UserInput{
		Name:     "Lorem",
		Email:    "lorem@ipsum.com",
		Password: "Lorem123!",
	})

	response := <-ch
	assert.Equal(t, response.Code, 400)
	assert.False(t, response.Success)
	assert.Equal(t, response.Data, "email already registered")

	userRepositoryMock.AssertExpectations(t)

	close(ch)
}

func TestShouldReturnNilErrorToCreateUser(t *testing.T) {
	ctx := context.Background()
	ch := make(chan UseCaseResponse)

	userRepositoryMock := mocks.UserRepositoryMock{}
	userRepositoryMock.On("IsEmailRegistered", ctx, "lorem@ipsum.com").Return(false)
	userRepositoryMock.On("CreateUser", ctx, &entity.UserEntity{
		Name:     "LOREM",
		Email:    "lorem@ipsum.com",
		Password: "hashed password",
	})
	hashAdapterMock := mocks.HashAdapterMock{}
	hashAdapterMock.On("Generate", "Lorem123!").Return("hashed password")

	usecase := NewUserUseCase(
		&userRepositoryMock,
		&hashAdapterMock,
		&mocks.EncryptAdapterMock{},
	)
	go usecase.CreateUser(ctx, ch, dto.UserInput{
		Name:     "Lorem",
		Email:    "lorem@ipsum.com",
		Password: "Lorem123!",
	})

	response := <-ch
	assert.Equal(t, response.Code, 201)
	assert.True(t, response.Success)

	userRepositoryMock.AssertExpectations(t)
	hashAdapterMock.AssertExpectations(t)

	close(ch)
}

// Authenticate User
func TestShouldReturnAnErrorOnUserEntityAuthenticationValidator(t *testing.T) {
	ctx := context.Background()
	ch := make(chan UseCaseResponse)

	usecase := NewUserUseCase(
		&mocks.UserRepositoryMock{},
		&mocks.HashAdapterMock{},
		&mocks.EncryptAdapterMock{},
	)
	go usecase.AuthenticateUser(ctx, ch, dto.UserInput{
		Email:    "lorem@ipsum.com",
		Password: "lorem",
	})

	response := <-ch
	assert.Equal(t, response.Code, 400)
	assert.False(t, response.Success)
	assert.Equal(t, response.Data, "password should be of 8 characters long")
	close(ch)
}

func TestShouldReturnAnErroIfTheGivenEmailIsNotRegister(t *testing.T) {
	ctx := context.Background()
	ch := make(chan UseCaseResponse)

	userRepositoryMock := mocks.UserRepositoryMock{}
	userRepositoryMock.On("LoadUserByEmail", ctx, "lorem@ipsum.com").Return(&entity.UserEntity{}, errors.New("invalidd email or password"))

	usecase := NewUserUseCase(
		&userRepositoryMock,
		&mocks.HashAdapterMock{},
		&mocks.EncryptAdapterMock{},
	)
	go usecase.AuthenticateUser(ctx, ch, dto.UserInput{
		Email:    "lorem@ipsum.com",
		Password: "Lorem123!",
	})

	response := <-ch
	assert.Equal(t, 401, response.Code)
	assert.False(t, response.Success)
	assert.Equal(t, "invalid email or password", response.Data)

	userRepositoryMock.AssertExpectations(t)

	close(ch)
}

func TestShouldReturnNilErrorToAuthenticateUser(t *testing.T) {
	ctx := context.Background()
	ch := make(chan UseCaseResponse)

	mockedUser := &entity.UserEntity{
		ID:       primitive.NewObjectID(),
		Name:     "ANY_NAME",
		Email:    "ANY_EMAIL",
		Password: "hashedPassword",
	}
	userRepositoryMock := mocks.UserRepositoryMock{}
	userRepositoryMock.On("LoadUserByEmail", ctx, "lorem@ipsum.com").Return(mockedUser, nil)

	hashAdapterMock := mocks.HashAdapterMock{}
	hashAdapterMock.On("Compare", "hashedPassword", "Lorem123!").Return(true)

	encryptAdapterMock := mocks.EncryptAdapterMock{}
	encryptAdapterMock.On("GenerateToken", map[string]interface{}{
		"sub": mockedUser.ID,
	}, uint(15)).Return("access_token")

	usecase := NewUserUseCase(
		&userRepositoryMock,
		&hashAdapterMock,
		&encryptAdapterMock,
	)
	go usecase.AuthenticateUser(ctx, ch, dto.UserInput{
		Email:    "lorem@ipsum.com",
		Password: "Lorem123!",
	})

	response := <-ch
	assert.Equal(t, 200, response.Code)
	assert.True(t, response.Success)
	assert.NotNil(t, response.Data)

	userRepositoryMock.AssertExpectations(t)
	hashAdapterMock.AssertExpectations(t)
	encryptAdapterMock.AssertExpectations(t)

	close(ch)
}
