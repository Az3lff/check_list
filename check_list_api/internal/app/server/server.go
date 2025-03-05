package server

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/Az3lff/check_list/internal/config"
	checkListTask "github.com/Az3lff/check_list/internal/delivery/grpc"
	"github.com/Az3lff/check_list/internal/service"
)

const (
	TCP = "tcp"
)

type Server struct {
	cfg        config.GRPCServer
	grpcServer *grpc.Server
	svc        service.CheckList
}

func New(cfg config.GRPCServer, service service.CheckList) *Server {
	srv := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: cfg.MaxConnectionIdle,
			Timeout:           cfg.Timeout,
			Time:              cfg.TimePingClient,
		}),
	)

	// TODO: add validate middleware

	return &Server{
		cfg:        cfg,
		grpcServer: srv,
		svc:        service,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen(TCP, s.cfg.Address)
	if err != nil {
		return err
	}

	checkListTask.Register(s.grpcServer, s.svc)

	if err := s.grpcServer.Serve(listener); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
}
