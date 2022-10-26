package apierr

//go:generate mockgen -destination=./mocks.go -package=apierr -source=./contracts.go

type Provider interface {
	BadRequest(message string, err error) *Message
	InternalServerError(err error) *Message
	Unauthorized(message string) *Message
}
