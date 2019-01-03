package security

import (
	"github.com/golangdaddy/tarantula/router/common/openapi"
)

type Auth_HTTP struct {
	Scheme string
	BearerFormat string
}

func (self *Auth_HTTP) Validate(b []byte) error {
	return nil
}

func (self *Auth_HTTP) Spec() *openapi.SecuritySchemeObject {
	return &openapi.SecuritySchemeObject{
		Type: "http",
		Scheme: self.Scheme,
		BearerFormat: self.BearerFormat,
	}
}
