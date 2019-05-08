# db
--
    import "github.com/autom8ter/tasks/db"


## Usage

```go
var DB_ERROR = func(err error) error {
	return status.Error(codes.Internal, "database error: "+err.Error())
}
```
DB_ERROR is wraps an error to indicate that it is database related

#### type Database

```go
type Database struct {
	config.GrpcFunc
}
```

Database contains a Postgres DB connection along with a GrpcFunc which registers
the Database to a gRPC server

#### func  NewDatabase

```go
func NewDatabase(opts *pg.Options) *Database
```
NewDatabase creates a new Database from postgres options and the generated
RegisterTaskServiceServer function

#### func (*Database) Create

```go
func (d *Database) Create(ctx context.Context, t *tasks.Task) (*tasks.Task, error)
```
Create inserts a Task in the database from the received Task message and then
returns it.

#### func (*Database) Delete

```go
func (d *Database) Delete(ctx context.Context, i *tasks.IDRequest) (*empty.Empty, error)
```
Delete deletes the Task with the received ID

#### func (*Database) Init

```go
func (d *Database) Init() error
```
Init initializes a Task table in the database

#### func (*Database) List

```go
func (d *Database) List(e *empty.Empty, stream tasks.TaskService_ListServer) error
```
List sends each task as a message in the gRPC stream server

#### func (*Database) Read

```go
func (d *Database) Read(ctx context.Context, i *tasks.IDRequest) (*tasks.Task, error)
```
Read returns a Task by ID

#### func (*Database) Update

```go
func (d *Database) Update(ctx context.Context, t *tasks.Task) (*empty.Empty, error)
```
Update coverts a protobuf Task into a cleaned Task, and then overwrites the
existing entry(by ID)

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

Task is just a cleaned version of the Task protobuf message for database
transactions

#### func  TaskFromTask

```go
func TaskFromTask(t *tasks.Task) *Task
```
TaskFromTask converts a protobuf Task message into a clean Task for database
transactions

#### func (*Task) ToTask

```go
func (t *Task) ToTask() *tasks.Task
```
ToTask converts the cleaned Task into a protbuf Task message
