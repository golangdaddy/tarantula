package validation

import (
	//
	"github.com/golangdaddy/tarantula/web"
)

// Returns a validation object that checks for a bool which parses correctly
func Bool() *Config {

	return NewConfig(
		false,
		func (req web.RequestInterface, param string) (*web.ResponseStatus, interface{}) {

			switch param {

				case "true":	return nil, true
				case "false":	return nil, false

			}

			return req.Respond(400, ERR_NOT_BOOL), nil
		},
		func (req web.RequestInterface, param interface{}) (*web.ResponseStatus, interface{}) {

			b, ok := param.(bool); if !ok { return req.Respond(400, ERR_NOT_BOOL), nil }

			return nil, b
		},
	)
}