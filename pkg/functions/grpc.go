//go:generate godocdown -o README.md

package functions

import (
	"github.com/autom8ter/tasks/pkg/config"
	"google.golang.org/grpc"
)

//GrpcFunc does something to a gRPC server
type GrpcFunc func(s *grpc.Server)

//GrpcServer creates a grpc server from the configurations server ConfigFuncs, interceptors, and the provided GRPCRunc
func (g GrpcFunc) GRPCServer(c *config.Config) *grpc.Server {
	s := grpc.NewServer(c.ServerOptions()...)
	g(s)
	return s
}
