package date

import (
	"context"
	"fmt"
	"time"

	kitdate_ "github.com/kaydxh/sea/pkg/sealet/domain/kit/date"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	return nil, status.Error(codes.Internal, err.Error())
}
