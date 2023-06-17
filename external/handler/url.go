package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/marcoscoutinhodev/url_shortener_api/external/adapter"
	"github.com/marcoscoutinhodev/url_shortener_api/external/middlewares"
	"github.com/marcoscoutinhodev/url_shortener_api/external/repository"
	_ "github.com/marcoscoutinhodev/url_shortener_api/external/swagger"
	"github.com/marcoscoutinhodev/url_shortener_api/internal/dto"
	"github.com/marcoscoutinhodev/url_shortener_api/internal/usecase"
)

// Create url godoc
// @Summary			Create URL
// @Description Create URL
// @Tags				url
// @Accept			json
// @Produce			json
// @Param				request	body			swagger.ShortURLInput	true	"url request"
// @Success			200			{object}	swagger.ToJSONSuccess
// @Failure			400			{object}	swagger.ToJSONError
// @Failure			500			{object}	swagger.ToJSONError
// @Router			/url		[post]
// @Security            ApiKeyAuth
func CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var shortURLInput dto.ShortURLInput
	err := json.NewDecoder(r.Body).Decode(&shortURLInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ToJson(false, "no data was provided"))
		return
	}

	urlUseCase := usecase.NewURLUseCase(
		repository.NewURLRepository(),
		adapter.NewURLCheckerAdapter(),
		adapter.NewCryptoAdapter(),
	)

	props := r.Context().Value(middlewares.AuthProps{}).(jwt.MapClaims)

	ch := make(chan usecase.UseCaseResponse)
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	go urlUseCase.CreateShortURL(ctx, ch, &shortURLInput, props["sub"].(string))

	select {
	case res := <-ch:
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(ToJson(res.Success, res.Data))
	case <-ctx.Done():
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ToJson(false, "internal server error, please try again in a few minutes"))
	}
}

// Get original url godoc
// @Summary			Get original URL
// @Description Get original URL
// @Tags				url
// @Accept			json
// @Produce			json
// @Param				short_url	path			string	true "short_url"
// @Success			200				{object}	swagger.ToJSONSuccess
// @Failure			400				{object}	swagger.ToJSONError
// @Failure			500				{object}	swagger.ToJSONError
// @Router			/url/{short_url}			[get]
// @Security            	ApiKeyAuth
func GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	shortUrl := chi.URLParam(r, "shortURL")
	if s := strings.ReplaceAll(shortUrl, " ", ""); s == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ToJson(false, "no url was provided"))
		return
	}

	urlUseCase := usecase.NewURLUseCase(
		repository.NewURLRepository(),
		adapter.NewURLCheckerAdapter(),
		adapter.NewCryptoAdapter(),
	)

	ch := make(chan usecase.UseCaseResponse)
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	go urlUseCase.GetOriginalURL(ctx, ch, shortUrl)

	select {
	case res := <-ch:
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(ToJson(res.Success, res.Data))
	case <-ctx.Done():
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ToJson(false, "internal server error, please try again in a few minutes"))
	}
}
