// Copyright 2020 The kaydxh Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package date

import (
	"context"

	logs_ "github.com/kaydxh/golang/pkg/logs"
	"github.com/kaydxh/sea/api/protoapi-spec/date"
	"github.com/kaydxh/sea/pkg/sealet/application"
	"github.com/kaydxh/sea/pkg/sealet/domain/sealet"
)

type Controller struct {
	app application.Application

	// Embed the unimplemented server
	date.UnimplementedDateServiceServer
}

// 日期查询
func (c *Controller) Now(
	ctx context.Context,
	req *date.NowRequest,
) (resp *date.NowResponse, err error) {
	logger := logs_.GetLoggerOrFallback(ctx, req.GetRequestId())
	dateReq := &sealet.NowRequest{
		RequestId: req.GetRequestId(),
	}
	dateResp, err := c.app.Commands.SealetHandler.Now(ctx, dateReq)
	if err != nil {
		logger.WithError(err).WithField("cmd", "Sealet").Errorf("failed to run [Now] command")
		return nil, err
	}

	resp = &date.NowResponse{
		Date: dateResp.Date,
	}

	return resp, nil
}

func (c *Controller) NowError(
	ctx context.Context,
	req *date.NowErrorRequest,
) (resp *date.NowErrorResponse, err error) {
	logger := logs_.GetLoggerOrFallback(ctx, req.GetRequestId())
	dateReq := &sealet.NowErrorRequest{
		RequestId: req.GetRequestId(),
	}
	dateResp, err := c.app.Commands.SealetHandler.NowError(ctx, dateReq)
	if err != nil {
		logger.WithError(err).WithField("cmd", "Sealet").Errorf("failed to run [NowError] command")
		return nil, err
	}

	resp = &date.NowErrorResponse{
		Date: dateResp.Date,
	}

	return resp, nil
}
