package domain

type Result struct {
	Message string `json:"message,omitempty"`
	Context any    `json:"context,omitempty"`
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
