package service

import (
	"fmt"
	"github.com/autom8ter/tasks/config"
	"github.com/autom8ter/tasks/db"
	"google.golang.org/grpc"
	"net"
)

type Service struct {
	database   *db.Database
	grpcServer *grpc.Server
}

func NewService(opts ...config.Option) (*Service, error) {
	c := config.NewConfig(opts...)
	if err := c.Validate(); err != nil {
		return nil, err
	}
	d := db.NewDatabase(c.PGOptions())
	err := d.Init()
	if err != nil {
		return nil, err
	}
	return &Service{
		database:   d,
		grpcServer: c.GRPCServer(d.GrpcFunc),
	}, nil
}

func (s *Service) Serve(addr string) error {
	fmt.Println("starting grpc server on http://localhost:8080")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return s.grpcServer.Serve(lis)
}
