// Code generated by goa v2.0.5, DO NOT EDIT.
//
// another_secured_service gRPC client CLI support package
//
// Command:
// $ goa gen github.com/goadesign/examples/security/design

package client

import (
	"encoding/json"
	"fmt"

	anothersecuredservice "github.com/goadesign/examples/security/gen/another_secured_service"
	another_secured_servicepb "github.com/goadesign/examples/security/gen/grpc/another_secured_service/pb"
)

// BuildSigninPayload builds the payload for the another_secured_service signin
// endpoint from CLI flags.
func BuildSigninPayload(anotherSecuredServiceSigninUsername string, anotherSecuredServiceSigninPassword string) (*anothersecuredservice.SigninPayload, error) {
	var username string
	{
		username = anotherSecuredServiceSigninUsername
	}
	var password string
	{
		password = anotherSecuredServiceSigninPassword
	}
	v := &anothersecuredservice.SigninPayload{}
	v.Username = username
	v.Password = password
	return v, nil
}

// BuildSecurePayload builds the payload for the another_secured_service secure
// endpoint from CLI flags.
func BuildSecurePayload(anotherSecuredServiceSecureMessage string, anotherSecuredServiceSecureToken string) (*anothersecuredservice.SecurePayload, error) {
	var err error
	var message another_secured_servicepb.SecureRequest
	{
		if anotherSecuredServiceSecureMessage != "" {
			err = json.Unmarshal([]byte(anotherSecuredServiceSecureMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"fail\": true\n   }'")
			}
		}
	}
	var token string
	{
		token = anotherSecuredServiceSecureToken
	}
	v := &anothersecuredservice.SecurePayload{
		Fail: &message.Fail,
	}
	v.Token = token
	return v, nil
}

// BuildDoublySecurePayload builds the payload for the another_secured_service
// doubly_secure endpoint from CLI flags.
func BuildDoublySecurePayload(anotherSecuredServiceDoublySecureMessage string, anotherSecuredServiceDoublySecureToken string) (*anothersecuredservice.DoublySecurePayload, error) {
	var err error
	var message another_secured_servicepb.DoublySecureRequest
	{
		if anotherSecuredServiceDoublySecureMessage != "" {
			err = json.Unmarshal([]byte(anotherSecuredServiceDoublySecureMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"key\": \"abcdef12345\"\n   }'")
			}
		}
	}
	var token string
	{
		token = anotherSecuredServiceDoublySecureToken
	}
	v := &anothersecuredservice.DoublySecurePayload{
		Key: message.Key,
	}
	v.Token = token
	return v, nil
}

// BuildAlsoDoublySecurePayload builds the payload for the
// another_secured_service also_doubly_secure endpoint from CLI flags.
func BuildAlsoDoublySecurePayload(anotherSecuredServiceAlsoDoublySecureMessage string, anotherSecuredServiceAlsoDoublySecureOauthToken string, anotherSecuredServiceAlsoDoublySecureToken string) (*anothersecuredservice.AlsoDoublySecurePayload, error) {
	var err error
	var message another_secured_servicepb.AlsoDoublySecureRequest
	{
		if anotherSecuredServiceAlsoDoublySecureMessage != "" {
			err = json.Unmarshal([]byte(anotherSecuredServiceAlsoDoublySecureMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"key\": \"abcdef12345\",\n      \"password\": \"password\",\n      \"username\": \"user\"\n   }'")
			}
		}
	}
	var oauthToken *string
	{
		if anotherSecuredServiceAlsoDoublySecureOauthToken != "" {
			oauthToken = &anotherSecuredServiceAlsoDoublySecureOauthToken
		}
	}
	var token *string
	{
		if anotherSecuredServiceAlsoDoublySecureToken != "" {
			token = &anotherSecuredServiceAlsoDoublySecureToken
		}
	}
	v := &anothersecuredservice.AlsoDoublySecurePayload{}
	if message.Username != "" {
		v.Username = &message.Username
	}
	if message.Password != "" {
		v.Password = &message.Password
	}
	if message.Key != "" {
		v.Key = &message.Key
	}
	v.OauthToken = oauthToken
	v.Token = token
	return v, nil
}
