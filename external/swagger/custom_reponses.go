package swagger

type ToJSONSuccess struct {
	Success bool        `json:"success" `
	Data    interface{} `json:"data" `
}

type ToJSONError struct {
	Success bool        `json:"success" default:"false"`
	Error   interface{} `json:"error"`
}

type UserInputSignUp struct {
	Name     string `json:"name" `
	Email    string `json:"email" `
	Password string `json:"password" `
}

type UserInputSignIn struct {
	Email    string `json:"email" `
	Password string `json:"password" `
}

type ShortURLInput struct {
	OriginalURL string `json:"original_url"`
}
