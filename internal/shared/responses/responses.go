package responses

type ResponseData struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type ResponseError struct {
	Error []ResponseData `json:"error"`
}

func Error(errorMessage []string) ResponseError {
	return ResponseError{
		Error: []ResponseData{constructResponse(nil, errorMessage[0], errorMessage[1])},
	}
}

func Success(data any) ResponseData {
	return ResponseData{
		Code:    "00",
		Message: "SUCCESS",
		Data:    data,
	}
}

func constructResponse(data any, code, message string) ResponseData {
	return ResponseData{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
