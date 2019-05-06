package config

import (
	"google.golang.org/grpc"
)

func WithUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) Option {
	return func(c *Config) {
		c.UnaryInterceptors = append(c.UnaryInterceptors, interceptors...)
	}
}

func WithStreamInterceptors(interceptors ...grpc.StreamServerInterceptor) Option {
	return func(c *Config) {
		c.StreamInterceptors = append(c.StreamInterceptors, interceptors...)
	}
}

func WithServerOptions(opts ...grpc.ServerOption) Option {
	return func(c *Config) {
		c.Opts = append(c.Opts, opts...)
	}
}

func WithDBConfig(cfg *DBConfig) Option {
	return func(c *Config) {
		c.DBConfig = cfg
	}
}
