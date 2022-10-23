package helpers

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/brenos/qap/internal/core/domain"
	"github.com/brenos/qap/internal/core/domain/result"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// VALIDATE EMAIL

func ValidateEmail(email string) error {
	emailSplited := strings.Split(email, "@")
	domainError := validateEmailDomain(emailSplited)
	if domainError != nil {
		return domainError
	}
	domain := strings.ToLower(emailSplited[len(emailSplited)-1])
	return isDisposableEmail(domain)
}

func isDisposableEmail(domain string) error {
	f, _ := os.Open("blacklistEmails.conf")
	for scanner := bufio.NewScanner(f); scanner.Scan(); {
		if scanner.Text() == domain {
			f.Close()
			return errors.New("email is not valid")
		}
	}
	f.Close()
	return nil
}

func validateEmailDomain(domains []string) error {
	if len(domains) != 2 {
		return errors.New("error on email validation")
	}
	return nil
}

// VALIDATE BODY REQUEST
// Ex.: https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/

func ValidateOrCreateBodyRequest(c *gin.Context, bodyRequest any) *domain.Result {
	err := c.ShouldBindJSON(bodyRequest)
	if err != nil {

		var ve validator.ValidationErrors

		if errors.As(err, &ve) {

			out := make([]domain.FieldError, len(ve))

			for i, fe := range ve {
				out[i] = domain.FieldError{Field: fe.Field(), Message: getFieldErrorMsg(fe)}
			}

			return domain.NewResultMessageContextCode(
				"validation schema error",
				out,
				result.CodeUnprocessableEntity,
			)

		}
		return domain.NewResultMessageAndCode("validation schema error", result.CodeUnprocessableEntity)
	}
	return nil
}

func getFieldErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "This field is not a valid email"
	case "min":
		return "Should be less than " + fe.Param()
	case "lte":
		return "Should be less than " + fe.Param()
	case "max":
		return "Should be greater than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	}
	return "Unknown error"
}
