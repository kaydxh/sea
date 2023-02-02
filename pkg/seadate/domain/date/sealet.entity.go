package date

import (
	"context"

	errors_ "github.com/kaydxh/golang/go/errors"
	logs_ "github.com/kaydxh/golang/pkg/logs"
	kitdate_ "github.com/kaydxh/sea/pkg/seadate/domain/kit/date"
)

var _ Repository = (*SeaDate)(nil)

type SeaDate struct {
	DateRepository kitdate_.Repository
}

type NowRequest struct {
	RequestId string
}

type NowResponse struct {
	Date string
}

type NowErrorRequest struct {
	RequestId string
}

type NowErrorResponse struct {
	Date string
}

func (s *SeaDate) Now(ctx context.Context, req *NowRequest) (resp *NowResponse, err error) {
	logger := logs_.GetLoggerOrFallback(ctx, req.RequestId)

	dateReq := &kitdate_.NowRequest{}
	dateResp, err := s.DateRepository.Now(ctx, dateReq)
	if err != nil {
		logger.Errorf("failed to call Now, err: %v", err)
		return nil, errors_.Errore(
			err,
			ErrInterval,
		)
	}

	resp = &NowResponse{
		Date: dateResp.Date,
	}

	return resp, nil
}

func (s *SeaDate) NowError(ctx context.Context, req *NowErrorRequest) (resp *NowErrorResponse, err error) {
	logger := logs_.GetLoggerOrFallback(ctx, req.RequestId)

	dateReq := &kitdate_.NowErrorRequest{}
	dateResp, err := s.DateRepository.NowError(ctx, dateReq)
	if err != nil {
		logger.Errorf("failed to call NowError, err: %v", err)
		return nil, errors_.Errore(
			err,
			ErrInterval,
		)
	}

	resp = &NowErrorResponse{
		Date: dateResp.Date,
	}

	return resp, nil
}
