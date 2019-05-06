# db
--
    import "github.com/autom8ter/tasks/db"


## Usage

```go
var DB_ERROR = func(err error) error {
	return status.Error(codes.Internal, "database error: "+err.Error())
}
```

```go
var STREAM_ERROR = func(err error) error {
	return status.Error(codes.Internal, "streaming error: "+err.Error())
}
```

```go
var UNIMPLEMENTED_ERR = status.Error(codes.Unimplemented, "api is currently unimplemented")
```

#### type Database

```go
type Database struct {
	functions.GrpcFunc
}
```


#### func  NewDatabase

```go
func NewDatabase(opts *pg.Options) *Database
```

#### func (*Database) Create

```go
func (d *Database) Create(ctx context.Context, t *tasks.Task) (*tasks.Task, error)
```

#### func (*Database) Delete

```go
func (d *Database) Delete(ctx context.Context, i *tasks.IDRequest) (*empty.Empty, error)
```

#### func (*Database) Init

```go
func (d *Database) Init() error
```

#### func (*Database) List

```go
func (d *Database) List(e *empty.Empty, stream tasks.TaskService_ListServer) error
```

#### func (*Database) Read

```go
func (d *Database) Read(ctx context.Context, i *tasks.IDRequest) (*tasks.Task, error)
```

#### func (*Database) Update

```go
func (d *Database) Update(ctx context.Context, t *tasks.Task) (*empty.Empty, error)
```

#### type Task

```go
type Task struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Priority string `json:"priority"`
	Content  string `json:"content"`
	DueDate  string `json:"due_date"`
}
```


#### func  TaskFromTask

```go
func TaskFromTask(t *tasks.Task) *Task
```

#### func (*Task) ToTask

```go
func (t *Task) ToTask() *tasks.Task
```
