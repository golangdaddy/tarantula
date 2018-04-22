package validation

import (
	"github.com/golangdaddy/tarantula/web"
)

// Returns a validation object for request body that checks a property to see if it's an object
func MapStringInterface() *Config {

	return NewConfig(
		map[string]interface{}{},
		nil,
		func (req web.RequestInterface, param interface{}) (*web.ResponseStatus, interface{}) {

			x, ok := param.(map[string]interface{}); if !ok { return req.Respond(400, ERR_NOT_OBJECT), nil }

			return nil, x
		},
	)
}
