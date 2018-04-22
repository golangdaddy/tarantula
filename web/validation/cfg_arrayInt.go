package validation

import (
	//
	"github.com/golangdaddy/tarantula/web"
)

// Returns a validation object that checks for a slice of integers
func ArrayInt() *Config {

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

			b := make([]int, len(a))
			for x, item := range a {
				i, ok := item.(int)
				if !ok {
					return req.Respond(400, ERR_NOT_FLOAT64), nil
				}
				b[x] = int(i)
			}

			return status, b
		},
	)
}
