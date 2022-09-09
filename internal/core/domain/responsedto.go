package domain

import "go/types"

type Response struct {
	Message string       `json:"message"`
	Code    int          `json:"code"`
	Context types.Object `json:"context"`
}

func NewMessageResponse(message string) *Response {
	return &Response{
		Message: message,
	}
}

func NewMessageCodeResponse(message string, code int) *Response {
	return &Response{
		Message: message,
		Code:    code,
	}
}

func NewMessageContextResponse(message string, context types.Object) *Response {
	return &Response{
		Message: message,
		Context: context,
	}
}

func NewMessageCodeContextResponse(message string, code int, context types.Object) *Response {
	return &Response{
		Message: message,
		Code:    code,
		Context: context,
	}
}
