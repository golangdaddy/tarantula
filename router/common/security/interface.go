package security

import (
	"github.com/golangdaddy/tarantula/router/common/openapi"
)

type Authentication interface {
	Validate([]byte) error
	Spec() *openapi.SecuritySchemeObject
}
