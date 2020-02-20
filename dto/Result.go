package dto

import "acai-go/constant"

type Result struct {
	Code    int64       `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Success() Result {
	result := Result{
		Code:    constant.SUCCESS,
		Data:    nil,
		Message: "操作成功",
	}
	return result
}

func SuccessData(data interface{}) Result {
	result := Result{
		Code:    constant.SUCCESS,
		Data:    data,
		Message: "操作成功",
	}
	return result
}

func Error() Result {
	result := Result{
		Code:    constant.ERROR,
		Data:    nil,
		Message: "操作失败",
	}
	return result
}

func ErrorMessage(message string) Result {
	result := Result{
		Code:    constant.ERROR,
		Data:    nil,
		Message: message,
	}
	return result
}
