package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/marcoscoutinhodev/url_shortener_api/external/adapter"
	"github.com/marcoscoutinhodev/url_shortener_api/external/repository"
	_ "github.com/marcoscoutinhodev/url_shortener_api/external/swagger"
	"github.com/marcoscoutinhodev/url_shortener_api/internal/dto"
	"github.com/marcoscoutinhodev/url_shortener_api/internal/usecase"
)

// Create user godoc
// @Summary			Create User
// @Description Create User
// @Tags				users
// @Accept			json
// @Produce			json
// @Param				request				body			swagger.UserInputSignUp	true	"user request"
// @Success			201						{object}	swagger.ToJSONSuccess
// @Failure			400						{object}	swagger.ToJSONError
// @Failure			500						{object}	swagger.ToJSONError
// @Router			/user/signup	[post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userInput dto.UserInput
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ToJson(false, "no data was provided"))
		return
	}

	userUseCase := usecase.NewUserUseCase(
		repository.NewUserRepository(),
		adapter.NewHashAdapter(),
		adapter.NewEncryptAdapter(),
	)

	ch := make(chan usecase.UseCaseResponse)
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	go userUseCase.CreateUser(ctx, ch, userInput)

	select {
	case res := <-ch:
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(ToJson(res.Success, res.Data))
	case <-ctx.Done():
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ToJson(false, "internal server error, please try again in a few minutes"))
	}
}

// Authenticate user godoc
// @Summary			Authenticate User
// @Description Authenticate User
// @Tags				users
// @Accept			json
// @Produce			json
// @Param				request				body			swagger.UserInputSignIn	true	"user request"
// @Success			200						{object}	swagger.ToJSONSuccess
// @Failure			400						{object}	swagger.ToJSONError
// @Failure			500						{object}	swagger.ToJSONError
// @Router			/user/signin	[post]
func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var userInput dto.UserInput
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ToJson(false, "no data was provided"))
		return
	}

	userUseCase := usecase.NewUserUseCase(
		repository.NewUserRepository(),
		adapter.NewHashAdapter(),
		adapter.NewEncryptAdapter(),
	)

	ch := make(chan usecase.UseCaseResponse)
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	go userUseCase.AuthenticateUser(ctx, ch, userInput)

	select {
	case res := <-ch:
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(ToJson(res.Success, res.Data))
	case <-ctx.Done():
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ToJson(false, "internal server error, please try again in a few minutes"))
	}
}
