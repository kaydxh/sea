// Copyright 2020 The kaydxh Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package seadate

import (
	"context"

	logs_ "github.com/kaydxh/golang/pkg/logs"
	v1 "github.com/kaydxh/sea/api/protoapi-spec/sea-date/v1"
	"github.com/kaydxh/sea/pkg/sea-date/application"
	"github.com/kaydxh/sea/pkg/sea-date/domain/date"
)

type Controller struct {
	app application.Application

	// Embed the unimplemented server
	v1.UnimplementedSeaDateServiceServer
}

// 日期查询
func (c *Controller) Now(
	ctx context.Context,
	req *v1.NowRequest,
) (resp *v1.NowResponse, err error) {
	logger := logs_.GetLoggerOrFallback(ctx, req.GetRequestId())
	dateReq := &date.NowRequest{
		RequestId: req.GetRequestId(),
	}
	dateResp, err := c.app.Commands.SeaDateHandler.Now(ctx, dateReq)
	if err != nil {
		logger.WithError(err).WithField("cmd", "Sealet").Errorf("failed to run [Now] command")
		return nil, APIError(err)
	}

	resp = &v1.NowResponse{
		Date: dateResp.Date,
	}

	return resp, nil
}

func (c *Controller) NowError(
	ctx context.Context,
	req *v1.NowErrorRequest,
) (resp *v1.NowErrorResponse, err error) {
	logger := logs_.GetLoggerOrFallback(ctx, req.GetRequestId())
	dateReq := &date.NowErrorRequest{
		RequestId: req.GetRequestId(),
	}
	dateResp, err := c.app.Commands.SeaDateHandler.NowError(ctx, dateReq)
	if err != nil {
		logger.WithError(err).WithField("cmd", "Sealet").Errorf("failed to run [NowError] command")
		return nil, APIError(err)
	}

	resp = &v1.NowErrorResponse{
		Date: dateResp.Date,
	}

	return resp, nil
}
