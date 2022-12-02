package date

import (
	"context"
	"time"

	kitdate_ "github.com/kaydxh/sea/pkg/sealet/domain/kit/date"
)

var _ kitdate_.Repository = (*Repository)(nil)

type Repository struct {
}

func (r *Repository) Date(ctx context.Context, req *kitdate_.DateRequest) (resp *kitdate_.DateResponse, err error) {
	resp = &kitdate_.DateResponse{
		Date: time.Now().String(),
	}
	return resp, nil
}
