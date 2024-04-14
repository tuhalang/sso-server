package grpc_server

import (
	"context"
	"errors"
	"fmt"
	"github.com/tuhalang/authen/config"
	"github.com/tuhalang/authen/domain"
	domainproto "github.com/tuhalang/authen/domain/proto"
	"github.com/tuhalang/authen/internal/logger"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	serverConf        config.ServerConfig
	loginUseCase      domain.LoginUseCase
	validationUseCase domain.ValidationUseCase
	domainproto.UnimplementedLoginServer
	domainproto.UnimplementedValidateServer
}

func NewGrpcServer(serverConf config.ServerConfig, loginUseCase domain.LoginUseCase, validationUseCase domain.ValidationUseCase) Server {
	return Server{serverConf: serverConf, loginUseCase: loginUseCase, validationUseCase: validationUseCase}
}

func (s *Server) HandleLogin(ctx context.Context, req *domainproto.LoginRequest) (*domainproto.LoginResponse, error) {
	if req == nil || req.Username == "" || req.Password == "" {
		return nil, errors.New("invalid params")
	}
	loginRes, errRes := s.loginUseCase.LoginByPassword(ctx, req.Username, req.Password)
	if errRes != nil {
		return nil, errors.New(errRes.Message)
	}
	return &domainproto.LoginResponse{
		AssesToken: loginRes.AccessToken,
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *domainproto.ValidateRequest) (*domainproto.ValidateResponse, error) {
	if req == nil || req.Token == "" {
		return nil, errors.New("invalid params")
	}
	validateReq := domain.ValidationRequest{
		Path:   req.Path,
		Token:  req.Token,
		Method: req.Method,
	}
	validateRes, errRes := s.validationUseCase.Validate(ctx, validateReq)
	if errRes != nil {
		return nil, errors.New(errRes.Message)
	}
	return &domainproto.ValidateResponse{
		IsAllowed: validateRes.IsAllowed,
	}, nil
}

func (s *Server) Start() {
	log := logger.Get()
	addr := fmt.Sprintf("%s:%d", s.serverConf.Host, s.serverConf.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Panic().Err(err).Msg(err.Error())
	}

	grpcServer := grpc.NewServer()
	domainproto.RegisterLoginServer(grpcServer, s)
	domainproto.RegisterValidateServer(grpcServer, s)
	log.Info().Msgf("Grpc server listening at %s", addr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Panic().Err(err).Msg(err.Error())
	}

}
