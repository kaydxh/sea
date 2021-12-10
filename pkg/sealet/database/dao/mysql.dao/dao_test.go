package mysqldao_test

import (
	"context"
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	mysqldao "github.com/kaydxh/sea/pkg/sealet/database/dao/mysql.dao"
	"github.com/kaydxh/sea/pkg/sealet/database/model"

	mysql_ "github.com/kaydxh/golang/pkg/database/mysql"
	viper_ "github.com/kaydxh/golang/pkg/viper"
)

var (
	onceDB sync.Once
	db     *sqlx.DB
	err    error
)

func GetDBOrDie() *sqlx.DB {
	onceDB.Do(func() {
		cfgFile := "../../../../../conf/sealet.yaml"
		config := mysql_.NewConfig(mysql_.WithViper(viper_.GetViper(cfgFile, "database.mysql")))

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
	taskDao *mysqldao.TaskDao
)

func GetTaskDao() *mysqldao.TaskDao {
	onceDao.Do(func() {
		taskDao = mysqldao.NewTaskDao(GetDBOrDie())
	})

	return taskDao
}

func TestGetTasks(t *testing.T) {

	results, err := GetTaskDao().GetTasks(context.Background(), []string{}, &model.Task{})
	if err != nil {
		t.Fatalf("failed to get tasks, err: %v", err)
	}

	t.Logf("result of get tasks: %v", model.Tasks(results))
}

func TestGetTasksByPage(t *testing.T) {
	const (
		offset int32 = 0
		limit  int32 = 10
	)

	filters := map[string]interface{}{
		model.ColTaskTaskId: "task1",
	}

	conds := []string{model.ColTaskTaskId}
	results, err := GetTaskDao().GetTasksByPage(context.Background(), offset, limit, conds, filters)
	if err != nil {
		t.Fatalf("failed to get tasks, err: %v", err)
	}

	t.Logf("result of get tasks: %v", model.Tasks(results))
}

func TestAddTask(t *testing.T) {

	err := GetTaskDao().AddTask(context.Background(), &model.Task{
		TaskName: "task3",
		TaskId:   uuid.New().String(),
		TaskType: 3,
	})
	if err != nil {
		t.Fatalf("failed to add tasks, err: %v", err)
	}
}

func TestDeleteTask(t *testing.T) {

	err := GetTaskDao().DeleteTask(context.Background(), []string{model.ColTaskTaskId}, &model.Task{
		TaskId: "task3",
	})
	if err != nil {
		t.Fatalf("failed to delete tasks, err: %v", err)
	}
}

func TestUpdateTask(t *testing.T) {

	err := GetTaskDao().UpdateTask(
		context.Background(),
		[]string{"task_name"},
		[]string{"task_type"},
		&model.Task{
			TaskName: "task20",
			TaskType: 2,
		},
	)
	if err != nil {
		t.Fatalf("failed to add tasks, err: %v", err)
	}
}
