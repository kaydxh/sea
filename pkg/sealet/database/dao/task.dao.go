package dao

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	mysql_ "github.com/kaydxh/golang/pkg/database/mysql"
	"github.com/kaydxh/sea/pkg/sealet/database/model"
)

type TaskDao struct{}

// AddTask
func (dao TaskDao) AddTask(ctx context.Context, db *sqlx.DB, arg model.Task) error {

	query := `INSERT INTO task 
			 (is_deleted, task_name, task_id, task_type, task_status)
			 VALUES (:is_deleted,
			         :task_name,
					 :task_id,
					 :task_type,
					 :task_status)
					 ON DUPLICATE KEY UPDATE
					 is_deleted   = :is_deleted,
                     task_name    = :task_name,
                     task_id      = :task_id,
                     task_type    = :task_type,
                     task_status  = :task_status
					 `
	return arg.ExecTaskByQuery(ctx, db, query)
}

// DeleteTask
func (dao TaskDao) DeleteTask(ctx context.Context, db *sqlx.DB, arg model.Task) error {
	query := `DELETE FROM task`
	return arg.ExecTaskByQuery(
		ctx,
		db,
		mysql_.GenerateCondition(mysql_.SqlCompareEqual, mysql_.SqlOperatorAnd, query, arg),
	)
}

// UpdateTask
// UPDATE task SET foo=:foo, bar=:bar WHERE thud=:thud AND grunt=:grunt
func (dao TaskDao) UpdateTask(ctx context.Context, db *sqlx.DB, cols, conds []string, arg model.Task) error {

	query := fmt.Sprintf(
		`UPDATE task SET %s %s`,
		mysql_.JoinNamedColumnsValues(cols...),
		mysql_.GenerateNameColumsCondition(mysql_.SqlCompareEqual, mysql_.SqlOperatorAnd, conds...),
	)

	return arg.ExecTaskByQuery(ctx, db, query)
}

// GetTasks
func (dao TaskDao) GetTasks(ctx context.Context, db *sqlx.DB, arg model.Task) ([]model.Task, error) {

	query := `SELECT * FROM task`
	return arg.GetTasksByQuery(
		ctx,
		db,
		mysql_.GenerateCondition(mysql_.SqlCompareEqual, mysql_.SqlOperatorAnd, query, arg),
	)

}
