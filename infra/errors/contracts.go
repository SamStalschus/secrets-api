package apierr

type Provider interface {
	BadRequest(message string, err error) *Message
}
