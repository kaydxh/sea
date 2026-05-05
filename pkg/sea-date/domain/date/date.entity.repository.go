// Package date 日期领域层，定义领域接口和实体。
package date

import "context"

// NowRequest 获取当前时间请求。
type NowRequest struct {
	RequestId string
}

// NowResponse 获取当前时间响应。
type NowResponse struct {
	Date string
}

// NowErrorRequest 获取当前时间（模拟错误）请求。
type NowErrorRequest struct {
	RequestId string
}

// NowErrorResponse 获取当前时间（模拟错误）响应。
type NowErrorResponse struct {
	Date string
}

// DateRepository 日期数据访问接口（领域层定义）。
type DateRepository interface {
	// Now 获取当前时间。
	Now(ctx context.Context, req *NowRequest) (*NowResponse, error)

	// NowError 获取当前时间（模拟内部错误场景）。
	NowError(ctx context.Context, req *NowErrorRequest) (*NowErrorResponse, error)
}