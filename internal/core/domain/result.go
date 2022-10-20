package domain

type Result struct {
	Message string `json:"message,omitempty"`
	Context any    `json:"context,omitempty"`
	Code    int    `json:"code,omitempty"`
}

func NewResultMessage(message string) *Result {
	return &Result{
		Message: message,
	}
}

func NewResultMessageAndContext(message string, context any) *Result {
	return &Result{
		Message: message,
		Context: context,
	}
}

func NewResultMessageAndCode(message string, code int) *Result {
	return &Result{
		Message: message,
		Code:    code,
	}
}

func NewResultMessageContextCode(message string, context any, code int) *Result {
	return &Result{
		Message: message,
		Context: context,
		Code:    code,
	}
}
