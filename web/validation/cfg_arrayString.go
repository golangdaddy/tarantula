package validation

import (
	//
	"github.com/golangdaddy/tarantula/web"
)

// Returns a validation object that allows any value
func ArrayString() *Config {

	return NewConfig(
		[]string{},
		func (req web.RequestInterface, param string) (status *web.ResponseStatus, _ interface{}) {

			// not really a possible scenario

			return status, param
		},
		func (req web.RequestInterface, param interface{}) (status *web.ResponseStatus, _ interface{}) {

			a, ok := param.([]interface{})
			if !ok {
				return req.Respond(400, ERR_NOT_ARRAY), nil
			}

			b := make([]string, len(a))
			for x, item := range a {
				b[x], ok = item.(string)
				if !ok {
					return req.Respond(400, ERR_NOT_STRING), nil
				}
			}

			return status, b
		},
	)
}
