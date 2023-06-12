package middlewares

type response struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
