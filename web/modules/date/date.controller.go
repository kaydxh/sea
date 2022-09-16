// Copyright 2020 The kaydxh Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package date

import (
	"context"
	"time"

	logs_ "github.com/kaydxh/golang/pkg/logs"
	"github.com/kaydxh/sea/api/openapi-spec/date"
)

type Controller struct {

	// Embed the unimplemented server
	date.UnimplementedDateServiceServer
}

// 日期查询
func (c *Controller) Now(
	ctx context.Context,
	req *date.DateRequest,
) (resp *date.DateResponse, err error) {
	logger := logs_.GetLogger(ctx)
	logger.Infof(">>>>Now")

	return &date.DateResponse{
		RequestId: req.GetRequestId(),
		Date:      time.Now().String(),
	}, nil

	//	fmt.Printf("req: %v", req.Image[0])

	//	err = fmt.Errorf("InvalidArgument")
	//	return nil, status.Error(codes.InvalidArgument, err.Error())
}
