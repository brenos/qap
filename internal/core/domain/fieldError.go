package domain

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
