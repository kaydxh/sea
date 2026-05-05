// Package seadate web 层错误码到前端响应的映射。
//
// sea-date 对外返回格式统一为 { "code": int32, "msg": string, "data": {...} }
// 本文件负责把 domain / application 层抛出的 sentinel error
// 映射为具体的 code（业务错误码）和 msg（面向前端的可读提示）。
package seadate

import (
	"errors"

	"github.com/kaydxh/sea/pkg/sea-date/domain/date"
)

// 业务错误码常量。
//
// 编码策略：
//   - 0        : 成功
//   - 1000-1999: InvalidParameter 参数非法
//   - 3000-3999: InternalError 内部错误
//   - 999      : 兜底未知错误
const (
	// CodeOK 成功。
	CodeOK int32 = 0

	// CodeUnknown 兜底错误码。
	CodeUnknown int32 = 999

	// ---- InvalidParameter (1000-1999) ----

	// CodeInvalidParameter 请求参数非法。
	CodeInvalidParameter int32 = 1001

	// ---- InternalError (3000-3999) ----

	// CodeInternalError 内部错误。
	CodeInternalError int32 = 3001
)

// toResponseCode 将 application / domain 层抛出的 error
// 映射为 (code, msg)。defaultCode 用于兜底（未识别的错误）。
func toResponseCode(err error, defaultCode int32) (int32, string) {
	if err == nil {
		return CodeOK, "ok"
	}
	msg := err.Error()

	switch {
	case errors.Is(err, date.ErrInvalidRequest):
		return CodeInvalidParameter, msg
	case errors.Is(err, date.ErrInternal):
		return CodeInternalError, msg
	}

	return defaultCode, msg
}
