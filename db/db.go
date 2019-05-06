//go:generate godocdown -o README.md

package db

import (
	"context"
	"github.com/autom8ter/tasks/functions"
	"github.com/autom8ter/tasks/sdk/go/tasks"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var UNIMPLEMENTED_ERR = status.Error(codes.Unimplemented, "api is currently unimplemented")
var DB_ERROR = func(err error) error {
	return status.Error(codes.Internal, "database error: "+err.Error())
}
var STREAM_ERROR = func(err error) error {
	return status.Error(codes.Internal, "streaming error: "+err.Error())
}

type Database struct {
	db *pg.DB
	functions.GrpcFunc
}

func (d *Database) List(e *empty.Empty, stream tasks.TaskService_ListServer) error {
	tsks := []*Task{}
	err := d.db.Model(&Task{}).Column("*").Select(&tsks)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	for _, t := range tsks {
		err = stream.Send(t.ToTask())
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}
	return nil
}

type Task struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Priority string `json:"priority"`
	Content  string `json:"content"`
	DueDate  string `json:"due_date"`
}

func TaskFromTask(t *tasks.Task) *Task {
	return &Task{
		Id:       int(t.Id),
		Title:    t.Title,
		Category: t.Category,
		Priority: t.Priority.String(),
		Content:  t.Content,
		DueDate:  t.DueDate,
	}
}

func (t *Task) ToTask() *tasks.Task {
	var pri tasks.Priority
	switch t.Priority {
	case "high", "High", "HIGH":
		pri = tasks.Priority_High
	case "medium", "Medium", "MEDIUM":
		pri = tasks.Priority_Medium
	case "low", "Low", "LOW":
		pri = tasks.Priority_Low
	}
	return &tasks.Task{
		Id:       int64(t.Id),
		Title:    t.Title,
		Category: t.Category,
		Priority: pri,
		Content:  t.Content,
		DueDate:  t.DueDate,
	}
}

func NewDatabase(opts *pg.Options) *Database {
	db := pg.Connect(opts)
	d := &Database{
		db: db,
	}
	d.GrpcFunc = func(s *grpc.Server) {
		tasks.RegisterTaskServiceServer(s, d)
	}
	return d
}

func (d *Database) Init() error {
	for _, model := range []interface{}{(*Task)(nil)} {
		err := d.db.CreateTable(model, &orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
			Varchar:     64,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Database) Read(ctx context.Context, i *tasks.IDRequest) (*tasks.Task, error) {
	t, err := d.selectById(int(i.Id))
	if err != nil {
		return nil, DB_ERROR(err)
	}

	return t, nil
}

func (d *Database) Create(ctx context.Context, t *tasks.Task) (*tasks.Task, error) {
	err := d.insertTask(t)
	if err != nil {
		return nil, DB_ERROR(err)
	}
	return t, nil
}

func (d *Database) Update(ctx context.Context, t *tasks.Task) (*empty.Empty, error) {
	err := d.updateTask(t)
	if err != nil {
		return nil, DB_ERROR(err)
	}
	return &empty.Empty{}, nil
}

func (d *Database) Delete(ctx context.Context, i *tasks.IDRequest) (*empty.Empty, error) {
	err := d.deleteTask(i.Id)
	if err != nil {
		return nil, DB_ERROR(err)
	}
	return &empty.Empty{}, nil
}

func (d *Database) insertTask(task *tasks.Task) error {
	return d.db.Insert(TaskFromTask(task))
}

func (d *Database) updateTask(t *tasks.Task) error {
	tsk := TaskFromTask(t)
	_, err := d.db.Model(tsk).Column("title", "category", "priority", "content", "due_date").Where("id = ?", tsk.Id).Returning("*").Update()
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) deleteTask(id int64) error {
	_, err := d.db.Model(&Task{}).Where("id = ?", id).Delete(&Task{Id: int(id)})
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) selectById(id int) (*tasks.Task, error) {
	t := &Task{}
	err := d.db.Model(t).Where("id = ?", id).Select(t)
	if err != nil {
		return nil, err
	}

	return t.ToTask(), nil
}
