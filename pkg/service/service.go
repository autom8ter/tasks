package service

import (
	"fmt"
	"github.com/autom8ter/tasks/pkg/db"
	"github.com/autom8ter/tasks/pkg/functions"
	"google.golang.org/grpc"
	"net"
)

//Service holds everything needed to create a gRPC task service server
type Service struct {
	database   *db.Database
	grpcServer *grpc.Server
}

//NewService creates a new Service from config options
func NewService(fn ...functions.ConfigFunc) (*Service, error) {
	c := functions.NewConfig(fn...)
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
		grpcServer: d.GRPCServer(c),
	}, nil
}

//Serve starts the gRPC server on the specified address
func (s *Service) Serve(addr string) error {
	fmt.Println("starting grpc server on http://localhost:8080")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return s.grpcServer.Serve(lis)
}
