package date

import (
	"context"
	"time"

	kitdate_ "github.com/kaydxh/sea/pkg/sealet/domain/kit/date"
)

var _ kitdate_.Repository = (*Repository)(nil)

type Repository struct {
}

func (r *Repository) Now(ctx context.Context, req *kitdate_.NowRequest) (resp *kitdate_.NowResponse, err error) {
	resp = &kitdate_.NowResponse{
		Date: time.Now().String(),
	}
	return resp, nil
}
