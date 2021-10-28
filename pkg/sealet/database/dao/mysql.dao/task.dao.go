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

type TaskDao struct {
	db *sqlx.DB
}

func NewTaskDao(db *sqlx.DB) *TaskDao {
	return &TaskDao{db: db}
}

// AddTask
func (d *TaskDao) AddTask(ctx context.Context, arg model.Task) error {

	ctx, cancel := context_.WithTimeout(ctx, dao.DatabaseExecuteTimeout)
	defer cancel()

	query := `INSERT INTO task 
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
                     task_status  = :task_status`

	return d.execTaskByQuery(ctx, query, arg)
}

// DeleteTask
func (d *TaskDao) DeleteTask(ctx context.Context, arg model.Task) error {

	ctx, cancel := context_.WithTimeout(ctx, dao.DatabaseExecuteTimeout)
	defer cancel()

	query := `DELETE FROM task`
	return d.execTaskByQuery(
		ctx,
		mysql_.GenerateCondition(mysql_.SqlCompareEqual, mysql_.SqlOperatorAnd, query, arg),
		arg,
	)
}

// UpdateTask
// UPDATE task SET foo=:foo, bar=:bar WHERE thud=:thud AND grunt=:grunt
func (d *TaskDao) UpdateTask(ctx context.Context, cols, conds []string, arg model.Task) error {

	ctx, cancel := context_.WithTimeout(ctx, dao.DatabaseExecuteTimeout)
	defer cancel()

	query := fmt.Sprintf(
		`UPDATE task SET %s %s`,
		mysql_.JoinNamedColumnsValues(cols...),
		mysql_.GenerateNameColumsCondition(mysql_.SqlCompareEqual, mysql_.SqlOperatorAnd, conds...),
	)
	return d.execTaskByQuery(ctx, query, arg)
}

// GetTasks
func (d *TaskDao) GetTasks(ctx context.Context, arg model.Task) ([]model.Task, error) {

	ctx, cancel := context_.WithTimeout(ctx, dao.DatabaseExecuteTimeout)
	defer cancel()

	query := `SELECT * FROM task`
	return d.getTasksByQuery(
		ctx,
		mysql_.GenerateCondition(mysql_.SqlCompareEqual, mysql_.SqlOperatorAnd, query, arg),
		arg,
	)

}

// GetTaskCount
func (d *TaskDao) GetTaskCount(ctx context.Context, conds []string, arg model.Task) (uint64, error) {

	ctx, cancel := context_.WithTimeout(ctx, dao.DatabaseExecuteTimeout)
	defer cancel()

	query := fmt.Sprintf(
		`SELECT count(*) FROM ti_model_repository %s`,
		mysql_.GenerateNameColumsCondition(mysql_.SqlCompareEqual, mysql_.SqlOperatorAnd, conds...),
	)

	return d.getCountByQuery(ctx, query, arg)
}

func (d *TaskDao) getTasksByQuery(ctx context.Context, query string, arg model.Task) ([]model.Task, error) {
	tc := time_.New(true)
	caller := runtime_.GetShortCaller()
	logger := logrus.WithField("caller", caller)
	clean := func() {
		tc.Tick(caller)
		logger.WithField("cost", tc.String()).Infof("SQL EXECL")
	}
	defer clean()

	// Check that invalid preparations fail
	ns, err := d.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer ns.Close()

	var dest []model.Task
	err = ns.SelectContext(ctx, &dest, arg)
	if err != nil {
		return nil, err
	}
	return dest, nil
}

func (d *TaskDao) getCountByQuery(ctx context.Context, query string, arg model.Task) (uint64, error) {
	tc := time_.New(true)
	caller := runtime_.GetShortCaller()
	logger := logrus.WithField("caller", caller)
	clean := func() {
		tc.Tick(caller)
		logger.WithField("cost", tc.String()).Infof("SQL EXECL")
	}
	defer clean()

	ns, err := d.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return 0, err
	}

	var count uint64
	err = ns.QueryRowContext(ctx, arg).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// exec sql for insert/update/delete
func (d *TaskDao) execTaskByQuery(ctx context.Context, query string, arg model.Task) error {
	tc := time_.New(true)
	caller := runtime_.GetShortCaller()
	logger := logrus.WithField("caller", caller)
	clean := func() {
		tc.Tick(caller)
		logger.WithField("cost", tc.String()).Infof("SQL EXECL")
	}
	defer clean()

	_, err := d.db.NamedExecContext(ctx, query, arg)
	if err != nil {
		return err
	}
	return nil
}
