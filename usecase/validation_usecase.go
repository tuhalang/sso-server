package usecase

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/tuhalang/authen/config"
	"github.com/tuhalang/authen/domain"
	"github.com/tuhalang/authen/internal/tokenutil"
)

type validationUseCase struct {
	keyStores map[string]config.KeyStore
}

// NewValidationUseCase returns a ValidateUseCase object
func NewValidationUseCase(keyStores map[string]config.KeyStore) domain.ValidationUseCase {
	return validationUseCase{keyStores: keyStores}
}

func (vu validationUseCase) Validate(ctx context.Context, req domain.ValidationRequest) (*domain.ValidationResponse, *domain.ErrorResponse) {
	log := zerolog.Ctx(ctx)
	claims, err := tokenutil.ExtractWithoutValidate(req.Token)
	if err != nil {
		log.Error().Err(err).Msg("error extracting token")
		return nil, domain.Error(domain.InvalidToken)
	}

	iss := claims["iss"].(string)
	isAuthorized, err := tokenutil.IsAuthorized(req.Token, vu.keyStores[iss].JwtKey)
	if err != nil {
		log.Error().Err(err).Msg("error verifying token")
		return nil, domain.Error(domain.InvalidToken)
	}

	return &domain.ValidationResponse{IsAllowed: isAuthorized}, nil
}
