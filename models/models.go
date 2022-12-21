package models

import (
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
	"xorm.io/xorm/names"
)

type ColName string

const (
	ColTitle    ColName = "Title"
	ColContent  ColName = "Content"
	ColParentId ColName = "ParentId"
	ColDoneAt   ColName = "DoneAt"
)

func MapNames(m names.Mapper, cn ...ColName) []string {
	values := make([]string, len(cn))
	for i, v := range cn {
		values[i] = m.Obj2Table(string(v))
	}
	return values
}

type Todo struct {
	Id        int64
	Title     string    `xorm:"TEXT notnull"`
	Content   string    `xorm:"TEXT"`
	ParentId  int64     // FK back reference if it's created from exiting one
	CreatedAt time.Time `xorm:"updated index"`
	UpdatedAt time.Time `xorm:"created"`
	DoneAt    time.Time
}

type DB struct {
	engine *xorm.Engine
	mapper names.Mapper
}

func (db *DB) mapName(filedName string) string {
	return db.mapper.Obj2Table(filedName)
}

func (db *DB) UpdateTodoForce(todo *Todo) (*Todo, error) {
	_, err := db.engine.AllCols().Update(todo)
	return todo, err
}

func (db *DB) DeleteTodo(id int64) (*Todo, error) {
	todo := &Todo{Id: id}
	_, err := db.engine.Delete(todo)
	return todo, err
}

func (db *DB) GetAll() ([]*Todo, error) {
	var todos []*Todo
	err := db.engine.Desc(db.mapName("CreatedAt")).Find(&todos)
	return todos, err
}

func (db *DB) Get(id int64) (*Todo, error) {
	todo := &Todo{Id: id}
	found, err := db.engine.Get(todo)
	if !found {
		return nil, xorm.ErrNotExist
	}
	return todo, err
}

func (db *DB) UpdateTodo(todo *Todo, cols ...ColName) (*Todo, error) {
	if len(cols) == 0 {
		panic("requested update with no column specified")
	}

	names := MapNames(db.mapper, cols...)
	_, err := db.engine.ID(todo.Id).Cols(names...).Update(todo)
	return todo, err
}

func (db *DB) UpdateContent(id int64, content string) (*Todo, error) {
	todo := &Todo{Id: id, Content: content}
	_, err := db.engine.ID(id).Update(todo)
	return todo, err
}

func (db *DB) ToggleTodo(id int64) (*Todo, error) {
	todo := &Todo{Id: id}
	ok, err := db.engine.Get(todo)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("record not found with id: %d", id)
	}

	if todo.DoneAt.IsZero() {
		todo.DoneAt = time.Now()
	} else {
		todo.DoneAt = time.Time{}
	}

	_, err = db.engine.Cols(db.mapName("DoneAt")).Update(todo)

	return todo, err
}

func (db *DB) CreateTodo(todo *Todo) (*Todo, error) {
	_, err := db.engine.Insert(todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func NewDB() *DB {
	engine, err := xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	engine.ShowSQL(true)
	engine.Logger().SetLevel(log.LOG_DEBUG)

	engine.SetMapper(names.GonicMapper{})

	err = engine.Sync2(new(Todo))
	if err != nil {
		panic(err)
	}

	return &DB{
		engine: engine,
		mapper: engine.GetColumnMapper(),
	}
}
