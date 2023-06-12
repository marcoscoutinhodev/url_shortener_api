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
}

func NewURLUseCase(urlRepository URLRepositoryInterface, urlCheckerAdapter URLCheckerAdapterInterface) *URLUseCase {
	return &URLUseCase{
		url_repository:      urlRepository,
		url_checker_adapter: urlCheckerAdapter,
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

	url.ShortUrl = urlEnconded
	u.url_repository.CreateShortURL(ctx, url, userId)

	ch <- UseCaseResponse{
		Success: true,
		Code:    201,
	}
}
