package sealet

import "context"

type Repository interface {
	Date(ctx context.Context, req *DateRequest) (resp *DateResponse, err error)
}
