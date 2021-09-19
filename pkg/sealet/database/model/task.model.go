package model

import (
	"database/sql"
)

type Task struct {
	Id sql.NullInt64 `db:"id"` // primary key ID

	// NullTime represents a time.Time that may be null.
	// NullTime implements the Scanner interface so
	// it can be used as a scan destination, similar to NullString.
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`

	IsDeleted bool         `db:"is_deleted"` // soft delete, 0 for not deleted, 1 for deleted
	DeletedAt sql.NullTime `db:"deleted_at"`
	Version   int          `db:"version"`

	TaskName   string `db:"task_name"`
	TaskId     string `db:"task_id"`
	TaskType   int    `db:"task_type"`
	TaskStatus int    `db:"task_status"`
}

/*
func (arg Task) GetTasksByQuery(ctx context.Context, db *sqlx.DB, query string) ([]Task, error) {
	tc := time_.New(true)
	caller := runtime_.GetShortCaller()
	logger := logrus.WithField("caller", caller)
	clean := func() {
		tc.Tick(caller)
		logger.WithField("cost", tc.String()).Infof("SQL EXECL")
	}
	defer clean()

	// Check that invalid preparations fail
	ns, err := db.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer ns.Close()

	var dest []Task
	err = ns.SelectContext(ctx, &dest, arg)
	if err != nil {
		return nil, err
	}
	return dest, nil
}

// exec sql for insert/update/delete
func (arg Task) ExecTaskByQuery(ctx context.Context, db *sqlx.DB, query string) error {
	tc := time_.New(true)
	caller := runtime_.GetShortCaller()
	logger := logrus.WithField("caller", caller)
	clean := func() {
		tc.Tick(caller)
		logger.WithField("cost", tc.String()).Infof("SQL EXECL")
	}
	defer clean()

	_, err := db.NamedExecContext(ctx, query, arg)
	if err != nil {
		return err
	}
	return nil
}
*/
