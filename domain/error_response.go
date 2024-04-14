package domain

type ErrorResponse struct {
	HTTPCode int    `json:"httpCode"`
	Code     string `json:"code"`
	Message  string `json:"message"`
}

const (
	InternalErr    = "000"
	AuthInvalidErr = "001"
	InvalidToken   = "002"
)

func Error(code string) *ErrorResponse {
	switch code {
	case AuthInvalidErr:
		return &ErrorResponse{HTTPCode: 400, Code: AuthInvalidErr, Message: "Username or password is invalid"}
	case InvalidToken:
		return &ErrorResponse{HTTPCode: 400, Code: InvalidToken, Message: "Token is invalid"}
	default:
		return nil
	}
}

func InternalError(msg string) *ErrorResponse {
	return &ErrorResponse{HTTPCode: 500, Code: InternalErr, Message: msg}
}
