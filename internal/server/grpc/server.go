package grpc

import (
	"fmt"
	"gohub/configs"
	"gohub/database"
	"gohub/internal/libs/logger"
	"gohub/internal/libs/validation"
	middleware "gohub/pkg/middleware"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	engine    *grpc.Server
	cfg       *configs.Config
	validator validation.Validation
	db        database.IDatabase
}

func NewServer(validator validation.Validation, db database.IDatabase) *Server {
	interceptor := middleware.NewAuthInterceptor(configs.AuthIgnoreMethods)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.Unary(),
		),
	)

	return &Server{
		engine:    grpcServer,
		cfg:       configs.GetConfig(),
		validator: validator,
		db:        db,
	}
}

func (s Server) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.GrpcPort))
	logger.Info("GRPC server is listening on PORT: ", s.cfg.GrpcPort)
	if err != nil {
		logger.Error("Failed to listen: ", err)
		return err
	}

	err = s.engine.Serve(lis)
	if err != nil {
		logger.Fatal("Failed to serve grpc: ", err)
		return err
	}

	return nil
}
