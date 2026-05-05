package date

import (
	"context"
	"fmt"
	"time"

	"github.com/kaydxh/sea/pkg/sea-date/domain/date"
)

var _ date.DateRepository = (*Repository)(nil)

// Repository 本地日期数据访问实现（用于开发/测试）。
type Repository struct {
}

// Now 获取当前时间。
func (r *Repository) Now(ctx context.Context, req *date.NowRequest) (*date.NowResponse, error) {
	return &date.NowResponse{
		Date: time.Now().String(),
	}, nil
}

// NowError 获取当前时间（模拟内部错误场景）。
func (r *Repository) NowError(ctx context.Context, req *date.NowErrorRequest) (*date.NowErrorResponse, error) {
	return nil, fmt.Errorf("internal error")
}
