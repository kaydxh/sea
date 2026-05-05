package date

import "errors"

// 日期领域的业务错误。
// 统一使用 sentinel error，上层通过 errors.Is 进行识别，
// 然后由 web 层 (seadate.error.go) 映射为前端可识别的 code/msg。
var (
	// ErrInternal 内部错误。
	ErrInternal = errors.New("date: internal error")

	// ErrInvalidRequest 请求参数非法。
	ErrInvalidRequest = errors.New("date: invalid request")
)
