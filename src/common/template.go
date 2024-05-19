package common

import "fmt"

const (
	BaseTemplateFolder = "files/templates"
	ForgotPassword     = "forgotpass"
)

var (
	BaseEmailTemplateFolder = fmt.Sprintf("%s/email", BaseTemplateFolder)

	EmailTemplate = map[string]string{
		ForgotPassword: fmt.Sprintf("%s/%s.html", BaseEmailTemplateFolder, ForgotPassword),
	}
)
