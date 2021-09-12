package dao_test

import (
	"context"
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kaydxh/sea/pkg/sealet/database/dao"
	"github.com/kaydxh/sea/pkg/sealet/database/model"

	mysql_ "github.com/kaydxh/golang/pkg/database/mysql"
	viper_ "github.com/kaydxh/golang/pkg/viper"
)

func GetDBOrDie() *sqlx.DB {

	var (
		once sync.Once
		db   *sqlx.DB
		err  error
	)
	once.Do(func() {
		cfgFile := "../../../../conf/sealet.yaml"
		config := mysql_.NewConfig(mysql_.WithViper(viper_.GetViper(cfgFile, "database.mysql")))

		db, err = config.Complete().New()
		if err != nil {
			panic(err)
		}

		if db == nil {
			panic("db is not enable")
		}
	})

	return db
}

func TestGetTasks(t *testing.T) {

	result, err := dao.TaskDao{}.GetTasks(context.Background(), GetDBOrDie(),
		model.Task{})
	if err != nil {
		t.Fatalf("failed to get tasks, err: %v", err)
	}

	t.Logf("result of get tasks: %#v", result)
}

func TestGetTasksWithCondtion(t *testing.T) {

	result, err := dao.TaskDao{}.GetTasks(context.Background(), GetDBOrDie(),
		model.Task{
			TaskType: 1,
			TaskName: "task1",
		})
	if err != nil {
		t.Fatalf("failed to get tasks, err: %v", err)
	}

	t.Logf("result of get tasks: %#v", result)
}

func TestAddTask(t *testing.T) {

	err := dao.TaskDao{}.AddTask(context.Background(), GetDBOrDie(), model.Task{
		TaskName: "task3",
		TaskId:   uuid.New().String(),
		TaskType: 3,
	})
	if err != nil {
		t.Fatalf("failed to add tasks, err: %v", err)
	}
}

func TestDeleteTask(t *testing.T) {

	err := dao.TaskDao{}.DeleteTask(context.Background(), GetDBOrDie(), model.Task{
		TaskName: "task3",
		TaskType: 3,
	})
	if err != nil {
		t.Fatalf("failed to delete tasks, err: %v", err)
	}
}

func TestUpdateTask(t *testing.T) {

	err := dao.TaskDao{}.UpdateTask(
		context.Background(),
		GetDBOrDie(),
		[]string{"task_name"},
		[]string{"task_type"},
		model.Task{
			TaskName: "task20",
			TaskType: 2,
		},
	)
	if err != nil {
		t.Fatalf("failed to add tasks, err: %v", err)
	}
}
