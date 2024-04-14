package domain

import "context"

type ValidationRequest struct {
	Path   string `json:"path"`
	Token  string `json:"token"`
	Method string `json:"method"`
}

type ValidationResponse struct {
	IsAllowed bool `json:"isAllowed"`
}

type ValidationUseCase interface {
	Validate(context context.Context, req ValidationRequest) (*ValidationResponse, *ErrorResponse)
}
