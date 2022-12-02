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

type DateRequest struct {
	RequestId string
}

type DateResponse struct {
	Date string
}

func (s *Sealet) Date(ctx context.Context, req *DateRequest) (resp *DateResponse, err error) {
	logger := logs_.GetLoggerOrFallback(ctx, req.RequestId)

	dateReq := &kitdate_.DateRequest{}
	dateResp, err := s.DateRepository.Date(ctx, dateReq)
	if err != nil {
		logger.Errorf("failed to call Date, err: %v", err)
		return nil, err
	}

	resp = &DateResponse{
		Date: dateResp.Date,
	}

	return resp, nil
}
