package seadate

import (
	"context"

	logs_ "github.com/kaydxh/golang/pkg/logs"
	v1 "github.com/kaydxh/sea/api/protoapi-spec/sea-date/v1"
	"github.com/kaydxh/sea/pkg/sea-date/domain/date"
)

// Now 获取当前时间。
func (c *Controller) Now(
	ctx context.Context,
	req *v1.NowRequest,
) (*v1.NowResponse, error) {
	logger := logs_.GetLoggerOrFallback(ctx, req.GetRequestId())

	dateReq := &date.NowRequest{
		RequestId: req.GetRequestId(),
	}
	dateResp, err := c.app.Commands.SeaDateHandler.Now(ctx, dateReq)
	if err != nil {
		logger.WithError(err).Errorf("Now failed")
		code, msg := toResponseCode(err, CodeInternalError)
		return &v1.NowResponse{Code: code, Message: msg}, nil
	}

	return &v1.NowResponse{
		Code:    CodeOK,
		Message: "ok",
		Date:    dateResp.Date,
	}, nil
}
