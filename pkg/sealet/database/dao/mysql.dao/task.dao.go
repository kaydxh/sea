package mysqldao

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	context_ "github.com/kaydxh/golang/go/context"
	runtime_ "github.com/kaydxh/golang/go/runtime"
	time_ "github.com/kaydxh/golang/go/time"
	mysql_ "github.com/kaydxh/golang/pkg/database/mysql"
	"github.com/kaydxh/sea/pkg/sealet/database/dao"
	"github.com/kaydxh/sea/pkg/sealet/database/model"
	"github.com/sirupsen/logrus"
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
func (d *TaskDao) GetTasks(ctx context.Context, conds []string, arg *model.Task) ([]*model.Task, error) {

	ctx, cancel := context_.WithTimeout(ctx, dao.DatabaseExecuteTimeout)
	defer cancel()

	query := fmt.Sprintf(
		`SELECT * FROM %s
	             WHERE %s`,
		taskTableName,
		mysql_.ConditionWithEqualAnd(conds...),
	)
	return d.getTasks(ctx, query, arg)
}

func (d *TaskDao) GetTasksByPage(
	ctx context.Context,
	offset, limit int32,
	conds []string,
	arg map[string]interface{},
) ([]*model.Task, error) {
	ctx, cancel := context_.WithTimeout(ctx, dao.DatabaseExecuteTimeout)
	defer cancel()

	query := fmt.Sprintf(
		`SELECT * FROM  %s
		     WHERE %s
			   ORDER BY create_time DESC, id DESC limit %d, %d`,
		taskTableName,
		mysql_.ConditionWithEqualAnd(conds...),
	)
	return m.getTasks(ctx, query, arg)
}

// GetTasksCount
func (d *TaskDao) GetTasksCount(ctx context.Context, conds []string, arg *model.Task) (uint64, error) {

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

func (d *TaskDao) getTasks(ctx context.Context, query string, arg interface{}) (dest []*model.Task, err error) {
	tc := time_.New(true)
	caller := runtime_.GetShortCaller()
	logger := logrus.WithField("caller", caller)
	clean := func() {
		tc.Tick(caller)
		logger.WithField("cost", tc.String()).Infof("SQL EXECL")
		if err != nil {
			logger.WithError(err).Errorf("sql: %s", query)
		}
	}
	defer clean()

	// Check that invalid preparations fail
	ns, err := d.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer ns.Close()

	err = ns.SelectContext(ctx, &dest, arg)
	if err != nil {
		return nil, err
	}
	/*
		rows, err := d.db.NamedQueryContext(ctx, query, arg)
		if err != nil {
			fmt.Printf("===err: %v\n", err)
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var task model.Task
			err = rows.StructScan(&task)
			if err != nil {
				fmt.Printf("StructScan===err: %v\n", err)
				return nil, err
			}

			dest = append(dest, &task)
		}
	*/

	return dest, nil
}
