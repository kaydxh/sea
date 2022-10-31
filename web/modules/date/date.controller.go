// Copyright 2020 The kaydxh Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package date

import (
	"context"
	"fmt"
	"time"

	"github.com/kaydxh/sea/api/openapi-spec/date"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Controller struct {

	// Embed the unimplemented server
	date.UnimplementedDateServiceServer
}

// 日期查询
func (c *Controller) Now(
	ctx context.Context,
	req *date.NowRequest,
) (resp *date.NowResponse, err error) {
	return &date.NowResponse{
		RequestId: req.GetRequestId(),
		Date:      time.Now().String(),
	}, nil
}

func (c *Controller) NowError(
	ctx context.Context,
	req *date.NowErrorRequest,
) (resp *date.NowErrorResponse, err error) {
	err = fmt.Errorf("InvalidArgument")
	return nil, status.Error(codes.InvalidArgument, err.Error())
}
