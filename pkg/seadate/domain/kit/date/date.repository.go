package date

import "context"

type NowRequest struct {
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

type Repository interface {
	Now(ctx context.Context, req *NowRequest) (resp *NowResponse, err error)
	NowError(ctx context.Context, req *NowErrorRequest) (resp *NowErrorResponse, err error)
}