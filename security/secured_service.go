package multiauth

import (
	"context"
	"log"

	securedservice "github.com/goadesign/examples/security/gen/secured_service"
)

// secured_service service example implementation.
// The example methods log the requests and return zero values.
type securedServicesrvc struct {
	logger *log.Logger
}

// NewSecuredService returns the secured_service service implementation.
func NewSecuredService(logger *log.Logger) securedservice.Service {
	return &securedServicesrvc{logger}
}

// Creates a valid JWT
func (s *securedServicesrvc) Signin(ctx context.Context, p *securedservice.SigninPayload) (res *securedservice.Creds, err error) {
	res = &securedservice.Creds{}
	s.logger.Print("securedService.signin")
	return
}

// This action is secured with the jwt scheme
func (s *securedServicesrvc) Secure(ctx context.Context, p *securedservice.SecurePayload) (res string, err error) {
	s.logger.Print("securedService.secure")
	return
}

// This action is secured with the jwt scheme and also requires an API key
// query string.
func (s *securedServicesrvc) DoublySecure(ctx context.Context, p *securedservice.DoublySecurePayload) (res string, err error) {
	s.logger.Print("securedService.doubly_secure")
	return
}

// This action is secured with the jwt scheme and also requires an API key
// header.
func (s *securedServicesrvc) AlsoDoublySecure(ctx context.Context, p *securedservice.AlsoDoublySecurePayload) (res string, err error) {
	s.logger.Print("securedService.also_doubly_secure")
	return
}
