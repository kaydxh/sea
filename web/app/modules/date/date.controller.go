// Copyright 2020 The kaydxh Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package date

import (
	"context"
	"fmt"
	"time"

	"github.com/kaydxh/sea/api/openapi-spec/v1/date"
)

type Controller struct {

	// Embed the unimplemented server
	date.UnimplementedDateServiceServer
}

// 日期查询
func (c *Controller) Now(
	_ context.Context,
	req *date.DateRequest,
) (resp *date.DateResponse, err error) {
	fmt.Println(">>>>>NOW")
	return &date.DateResponse{
		RequestId: req.GetRequestId(),
		Date:      time.Now().String(),
	}, nil
}
