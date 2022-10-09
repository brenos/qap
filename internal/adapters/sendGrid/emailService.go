package sendgrid

import (
	"fmt"
	"log"

	"github.com/brenos/qap/helpers"
	emailPorts "github.com/brenos/qap/internal/core/ports/email"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type emailAdapter struct{}

func NewEmailAdapter() emailPorts.EmailAdapter {
	return &emailAdapter{}
}

func (*emailAdapter) SendEmail(emailTo string, token string) error {

	from := mail.NewEmail("QAP", "qapsolution@gmail.com")
	subject := "QAP Token"
	to := mail.NewEmail("", emailTo)
	email := mail.NewV3Mail()
	email.SetFrom(from)
	personalization := mail.NewPersonalization()
	personalization.SetDynamicTemplateData("token", token)
	personalization.To = append(personalization.To, to)
	personalization.Subject = subject
	email.SetTemplateID("d-fb162547a9f24a5c8960de2d4bf55b5c")
	email.Personalizations = append(email.Personalizations, personalization)
	client := sendgrid.NewSendClient(helpers.TOKEN_EMAIL())
	response, err := client.Send(email)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	return err
}
