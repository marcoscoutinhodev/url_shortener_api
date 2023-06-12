package usecase

import "fmt"

type UseCaseResponse struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

func RecoverPanic(ch chan<- UseCaseResponse, origin string) func() {
	return func() {
		if err := recover(); err != nil {
			fmt.Printf("panic origin: %s | error: %v\n", origin, err)
			ch <- UseCaseResponse{
				Code:    500,
				Success: false,
				Data:    "internal server error, please try again in a few minutes...",
			}
		}
	}
}
