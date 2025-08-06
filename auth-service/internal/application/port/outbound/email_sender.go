package outbound

type EmailSender interface {
	SendEmail(to string, subject string, body string) error
	SendEmailWithTemplate(to string, subject string, templateName string, data interface{}) error
}
