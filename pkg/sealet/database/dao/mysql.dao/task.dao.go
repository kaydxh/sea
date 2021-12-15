package mysqldao

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	context_ "github.com/kaydxh/golang/go/context"
	mysql_ "github.com/kaydxh/golang/pkg/database/mysql"
	"github.com/kaydxh/sea/pkg/sealet/database/dao"
	"github.com/kaydxh/sea/pkg/sealet/database/model"
)

const taskTableName = "task"

type TaskDao struct {
	db *sqlx.DB
}

func NewTaskDao(db *sqlx.DB) *TaskDao {
	return &TaskDao{db: db}
}

// AddTask
func (d *TaskDao) AddTask(ctx context.Context, arg *model.Task) error {

	ctx, cancel := context_.WithTimeout(ctx, dao.DatabaseExecuteTimeout)
	defer cancel()

	query := fmt.Sprintf(`INSERT INTO %s
			   (
			   is_deleted,
			   task_name,
			   task_id,
			   task_type,
			   task_status
			   )
			   VALUES (
			         :is_deleted,
			         :task_name,
					 :task_id,
					 :task_type,
					 :task_status
					 )
					 ON DUPLICATE KEY UPDATE
					 is_deleted   = :is_deleted,
                     task_name    = :task_name,
                     task_id      = :task_id,
                     task_type    = :task_type,
                     task_status  = :task_status`, taskTableName)

	return mysql_.ExecContext(ctx, query, arg, nil, d.db)
}

// DeleteTask
func (d *TaskDao) DeleteTask(ctx context.Context, conds []string, arg *model.Task) error {

	ctx, cancel := context_.WithTimeout(ctx, dao.DatabaseExecuteTimeout)
	defer cancel()

	query := fmt.Sprintf(
		`DELETE FROM %s
	           WHERE %s`,
		taskTableName,
		mysql_.ConditionWithEqualAnd(conds...))

	return mysql_.ExecContext(ctx, query, arg, nil, d.db)
}

// UpdateTask
// UPDATE task SET foo=:foo, bar=:bar WHERE thud=:thud AND grunt=:grunt
func (d *TaskDao) UpdateTask(ctx context.Context, cols, conds []string, arg *model.Task) error {

	ctx, cancel := context_.WithTimeout(ctx, dao.DatabaseExecuteTimeout)
	defer cancel()

	query := fmt.Sprintf(
		`UPDATE %s
		    SET %s
	      WHERE %s`,
		taskTableName,
		mysql_.JoinNamedColumnsValues(cols...),
		mysql_.ConditionWithEqualAnd(conds...),
	)
	return mysql_.ExecContext(ctx, query, arg, nil, d.db)
}

// GetTasks
func (d *TaskDao) GetTasks(ctx context.Context, conds []string, arg *model.Task) (tasks []*model.Task, err error) {

	ctx, cancel := context_.WithTimeout(ctx, dao.DatabaseExecuteTimeout)
	defer cancel()

	query := fmt.Sprintf(
		`SELECT * FROM %s
	             WHERE %s`,
		taskTableName,
		mysql_.ConditionWithEqualAnd(conds...),
	)

	err = mysql_.SelectNamedContext(ctx, query, arg, &tasks, d.db)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (d *TaskDao) GetTasksByPage(
	ctx context.Context,
	offset, limit int32,
	conds []string,
	arg map[string]interface{},
) (tasks []*model.Task, err error) {
	ctx, cancel := context_.WithTimeout(ctx, dao.DatabaseExecuteTimeout)
	defer cancel()

	query := fmt.Sprintf(
		`SELECT * FROM  %s
		          WHERE %s
			      ORDER BY create_time DESC, id DESC limit %d, %d`,
		taskTableName,
		mysql_.ConditionWithEqualAnd(conds...),
		offset,
		limit,
	)

	err = mysql_.SelectNamedContext(ctx, query, arg, &tasks, d.db)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// GetTasksCount
func (d *TaskDao) GetTasksCount(ctx context.Context, conds []string, arg *model.Task) (uint32, error) {

	ctx, cancel := context_.WithTimeout(ctx, dao.DatabaseExecuteTimeout)
	defer cancel()

	query := fmt.Sprintf(
		`SELECT count(*) FROM %s
		                WHERE %s`,
		taskTableName,
		mysql_.ConditionWithEqualAnd(conds...),
	)

	return mysql_.GetCountContext(ctx, query, arg, d.db)
}
