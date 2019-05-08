//go:generate godocdown -o README.md

package config

import (
	"github.com/go-pg/pg"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"gopkg.in/go-playground/validator.v9"
)

var valid = validator.New()

type DBConfig struct {
	Addr     string `validate:"required"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	Database string `validate:"required"`
}

func (c *Config) Validate() error {
	return valid.Struct(c)
}

//Config holds parameters for creating a Task server
type Config struct {
	DBConfig           *DBConfig
	UnaryInterceptors  []grpc.UnaryServerInterceptor
	StreamInterceptors []grpc.StreamServerInterceptor
	Opts               []grpc.ServerOption
}

func (c *Config) serverOptions() []grpc.ServerOption {
	return append(
		[]grpc.ServerOption{
			grpc_middleware.WithUnaryServerChain(c.UnaryInterceptors...),
			grpc_middleware.WithStreamServerChain(c.StreamInterceptors...),
		},

		c.Opts...,
	)
}

//GrpcFunc does something to a gRPC server
type GrpcFunc func(s *grpc.Server)

//GrpcServer creates a grpc server from the configurations server options, interceptors, and the provided GRPCRunc
func (c *Config) GRPCServer(grpcFunc GrpcFunc) *grpc.Server {
	s := grpc.NewServer(c.serverOptions()...)
	grpcFunc(s)
	return s
}

//PGOptions returns postgres options used for creating a database connection
func (s *Config) PGOptions() *pg.Options {
	return &pg.Options{
		Addr:     s.DBConfig.Addr,
		User:     s.DBConfig.User,
		Password: s.DBConfig.Password,
		Database: s.DBConfig.Database,
	}
}

type Option func(c *Config)

func NewConfig(opts ...Option) *Config {
	c := &Config{}
	for _, o := range opts {
		o(c)
	}
	return c
}
