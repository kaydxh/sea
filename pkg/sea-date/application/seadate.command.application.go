package application

import (
	"context"

	"github.com/kaydxh/sea/pkg/sea-date/domain/date"
)

// SeaDateHandler 日期应用层处理器。
type SeaDateHandler struct {
	repo date.DateRepository
}

// NewSeaDateHandler 创建日期应用层处理器。
func NewSeaDateHandler(repo date.DateRepository) SeaDateHandler {
	return SeaDateHandler{repo: repo}
}

// Now 获取当前时间。
func (h SeaDateHandler) Now(ctx context.Context, req *date.NowRequest) (*date.NowResponse, error) {
	return h.repo.Now(ctx, req)
}

// NowError 获取当前时间（模拟内部错误场景）。
func (h SeaDateHandler) NowError(ctx context.Context, req *date.NowErrorRequest) (*date.NowErrorResponse, error) {
	return h.repo.NowError(ctx, req)
}