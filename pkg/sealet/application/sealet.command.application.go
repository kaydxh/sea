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

func (s SealetHandler) Date(ctx context.Context, req *sealet.DateRequest) (resp *sealet.DateResponse, err error) {

	handler, err := s.sealetFactory.NewSealet(ctx)
	if err != nil {
		return nil, err
	}

	return handler.Date(ctx, req)
}
