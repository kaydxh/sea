package seadate

import (
	"context"

	middleware_ "github.com/kaydxh/golang/pkg/middleware"
	v1 "github.com/kaydxh/sea/api/protoapi-spec/sea-date/v1"
)

type LocalController struct {
	*Controller
}

// 日期查询
func (c *LocalController) Now(
	ctx context.Context,
	req *v1.NowRequest,
) (resp *v1.NowResponse, err error) {

	return middleware_.LocalMiddlewareWrap(c.Controller.Now)(ctx, req)
}

func (c *LocalController) NowError(
	ctx context.Context,
	req *v1.NowErrorRequest,
) (resp *v1.NowErrorResponse, err error) {
	return middleware_.LocalMiddlewareWrap(c.Controller.NowError)(ctx, req)
}
