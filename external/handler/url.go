package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/marcoscoutinhodev/url_shortener_api/external/adapter"
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

	ch := make(chan usecase.UseCaseResponse)
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	go urlUseCase.CreateShortURL(ctx, ch, &shortURLInput, "648663447fd8ef5c8687ddb3")

	select {
	case res := <-ch:
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(ToJson(res.Success, res.Data))
	case <-ctx.Done():
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ToJson(false, "internal server error, please try again in a few minutes"))
	}
}
