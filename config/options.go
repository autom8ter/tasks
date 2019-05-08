package config

import (
	"google.golang.org/grpc"
)

//WithUnaryInterceptors adds the provided gRPC unary server interceptors(variadic) to a configuration
func WithUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) Option {
	return func(c *Config) {
		c.UnaryInterceptors = append(c.UnaryInterceptors, interceptors...)
	}
}

//WithStreamInterceptors adds the provided gRPC stream server interceptors(variadic) to a configuration
func WithStreamInterceptors(interceptors ...grpc.StreamServerInterceptor) Option {
	return func(c *Config) {
		c.StreamInterceptors = append(c.StreamInterceptors, interceptors...)
	}
}

//WithServerOptions adds the provided gRPC server options(variadic) to a configuration
func WithServerOptions(opts ...grpc.ServerOption) Option {
	return func(c *Config) {
		c.Opts = append(c.Opts, opts...)
	}
}

//WithDBConfig adds the provided DBConfig to a configuration
func WithDBConfig(cfg *DBConfig) Option {
	return func(c *Config) {
		c.DBConfig = cfg
	}
}
