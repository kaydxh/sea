package date

import "context"

type DateRequest struct {
}

type DateResponse struct {
	Date string
}

type Repository interface {
	Date(ctx context.Context, req *DateRequest) (resp *DateResponse, err error)
}
