package date

import (
	"context"

	logs_ "github.com/kaydxh/golang/pkg/logs"
)

// SeaDate 日期领域实体，持有 DateRepository 接口实现业务逻辑。
type SeaDate struct {
	repo DateRepository
}

// NewSeaDate 创建日期领域实体实例。
func NewSeaDate(repo DateRepository) *SeaDate {
	return &SeaDate{repo: repo}
}

// Now 获取当前时间。
func (s *SeaDate) Now(ctx context.Context, req *NowRequest) (*NowResponse, error) {
	logger := logs_.GetLoggerOrFallback(ctx, req.RequestId)

	resp, err := s.repo.Now(ctx, req)
	if err != nil {
		logger.Errorf("failed to call Now, err: %v", err)
		return nil, ErrInternal
	}

	return resp, nil
}

// NowError 获取当前时间（模拟内部错误场景）。
func (s *SeaDate) NowError(ctx context.Context, req *NowErrorRequest) (*NowErrorResponse, error) {
	logger := logs_.GetLoggerOrFallback(ctx, req.RequestId)

	resp, err := s.repo.NowError(ctx, req)
	if err != nil {
		logger.Errorf("failed to call NowError, err: %v", err)
		return nil, ErrInternal
	}

	return resp, nil
}