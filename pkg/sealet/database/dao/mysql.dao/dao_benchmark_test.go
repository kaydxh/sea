package mysqldao_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/kaydxh/sea/pkg/sealet/database/model"
)

// -count the benchmark times, -benchtime the test execute times(用例执行次数) or execute time(用例执行时间)
//go test -run=dao_benchmark_test.go  -test.bench="AddTask" -benchtime=5s -count=3 .
//go test -bench="AddTask" -benchtime=50x -count=3 .
func BenchmarkAddTask(t *testing.B) {
	for n := 0; n < t.N; n++ {
		fmt.Println("n: ", n)
		err := GetTaskDao().AddTask(context.Background(), model.Task{
			TaskName: "task3",
			TaskId:   uuid.New().String(),
			TaskType: 3,
		})
		if err != nil {
			t.Fatalf("failed to add tasks, err: %v", err)
		}
	}
}

func BenchmarkParallelAddTask(t *testing.B) {
	t.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := GetTaskDao().AddTask(context.Background(), model.Task{
				TaskName: "task3",
				TaskId:   uuid.New().String(),
				TaskType: 3,
			})
			if err != nil {
				t.Fatalf("failed to add tasks, err: %v", err)
			}
		}
	})
}
