package security

import (
	"github.com/golangdaddy/tarantula/router/common/openapi"
)

type Auth_Key struct {
	In string
	Name string
}

func (self *Auth_Key) Validate(b []byte) error {
	return nil
}

func (self *Auth_Key) Spec() *openapi.SecuritySchemeObject {
	return &openapi.SecuritySchemeObject{
		Type: "apiKey",
		In: self.In,
		Name: self.Name,
	}
}
