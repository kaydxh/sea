package date

import "context"

type NowRequest struct {
}

type NowResponse struct {
	Date string
}

type Repository interface {
	Now(ctx context.Context, req *NowRequest) (resp *NowResponse, err error)
}
