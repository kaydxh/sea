/*
 *Copyright (c) 2022, kaydxh
 *
 *Permission is hereby granted, free of charge, to any person obtaining a copy
 *of this software and associated documentation files (the "Software"), to deal
 *in the Software without restriction, including without limitation the rights
 *to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *copies of the Software, and to permit persons to whom the Software is
 *furnished to do so, subject to the following conditions:
 *
 *The above copyright notice and this permission notice shall be included in all
 *copies or substantial portions of the Software.
 *
 *THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *SOFTWARE.
 */
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
		err := GetTaskDao().AddTask(context.Background(), &model.Task{
			TaskName: "task3",
			TaskId:   uuid.New().String(),
			TaskType: 3,
		})
		if err != nil {
			t.Fatalf("failed to add tasks, err: %v", err)
		}
	}
}

//go test -v dao_benchmark_test.go  -test.bench=" BenchmarkParallelAddTask" .
func BenchmarkParallelAddTask(t *testing.B) {
	t.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := GetTaskDao().AddTask(context.Background(), &model.Task{
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
