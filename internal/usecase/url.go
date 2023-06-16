package usecase

import (
	"context"
	"encoding/base64"

	"github.com/marcoscoutinhodev/url_shortener_api/internal/dto"
	"github.com/marcoscoutinhodev/url_shortener_api/internal/entity"
)

type URLUseCase struct {
	url_repository      URLRepositoryInterface
	url_checker_adapter URLCheckerAdapterInterface
	crypto_adapter      CryptoAdapterInterface
}

func NewURLUseCase(urlRepository URLRepositoryInterface, urlCheckerAdapter URLCheckerAdapterInterface, cryptoAdapter CryptoAdapterInterface) *URLUseCase {
	return &URLUseCase{
		url_repository:      urlRepository,
		url_checker_adapter: urlCheckerAdapter,
		crypto_adapter:      cryptoAdapter,
	}
}

func (u URLUseCase) CreateShortURL(ctx context.Context, ch chan<- UseCaseResponse, urlInput *dto.ShortURLInput, userId string) {
	defer RecoverPanic(ch, "CreateShortURL")()

	url := entity.NewURLEntity(urlInput.OriginalURL, "")
	if err := url.OriginalURLValidator(); err != nil {
		ch <- UseCaseResponse{
			Success: false,
			Data:    err.Error(),
			Code:    400,
		}
		return
	}

	urlEnconded := base64.RawURLEncoding.EncodeToString([]byte(url.OriginalUrl))
	if isURLSafe := u.url_checker_adapter.IsURLSafe(ctx, urlEnconded); !isURLSafe {
		ch <- UseCaseResponse{
			Success: false,
			Data:    "the provided URL is not secure",
			Code:    400,
		}
		return
	}

	randomBytes := u.crypto_adapter.GenerateRandomBytes()
	url.ShortUrl = randomBytes

	u.url_repository.CreateShortURL(ctx, url, userId)

	ch <- UseCaseResponse{
		Success: true,
		Code:    201,
	}
}

func (u URLUseCase) GetOriginalURL(ctx context.Context, ch chan<- UseCaseResponse, shortUrl string) {
	defer RecoverPanic(ch, "GetOriginalURL")()

	url, err := u.url_repository.GetOriginalURL(ctx, shortUrl)
	if err != nil {
		ch <- UseCaseResponse{
			Success: false,
			Data:    "no matching url",
			Code:    404,
		}
		return
	}

	ch <- UseCaseResponse{
		Success: true,
		Data:    url,
		Code:    200,
	}
}
