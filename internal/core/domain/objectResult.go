package domain

type Result struct {
	Message string `json:"message"`
	Context any    `json:"context"`
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
