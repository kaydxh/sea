package redisdao_test

import (
	"context"
	"sync"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"

	redis_ "github.com/kaydxh/golang/pkg/database/redis"
	viper_ "github.com/kaydxh/golang/pkg/viper"
	redisdao "github.com/kaydxh/sea/pkg/sealet/database/dao/redis.dao"
	"github.com/kaydxh/sea/pkg/sealet/database/model"
)

var (
	onceDB sync.Once
	db     *redis.Client
	err    error
)

func GetDBOrDie() *redis.Client {
	onceDB.Do(func() {
		cfgFile := "../../../../../conf/sealet.yaml"
		config := redis_.NewConfig(redis_.WithViper(viper_.GetViper(cfgFile, "database.redis")))

		db, err = config.Complete().New(context.Background())
		if err != nil {
			panic(err)
		}

		if db == nil {
			panic("db is not enable")
		}
	})

	return db
}

var (
	onceDao sync.Once
	taskDao *redisdao.TaskDao
)

func GetTaskDao() *redisdao.TaskDao {
	onceDao.Do(func() {
		taskDao = redisdao.NewTaskDao(GetDBOrDie())
	})

	return taskDao
}

func TestAddTask(t *testing.T) {

	err := GetTaskDao().AddTask(context.Background(), model.Task{
		TaskName: "task3",
		TaskId:   uuid.New().String(),
		TaskType: 3,
	})
	if err != nil {
		t.Fatalf("failed to add tasks, err: %v", err)
	}
}
