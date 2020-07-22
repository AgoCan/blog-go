package response

// 错误码
const (
	ErrCodeSuccess        = 0
	ErrCodeParameter      = 1001
	ErrCodeNotLogin       = 1002
	ErrCodeLoginParameter = 1003
)

func getMessage(code int) (message string) {
	switch code {
	case ErrCodeSuccess:
		message = "success"
	case ErrCodeParameter:
		message = "参数错误"
	case ErrCodeNotLogin:
		message = "认证错误"
	case ErrCodeLoginParameter:
		message = "账号密码错误"
	default:
		message = "未知错误"
	}
	return
}
