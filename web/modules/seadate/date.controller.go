// Package seadate SeaDate gRPC Controller 骨架。
//
//   - date.controller.go : 定义 Controller 结构体与构造函数（本文件）
//   - now.go             : Now 方法
//   - now_error.go       : NowError 方法
//   - seadate.error.go   : domain/application error → (code, msg) 映射
//   - date.router.go     : HTTP 路由、中间件注册
//
// 一个 RPC 方法一个文件，便于阅读与后续新增方法。
package seadate

import (
	v1 "github.com/kaydxh/sea/api/protoapi-spec/sea-date/v1"
	"github.com/kaydxh/sea/pkg/sea-date/application"
)

// Controller SeaDate 控制器，实现 gRPC SeaDateService 接口。
type Controller struct {
	app application.Application

	// 嵌入未实现的 gRPC server，保证前向兼容。
	v1.UnimplementedSeaDateServiceServer
}

// NewController 创建 SeaDate Controller 实例。
func NewController(app application.Application) *Controller {
	return &Controller{app: app}
}