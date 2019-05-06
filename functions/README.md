# functions
--
    import "github.com/autom8ter/tasks/functions"


## Usage

#### type GrpcFunc

```go
type GrpcFunc func(s *grpc.Server)
```


#### type RunFunc

```go
type RunFunc func() error
```


#### func  GRPCRunFunc

```go
func GRPCRunFunc(s *grpc.Server) RunFunc
```
