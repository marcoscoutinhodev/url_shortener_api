package usecase

import (
	"context"

	"github.com/marcoscoutinhodev/url_shortener_api/internal/dto"
	"github.com/marcoscoutinhodev/url_shortener_api/internal/entity"
)

type UserUseCase struct {
	UserRepository UserRepositoryInterface
	HashAdapter    HashAdapterInterface
	EncryptAdapter EncryptAdapterInterface
}

func NewUserUseCase(userRepository UserRepositoryInterface, hashAdapter HashAdapterInterface, encryptAdapter EncryptAdapterInterface) *UserUseCase {
	return &UserUseCase{
		UserRepository: userRepository,
		HashAdapter:    hashAdapter,
		EncryptAdapter: encryptAdapter,
	}
}

func (u UserUseCase) CreateUser(ctx context.Context, ch chan<- UseCaseResponse, userInput dto.UserInput) {
	defer RecoverPanic(ch, "CreateUser")()

	user := entity.NewUserEntity(userInput.Name, userInput.Email, userInput.Password)
	if err := user.RegistrationValidator(); err != nil {
		ch <- UseCaseResponse{
			Code:    400,
			Success: false,
			Data:    err.Error(),
		}
		return
	}

	if isRegistered := u.UserRepository.IsEmailRegistered(ctx, user.Email); isRegistered {
		ch <- UseCaseResponse{
			Code:    400,
			Success: false,
			Data:    "email already registered",
		}
		return
	}

	user.Password = u.HashAdapter.Generate(user.Password)

	u.UserRepository.CreateUser(ctx, user)

	ch <- UseCaseResponse{
		Code:    201,
		Success: true,
	}
}

func (u UserUseCase) AuthenticateUser(ctx context.Context, ch chan<- UseCaseResponse, userInput dto.UserInput) {
	defer RecoverPanic(ch, "AuthenticateUser")()

	user := entity.NewUserEntity(userInput.Name, userInput.Email, userInput.Password)
	if err := user.AuthenticationValidator(); err != nil {
		ch <- UseCaseResponse{
			Code:    400,
			Success: false,
			Data:    err.Error(),
		}
		return
	}

	userFromRepository, err := u.UserRepository.LoadUserByEmail(ctx, user.Email)
	if err != nil {
		ch <- UseCaseResponse{
			Success: false,
			Data:    "invalid email or password",
			Code:    401,
		}
		return
	}

	isPasswordValid := u.HashAdapter.Compare(userFromRepository.Password, user.Password)
	if !isPasswordValid {
		ch <- UseCaseResponse{
			Success: false,
			Data:    "invalid email or password",
			Code:    401,
		}
		return
	}

	token := u.EncryptAdapter.GenerateToken(map[string]interface{}{
		"sub": userFromRepository.ID,
	}, 15)

	ch <- UseCaseResponse{
		Success: true,
		Data: map[string]interface{}{
			"token": token,
			"user":  userFromRepository,
		},
		Code: 200,
	}
}
