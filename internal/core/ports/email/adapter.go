package emailPorts

type EmailAdapter interface {
	SendEmail(emailTo, token string) error
}
