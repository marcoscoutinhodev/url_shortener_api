package handler

type JSONResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func ToJson(success bool, data interface{}) *JSONResponse {
	r := &JSONResponse{
		Success: success,
	}

	if success {
		r.Data = data
	} else {
		r.Error = data
	}

	return r
}
