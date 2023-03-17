/*
 *Copyright (c) 2022, kaydxh
 *
 *Permission is hereby granted, free of charge, to any person obtaining a copy
 *of this software and associated documentation files (the "Software"), to deal
 *in the Software without restriction, including without limitation the rights
 *to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *copies of the Software, and to permit persons to whom the Software is
 *furnished to do so, subject to the following conditions:
 *
 *The above copyright notice and this permission notice shall be included in all
 *copies or substantial portions of the Software.
 *
 *THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *SOFTWARE.
 */
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
