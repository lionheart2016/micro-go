package model

// Response 通用响应模型
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// NewResponse 创建新的响应
func NewResponse(code int, message string, data interface{}) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// SuccessResponse 创建成功响应
func SuccessResponse(data interface{}) Response {
	return NewResponse(200, "success", data)
}

// ErrorResponse 创建错误响应
func ErrorResponse(code int, message string) Response {
	return NewResponse(code, message, nil)
}
