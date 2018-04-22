package validation

import (
	//
	"github.com/golangdaddy/tarantula/web"
)

// Returns a validation object that checks for a string with exact length
func StringExact(max float64) *Config {

	min := max

	return NewConfig(
		"",
		func (req web.RequestInterface, param string) (*web.ResponseStatus, interface{}) {

			return checkString(req, min, max, param)
		},
		func (req web.RequestInterface, param interface{}) (*web.ResponseStatus, interface{}) {

			if min == 0 && param == nil { return nil, "" }

			s, ok := param.(string); if !ok { return req.Respond(400, ERR_NOT_STRING), nil }

			return checkString(req, min, max, s)
		},
		min,
		max,
	)
}
