package security

import (
	"github.com/golangdaddy/tarantula/router/common/openapi"
)

type Auth_Basic struct {}

func (self *Auth_Basic) Validate(b []byte) error {
	return nil
}

func (self *Auth_Basic) Spec() *openapi.SecuritySchemeObject {
	return &openapi.SecuritySchemeObject{
		Type: "basic",
	}
}
