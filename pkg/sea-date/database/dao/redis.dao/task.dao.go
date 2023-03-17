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
package redisdao

import (
	"context"

	"github.com/go-redis/redis/v8"

	runtime_ "github.com/kaydxh/golang/go/runtime"
	time_ "github.com/kaydxh/golang/go/time"
	database_ "github.com/kaydxh/golang/pkg/database"
	redis_ "github.com/kaydxh/golang/pkg/database/redis"

	"github.com/kaydxh/sea/pkg/sea-date/database/dao"
	"github.com/kaydxh/sea/pkg/sea-date/database/model"
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
