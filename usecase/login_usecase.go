package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/tuhalang/authen/config"
	"github.com/tuhalang/authen/domain"
	"github.com/tuhalang/authen/internal/tenantutil"
	"github.com/tuhalang/authen/internal/tokenutil"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type loginUseCase struct {
	userRepository    domain.UserRepository
	sessionRepository domain.SessionRepository
	keyStores         map[string]config.KeyStore
}

// NewLoginUseCase returns a LoginUseCase
func NewLoginUseCase(repository domain.UserRepository, sessionRepository domain.SessionRepository, keyStores map[string]config.KeyStore) domain.LoginUseCase {
	return &loginUseCase{
		userRepository:    repository,
		sessionRepository: sessionRepository,
		keyStores:         keyStores,
	}
}

func (lu *loginUseCase) getUserByUsername(ctx context.Context, tenant, username string) (*domain.User, error) {
	return lu.userRepository.GetByUsername(ctx, tenant, username)
}

// LoginByPassword handles login by username and password
func (lu *loginUseCase) LoginByPassword(ctx context.Context, username, password string) (*domain.LoginResponse, *domain.ErrorResponse) {
	log := zerolog.Ctx(ctx)

	tenant, username := tenantutil.GetTenant(username)
	user, err := lu.getUserByUsername(ctx, tenant, username)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return nil, domain.InternalError(err.Error())
	}

	if user == nil {
		log.Warn().Msgf("User not found: %s", username)
		return nil, domain.Error(domain.AuthInvalidErr)
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		log.Warn().Msgf("User %s invalid password!", username)
		return nil, domain.Error(domain.AuthInvalidErr)
	}

	keyStore := lu.keyStores[tenant]
	sessionID := uuid.New().String()
	accessToken, err := lu.createAccessToken(user, tenant, sessionID, keyStore.JwtKey, keyStore.JwtExp)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return nil, domain.Error(domain.AuthInvalidErr)
	}

	go lu.saveLoginInfo(ctx, tenant, domain.Session{
		SessionID: sessionID,
		UserID:    user.ID,
		LoginTime: time.Now(),
		Status:    1,
	})

	return &domain.LoginResponse{AccessToken: accessToken}, nil
}

func (lu *loginUseCase) createAccessToken(user *domain.User, issuer, sessionID, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, issuer, sessionID, secret, expiry)
}

func (lu *loginUseCase) saveLoginInfo(ctx context.Context, tenant string, session domain.Session) {
	log := zerolog.Ctx(ctx)
	err := lu.sessionRepository.Save(tenant, session)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
	}
}
