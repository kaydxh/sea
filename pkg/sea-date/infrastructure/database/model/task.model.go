package model

import (
	"database/sql"
	"fmt"
	"strings"
)

const (
	ColTaskTaskId = "task_id"
)

type Task struct {
	Id sql.NullInt64 `db:"id"` // primary key ID

	// NullTime represents a time.Time that may be null.
	// NullTime implements the Scanner interface so
	// it can be used as a scan destination, similar to NullString.
	CreateTime sql.NullTime `db:"create_time"`
	UpdateTime sql.NullTime `db:"update_time"`

	IsDeleted  bool         `db:"is_deleted"  redis:"is_deleted"` // soft delete, 0 for not deleted, 1 for deleted
	DeleteTime sql.NullTime `db:"delete_time"`
	Version    int          `db:"version"     redis:"version"`

	TaskName   string `db:"task_name"   redis:"task_name"`
	TaskId     string `db:"task_id"     redis:"task_id"`
	TaskType   int    `db:"task_type"   redis:"task_type"`
	TaskStatus int    `db:"task_status" redis:"task_status"`
}

type Tasks []*Task

func (t Tasks) String() string {
	s := "["
	for _, task := range t {
		s += fmt.Sprintf("%v,", task)
	}
	if len(t) > 0 {
		s = strings.TrimRight(s, ",")
	}

	return s + "]"
}
