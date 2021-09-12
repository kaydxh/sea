package dao_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/kaydxh/sea/pkg/sealet/database/dao"
	"github.com/kaydxh/sea/pkg/sealet/database/model"
)

func BenchmarkAddTask(t *testing.B) {

	err := dao.TaskDao{}.AddTask(context.Background(), GetDBOrDie(), model.Task{
		TaskName: "task3",
		TaskId:   uuid.New().String(),
		TaskType: 3,
	})
	if err != nil {
		t.Fatalf("failed to add tasks, err: %v", err)
	}
}
