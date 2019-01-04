package security

import (
	"github.com/golangdaddy/tarantula/router/common/openapi"
)

type Auth_OAuth2 struct {
	Flow string
	AuthorizationUrl string
	TokenUrl string
	Scopes map[string]string
}

func (self *Auth_OAuth2) Validate(b []byte) error {
	return nil
}

func (self *Auth_OAuth2) Spec() *openapi.SecuritySchemeObject {
	return &openapi.SecuritySchemeObject{
		Type: []string{"oauth2"},
		Flow: self.Flow,
		AuthorizationUrl: self.AuthorizationUrl,
		TokenUrl: self.TokenUrl,
		Scopes: self.Scopes,
	}
}
