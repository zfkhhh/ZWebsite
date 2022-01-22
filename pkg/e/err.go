package e

// http返回错误信息

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	SUCCESS        = 200
	INVALID_PARAMS = 400
	ERROR          = 500
)

var MsgFlags = map[int]string{
	SUCCESS:        "请求成功",
	ERROR:          "系统错误",
	INVALID_PARAMS: "请求参数错误",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}

func NewErr(code int) *AppError {
	return &AppError{
		Code:    code,
		Message: GetMsg(code),
	}
}
