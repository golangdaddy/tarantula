package validation

import (
	"github.com/golangdaddy/tarantula/web"
)

// Returns a validation object for request body that checks a property to see if it's an array
func ArrayInterface() *Config {

	return NewConfig(
		[]interface{}{},
		nil,
		func (req web.RequestInterface, param interface{}) (*web.ResponseStatus, interface{}) {

			m, ok := param.([]interface{}); if !ok { return req.Respond(400, ERR_NOT_ARRAY), nil }

			return nil, m
		},
	)
}
