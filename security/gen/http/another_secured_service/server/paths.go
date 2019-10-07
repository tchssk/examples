// Code generated by goa v2.0.5, DO NOT EDIT.
//
// HTTP request path constructors for the another_secured_service service.
//
// Command:
// $ goa gen github.com/goadesign/examples/security/design

package server

// SigninAnotherSecuredServicePath returns the URL path to the another_secured_service service signin HTTP endpoint.
func SigninAnotherSecuredServicePath() string {
	return "/another/signin"
}

// SecureAnotherSecuredServicePath returns the URL path to the another_secured_service service secure HTTP endpoint.
func SecureAnotherSecuredServicePath() string {
	return "/another/secure"
}

// DoublySecureAnotherSecuredServicePath returns the URL path to the another_secured_service service doubly_secure HTTP endpoint.
func DoublySecureAnotherSecuredServicePath() string {
	return "/another/secure"
}

// AlsoDoublySecureAnotherSecuredServicePath returns the URL path to the another_secured_service service also_doubly_secure HTTP endpoint.
func AlsoDoublySecureAnotherSecuredServicePath() string {
	return "/another/secure"
}
