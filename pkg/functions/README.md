# functions
--
    import "github.com/autom8ter/tasks/pkg/functions"


## Usage

#### func  NewConfig

```go
func NewConfig(opts ...ConfigFunc) *config.Config
```

#### type AuthFunc

```go
type AuthFunc func() string
```


#### func (AuthFunc) AsNatsHandler

```go
func (a AuthFunc) AsNatsHandler() nats.AuthTokenHandler
```

#### type BasicFunc

```go
type BasicFunc func()
```


#### type ConfigFunc

```go
type ConfigFunc func(c *config.Config)
```

ConfigFunc does something to a configuration

#### func  WithDBConfig

```go
func WithDBConfig(cfg *config.DBConfig) ConfigFunc
```
WithDBConfig adds the provided DBConfig to a configuration

#### func  WithServerConfigFuncs

```go
func WithServerConfigFuncs(opts ...grpc.ServerOption) ConfigFunc
```
WithServerConfigFuncs adds the provided gRPC server ConfigFuncs(variadic) to a
configuration

#### func  WithStreamInterceptors

```go
func WithStreamInterceptors(interceptors ...grpc.StreamServerInterceptor) ConfigFunc
```
WithStreamInterceptors adds the provided gRPC stream server
interceptors(variadic) to a configuration

#### func  WithUnaryInterceptors

```go
func WithUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) ConfigFunc
```
WithUnaryInterceptors adds the provided gRPC unary server interceptors(variadic)
to a configuration

#### type GrpcFunc

```go
type GrpcFunc func(s *grpc.Server)
```

GrpcFunc does something to a gRPC server

#### func (GrpcFunc) GRPCServer

```go
func (g GrpcFunc) GRPCServer(c *config.Config) *grpc.Server
```
GrpcServer creates a grpc server from the configurations server ConfigFuncs,
interceptors, and the provided GRPCRunc

#### type RouterFunc

```go
type RouterFunc func(r *mux.Router)
```


#### func  WithDocs

```go
func WithDocs() RouterFunc
```

#### func  WithMetrics

```go
func WithMetrics() RouterFunc
```

#### func  WithProfiling

```go
func WithProfiling() RouterFunc
```

#### func  WithTaskCreate

```go
func WithTaskCreate(client tasks.TaskServiceClient) RouterFunc
```

#### func  WithTaskDelete

```go
func WithTaskDelete(client tasks.TaskServiceClient) RouterFunc
```

#### func  WithTaskList

```go
func WithTaskList(client tasks.TaskServiceClient) RouterFunc
```

#### func  WithTaskRead

```go
func WithTaskRead(client tasks.TaskServiceClient) RouterFunc
```

#### func  WithTaskUpdate

```go
func WithTaskUpdate() RouterFunc
```
