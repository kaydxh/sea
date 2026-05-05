package seadate

import (
	"context"

	logs_ "github.com/kaydxh/golang/pkg/logs"
	v1 "github.com/kaydxh/sea/api/protoapi-spec/sea-date/v1"
	"github.com/kaydxh/sea/pkg/sea-date/domain/date"
)

// NowError 获取当前时间（模拟内部错误场景）。
func (c *Controller) NowError(
	ctx context.Context,
	req *v1.NowErrorRequest,
) (*v1.NowErrorResponse, error) {
	logger := logs_.GetLoggerOrFallback(ctx, req.GetRequestId())

	dateReq := &date.NowErrorRequest{
		RequestId: req.GetRequestId(),
	}
	dateResp, err := c.app.Commands.SeaDateHandler.NowError(ctx, dateReq)
	if err != nil {
		logger.WithError(err).Errorf("NowError failed")
		code, msg := toResponseCode(err, CodeInternalError)
		return &v1.NowErrorResponse{Code: code, Message: msg}, nil
	}

	return &v1.NowErrorResponse{
		Code:    CodeOK,
		Message: "ok",
		Date:    dateResp.Date,
	}, nil
}
