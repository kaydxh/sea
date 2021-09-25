package redisdao

import (
	"context"

	"github.com/go-redis/redis/v8"

	runtime_ "github.com/kaydxh/golang/go/runtime"
	time_ "github.com/kaydxh/golang/go/time"
	database_ "github.com/kaydxh/golang/pkg/database"
	redis_ "github.com/kaydxh/golang/pkg/database/redis"

	"github.com/kaydxh/sea/pkg/sealet/database/dao"
	"github.com/kaydxh/sea/pkg/sealet/database/model"
	"github.com/sirupsen/logrus"
)

type TaskDao struct {
	db *redis.Client
}

func NewTaskDao(db *redis.Client) *TaskDao {
	return &TaskDao{db: db}
}

// AddTask
func (d *TaskDao) AddTask(ctx context.Context, arg model.Task) error {

	tc := time_.New(true)
	caller := runtime_.GetShortCaller()
	logger := logrus.WithField("caller", caller)
	clean := func() {
		tc.Tick(caller)
		logger.WithField("cost", tc.String()).Infof("REDIS EXECL")
	}
	defer clean()

	logger.WithField("request", arg).Infof("AddTask")

	ctx, cancel := database_.WithDatabaseExecuteTimeout(ctx, dao.DatabaseExecuteTimeout)
	defer cancel()

	err := redis_.HSetStruct(ctx, d.db, arg.TaskId, arg)
	if err != nil {
		return err
	}

	logger.WithField("cost", tc.String()).Infof("successed HSet")

	return nil
}
