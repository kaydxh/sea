package date

import (
	"context"

	middleware_ "github.com/kaydxh/golang/pkg/middleware"
	"github.com/kaydxh/sea/api/protoapi-spec/date"
)

type LocalController struct {
	*Controller
}

// 日期查询
func (c *LocalController) Now(
	ctx context.Context,
	req *date.NowRequest,
) (resp *date.NowResponse, err error) {

	return middleware_.LocalMiddlewareWrap(c.Controller.Now)(ctx, req)
}

func (c *LocalController) NowError(
	ctx context.Context,
	req *date.NowErrorRequest,
) (resp *date.NowErrorResponse, err error) {
	return middleware_.LocalMiddlewareWrap(c.Controller.NowError)(ctx, req)
}
