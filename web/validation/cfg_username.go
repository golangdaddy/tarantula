package validation

import (
	"strings"
	//
	"github.com/golangdaddy/tarantula/web"
)

// Returns a validation object which checks for valid username
func Username(min, max float64) *Config {

	return NewConfig(
		"nameofsomething",
		func (req web.RequestInterface, s string) (status *web.ResponseStatus, _ interface{}) {

			status, s = checkString(
				req,
				min,
				max,
				strings.TrimSpace(strings.ToLower(s)),
			)
			if status != nil {
				return status, nil
			}

			for _, char := range s {

				if !strings.Contains(USERNAME_CHARS, string(char)) { return req.Respond(400, ERR_INVALID_CHARS), nil }

			}

			return status, s
		},
		func (req web.RequestInterface, param interface{}) (status *web.ResponseStatus, _ interface{}) {

			s, ok := param.(string); if !ok { return req.Respond(400, ERR_NOT_STRING), nil }

			status, s = checkString(
				req,
				min,
				max,
				strings.TrimSpace(strings.ToLower(s)),
			)
			if status != nil {
				return status, nil
			}

			for _, char := range s {

				if !strings.Contains(USERNAME_CHARS, string(char)) { return req.Respond(400, ERR_INVALID_CHARS), nil }

			}

			return status, s
		},
		min,
		max,
	)
}
