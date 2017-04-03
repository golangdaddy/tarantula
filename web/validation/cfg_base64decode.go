package validation

import (
	"encoding/base64"
	//
	"github.com/golangdaddy/tarantula/web"
)

// Returns a validation object which checks for valid username
func Base64Decode() *Config {

	return NewConfig(
		"",
		func (req web.RequestInterface, s string) (status *web.ResponseStatus, _ interface{}) {

      		b, err := base64.StdEncoding.DecodeString(s); if err != nil { status = req.Respond(400, err.Error()) }

			return status, b
		},
		func (req web.RequestInterface, param interface{}) (status *web.ResponseStatus, _ interface{}) {

			s, ok := param.(string); if !ok { return req.Respond(400, ERR_NOT_STRING), nil }

			b, err := base64.StdEncoding.DecodeString(s); if err != nil { status = req.Respond(400, err.Error()) }

			return status, b
		},
	)
}
