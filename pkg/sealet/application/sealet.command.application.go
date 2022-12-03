package application

import (
	"context"

	"github.com/kaydxh/sea/pkg/sealet/domain/sealet"
)

type SealetHandler struct {
	sealetFactory sealet.Factory
}

func NewSealetHandler(f sealet.Factory) SealetHandler {
	return SealetHandler{
		sealetFactory: f,
	}
}

func (s SealetHandler) Now(ctx context.Context, req *sealet.NowRequest) (resp *sealet.NowResponse, err error) {

	handler, err := s.sealetFactory.NewSealet(ctx)
	if err != nil {
		return nil, err
	}

	return handler.Now(ctx, req)
}

func (s SealetHandler) NowError(ctx context.Context, req *sealet.NowErrorRequest) (resp *sealet.NowErrorResponse, err error) {

	handler, err := s.sealetFactory.NewSealet(ctx)
	if err != nil {
		return nil, err
	}

	return handler.NowError(ctx, req)
}
