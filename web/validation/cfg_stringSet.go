package validation

import (
	"strings"
	//
	"github.com/golangdaddy/tarantula/web"
)

// Returns a validation object which ensures string is whitelisted
func StringSet(set ...string) *Config {

	min := 0.0
	max := STRING_MAX_LENGTH

	filter := map[string]bool{}

	for _, item := range set { filter[item] = true  }

	return NewConfig(
		"",
		func (req web.RequestInterface, param string) (status *web.ResponseStatus, _ interface{}) {

			status, param = checkString(req, min, max, param)
			if status != nil {
				return status, nil
			}

			if !filter[param] { status = req.Respond(400, "PARAM IS NOT PART OF SET: "+strings.Join(set, ", ")) }

			return status, Sanitize(param)
		},
		func (req web.RequestInterface, param interface{}) (status *web.ResponseStatus, _ interface{}) {

			s, ok := param.(string); if !ok { return req.Respond(400, ERR_NOT_STRING), nil }

			status, s = checkString(req, min, max, s)
			if status != nil {
				return status, nil
			}

			if !filter[s] { status = req.Respond(400, "PARAM IS NOT PART OF SET: "+strings.Join(set, ", ")) }

			return status, s
		},
		min,
		max,
	)
}
