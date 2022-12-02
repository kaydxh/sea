package sealet

import (
	"context"

	logs_ "github.com/kaydxh/golang/pkg/logs"
	kitdate_ "github.com/kaydxh/sea/pkg/sealet/domain/kit/date"
)

var _ Repository = (*Sealet)(nil)

type Sealet struct {
	DateRepository kitdate_.Repository
}

type NowRequest struct {
	RequestId string
}

type NowResponse struct {
	Date string
}

func (s *Sealet) Now(ctx context.Context, req *NowRequest) (resp *NowResponse, err error) {
	logger := logs_.GetLoggerOrFallback(ctx, req.RequestId)

	dateReq := &kitdate_.NowRequest{}
	dateResp, err := s.DateRepository.Now(ctx, dateReq)
	if err != nil {
		logger.Errorf("failed to call Date, err: %v", err)
		return nil, err
	}

	resp = &NowResponse{
		Date: dateResp.Date,
	}

	return resp, nil
}
