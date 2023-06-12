package usecase

import (
	"context"

	"github.com/marcoscoutinhodev/url_shortener_api/internal/entity"
)

type UserRepositoryInterface interface {
	IsEmailRegistered(ctx context.Context, email string) bool
	CreateUser(ctx context.Context, user *entity.UserEntity)
	LoadUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error)
}

type HashAdapterInterface interface {
	Generate(plaintext string) string
	Compare(hash, plaintext string) bool
}

type EncryptAdapterInterface interface {
	GenerateToken(payload map[string]interface{}, minutesToExpire uint) string
}

type URLRepositoryInterface interface {
	CreateShortURL(ctx context.Context, url *entity.URLEntity, userId string)
}

type URLCheckerAdapterInterface interface {
	IsURLSafe(ctx context.Context, urlEncoded string) bool
}
