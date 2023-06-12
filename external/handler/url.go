package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/marcoscoutinhodev/url_shortener_api/external/adapter"
	"github.com/marcoscoutinhodev/url_shortener_api/external/middlewares"
	"github.com/marcoscoutinhodev/url_shortener_api/external/repository"
	"github.com/marcoscoutinhodev/url_shortener_api/internal/dto"
	"github.com/marcoscoutinhodev/url_shortener_api/internal/usecase"
)

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
