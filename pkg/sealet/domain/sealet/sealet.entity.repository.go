package sealet

import "context"

type Repository interface {
	Now(ctx context.Context, req *NowRequest) (resp *NowResponse, err error)
	NowError(ctx context.Context, req *NowErrorRequest) (resp *NowErrorResponse, err error)
}
