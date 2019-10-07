package multiauth

import (
	"context"
	"log"

	anothersecuredservice "github.com/goadesign/examples/security/gen/another_secured_service"
)

// another_secured_service service example implementation.
// The example methods log the requests and return zero values.
type anotherSecuredServicesrvc struct {
	logger *log.Logger
}

// NewAnotherSecuredService returns the another_secured_service service
// implementation.
func NewAnotherSecuredService(logger *log.Logger) anothersecuredservice.Service {
	return &anotherSecuredServicesrvc{logger}
}

// Creates a valid JWT
func (s *anotherSecuredServicesrvc) Signin(ctx context.Context, p *anothersecuredservice.SigninPayload) (res *anothersecuredservice.Creds, err error) {
	res = &anothersecuredservice.Creds{}
	s.logger.Print("anotherSecuredService.signin")
	return
}

// This action is secured with the jwt scheme
func (s *anotherSecuredServicesrvc) Secure(ctx context.Context, p *anothersecuredservice.SecurePayload) (res string, err error) {
	s.logger.Print("anotherSecuredService.secure")
	return
}

// This action is secured with the jwt scheme and also requires an API key
// query string.
func (s *anotherSecuredServicesrvc) DoublySecure(ctx context.Context, p *anothersecuredservice.DoublySecurePayload) (res string, err error) {
	s.logger.Print("anotherSecuredService.doubly_secure")
	return
}

// This action is secured with the jwt scheme and also requires an API key
// header.
func (s *anotherSecuredServicesrvc) AlsoDoublySecure(ctx context.Context, p *anothersecuredservice.AlsoDoublySecurePayload) (res string, err error) {
	s.logger.Print("anotherSecuredService.also_doubly_secure")
	return
}
