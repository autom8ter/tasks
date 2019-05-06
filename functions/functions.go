//go:generate godocdown -o README.md

package functions

import (
	"google.golang.org/grpc"
	"net"
)

type GrpcFunc func(s *grpc.Server)

type RunFunc func() error

func GRPCRunFunc(s *grpc.Server) RunFunc {
	return func() error {
		lis, err := net.Listen("tcp", ":8080")
		if err != nil {
			return err
		}
		return s.Serve(lis)
	}

}
