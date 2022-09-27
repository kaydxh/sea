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
