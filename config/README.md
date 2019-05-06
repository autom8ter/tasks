# config
--
    import "github.com/autom8ter/tasks/config"


## Usage

#### type Config

```go
type Config struct {
	DBConfig           *DBConfig
	UnaryInterceptors  []grpc.UnaryServerInterceptor
	StreamInterceptors []grpc.StreamServerInterceptor
	Opts               []grpc.ServerOption
}
```


#### func  NewConfig

```go
func NewConfig(opts ...Option) *Config
```

#### func (*Config) GRPCServer

```go
func (c *Config) GRPCServer(grpcFunc functions.GrpcFunc) *grpc.Server
```

#### func (*Config) PGOptions

```go
func (s *Config) PGOptions() *pg.Options
```

#### func (*Config) Validate

```go
func (c *Config) Validate() error
```

#### type DBConfig

```go
type DBConfig struct {
	Addr     string `validate:"required"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	Database string `validate:"required"`
}
```


#### type Option

```go
type Option func(c *Config)
```


#### func  WithDBConfig

```go
func WithDBConfig(cfg *DBConfig) Option
```

#### func  WithServerOptions

```go
func WithServerOptions(opts ...grpc.ServerOption) Option
```

#### func  WithStreamInterceptors

```go
func WithStreamInterceptors(interceptors ...grpc.StreamServerInterceptor) Option
```

#### func  WithUnaryInterceptors

```go
func WithUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) Option
```
