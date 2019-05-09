package functions

import (
	"github.com/autom8ter/tasks/pkg/config"
	"google.golang.org/grpc"
)

//ConfigFunc does something to a configuration
type ConfigFunc func(c *config.Config)

func NewConfig(opts ...ConfigFunc) *config.Config {
	c := &config.Config{}
	for _, o := range opts {
		o(c)
	}
	return c
}

//WithUnaryInterceptors adds the provided gRPC unary server interceptors(variadic) to a configuration
func WithUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) ConfigFunc {
	return func(c *config.Config) {
		c.UnaryInterceptors = append(c.UnaryInterceptors, interceptors...)
	}
}

//WithStreamInterceptors adds the provided gRPC stream server interceptors(variadic) to a configuration
func WithStreamInterceptors(interceptors ...grpc.StreamServerInterceptor) ConfigFunc {
	return func(c *config.Config) {
		c.StreamInterceptors = append(c.StreamInterceptors, interceptors...)
	}
}

//WithServerConfigFuncs adds the provided gRPC server ConfigFuncs(variadic) to a configuration
func WithServerConfigFuncs(opts ...grpc.ServerOption) ConfigFunc {
	return func(c *config.Config) {
		c.Opts = append(c.Opts, opts...)
	}
}

//WithDBConfig adds the provided DBConfig to a configuration
func WithDBConfig(cfg *config.DBConfig) ConfigFunc {
	return func(c *config.Config) {
		c.DBConfig = cfg
	}
}
