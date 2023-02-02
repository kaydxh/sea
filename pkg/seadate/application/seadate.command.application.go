package application

import (
	"context"

	"github.com/kaydxh/sea/pkg/seadate/domain/date"
)

type SeaDateHandler struct {
	seaDateFactory date.Factory
}

func NewSeaDateHandler(f date.Factory) SeaDateHandler {
	return SeaDateHandler{
		seaDateFactory: f,
	}
}

func (s SeaDateHandler) Now(ctx context.Context, req *date.NowRequest) (resp *date.NowResponse, err error) {

	handler, err := s.seaDateFactory.NewSeaDate(ctx)
	if err != nil {
		return nil, err
	}

	return handler.Now(ctx, req)
}

func (s SeaDateHandler) NowError(ctx context.Context, req *date.NowErrorRequest) (resp *date.NowErrorResponse, err error) {

	handler, err := s.seaDateFactory.NewSeaDate(ctx)
	if err != nil {
		return nil, err
	}

	return handler.NowError(ctx, req)
}
