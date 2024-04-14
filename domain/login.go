package domain

import "context"

type LoginRequest struct {
	Username string `json:"username" binding:"required,username"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}

type LoginUseCase interface {
	LoginByPassword(ctx context.Context, username, password string) (*LoginResponse, *ErrorResponse)
}
