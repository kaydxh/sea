package date

import (
	"context"
	"fmt"
	"time"

	kitdate_ "github.com/kaydxh/sea/pkg/seadate/domain/kit/date"
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

func (r *Repository) NowError(ctx context.Context, req *kitdate_.NowErrorRequest) (resp *kitdate_.NowErrorResponse, err error) {
	err = fmt.Errorf("Internal")
	return nil, fmt.Errorf("Internal")
}
