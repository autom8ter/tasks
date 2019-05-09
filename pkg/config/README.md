# config
--
    import "github.com/autom8ter/tasks/pkg/config"


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

Config holds parameters for creating a Task server

#### func (*Config) PGOptions

```go
func (s *Config) PGOptions() *pg.Options
```
PGOptions returns postgres options used for creating a database connection

#### func (*Config) ServerOptions

```go
func (c *Config) ServerOptions() []grpc.ServerOption
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
