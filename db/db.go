package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type Todo struct {
	Id        int64      `json:"id"`
	Title     string     `json:"title,omitempty"`
	Content   string     `json:"content,omitempty"`
	ParentId  int64      `json:"parent_id,omitempty"`
	Done      bool       `json:"done"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DoneAt    *time.Time `json:"done_at,omitempty"`
}

type RowScanner interface {
	Scan(dest ...any) error
}

func ScanTodo(rs RowScanner) (*Todo, error) {
	var todo Todo
	var nullContent sql.NullString
	var nullDone sql.NullBool
	var nullDoneAt sql.NullTime

	err := rs.Scan(
		&todo.Id,
		&todo.Title,
		&nullContent,
		&todo.ParentId,
		&nullDone,
		&todo.CreatedAt,
		&todo.UpdatedAt,
		&nullDoneAt,
	)
	if err != nil {
		return nil, err
	}

	if nullDone.Valid {
		todo.Done = nullDone.Bool
	}
	if nullDoneAt.Valid {
		todo.DoneAt = &nullDoneAt.Time
	}
	if nullContent.Valid {
		todo.Content = nullContent.String
	}

	return &todo, nil
}

type TodoParams struct {
	Title    string
	ParentId int64
	Content  sql.NullString
	Done     sql.NullBool
}

type Dao struct {
	db *sql.DB
}

func NewDao(dsn string) (*Dao, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(createTableStmt)
	if err != nil {
		return nil, err
	}

	return &Dao{db: db}, nil
}

func (dao *Dao) Read(id int64) (*Todo, error) {
	const readStmt = `SELECT * FROM todos WHERE id = ?`
	stmt, err := dao.db.Prepare(readStmt)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return ScanTodo(stmt.QueryRow(id))
}

func (dao *Dao) ReadAll() ([]*Todo, error) {
	const readStmt = `SELECT * FROM todos ORDER BY created_at DESC`
	rows, err := dao.db.Query(readStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*Todo

	for rows.Next() {
		todo, err := ScanTodo(rows)
		if err != nil {
			log.Println("error retriving some data", err)
			continue
		}

		todos = append(todos, todo)
	}

	err = rows.Err()
	return todos, err
}

func (dao *Dao) SetDone(id int64) (*Todo, error) {
	return dao.Update(id, TodoParams{
		Done: sql.NullBool{Valid: true, Bool: true},
	})
}

func (dao *Dao) SetUndone(id int64) (*Todo, error) {
	return dao.Update(id, TodoParams{
		Done: sql.NullBool{Valid: true, Bool: false},
	})
}

func (dao *Dao) ToggleDone(id int64) (*Todo, error) {
	todo, err := dao.Read(id)
	if err != nil {
		return nil, err
	}
	return dao.Update(id, TodoParams{
		Done: sql.NullBool{Valid: true, Bool: !todo.Done},
	})
}

func (dao *Dao) Update(id int64, param TodoParams) (*Todo, error) {
	builder := sq.StatementBuilder.Update("todos")
	if param.Title != "" {
		builder = builder.Set("title", param.Title)
	}
	if param.Content.Valid {
		builder = builder.Set("content", param.Content)
	}
	if param.ParentId != 0 {
		builder = builder.Set("parent_id", param.ParentId)
	}

	if param.Done.Valid {
		builder = builder.Set("done", param.Done)
	}

	builder = builder.Where("id = ?", id)
	builder = builder.Suffix("RETURNING *")

	return ScanTodo(builder.RunWith(dao.db).QueryRow())
}

func (dao *Dao) Delete(id int64) (*Todo, error) {
	todo, err := dao.Read(id)
	if err != nil {
		return nil, err
	}
	const deleteStmt = `DELETE FROM todos WHERE id = ?`
	_, err = dao.db.Exec(deleteStmt, id)
	if err != nil {
		return nil, err
	}
	return todo, err
}

func (dao *Dao) Create(param TodoParams) (*Todo, error) {
	const insertStmt = `
	INSERT INTO 
		todos (title, content, parent_id, done) 
		VALUES (?, ?, ?, ?) 
	RETURNING *`

	stmt, err := dao.db.Prepare(insertStmt)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return ScanTodo(stmt.QueryRow(
		param.Title,
		param.Content,
		param.ParentId,
		param.Done))
}

const createTableStmt = `
CREATE TABLE IF NOT EXISTS todos (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	content TEXT,
	parent_id INTEGER,
	done INTEGER,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	done_at DATETIME
);

CREATE TRIGGER IF NOT EXISTS UpdateField_updated_at 
UPDATE OF 
	title, content, parent_id, done, done_at
ON todos
BEGIN
  UPDATE todos SET updated_at=CURRENT_TIMESTAMP WHERE id=NEW.id;
END;

CREATE TRIGGER IF NOT EXISTS UpdateField_done_at  
UPDATE OF 
	done
ON todos
BEGIN
  UPDATE todos 
  SET done_at = CASE 
		WHEN NEW.done IS NOT NULL AND NEW.done = 1 
			THEN CURRENT_TIMESTAMP
			ELSE NULL
		END
  WHERE id=NEW.id;
END;
`
