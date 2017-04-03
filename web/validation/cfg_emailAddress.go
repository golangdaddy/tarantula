package validation

import (
	"strings"
	"net/mail"
	//
	"github.com/golangdaddy/tarantula/web"
)

// Returns a validation object which checks for valid email address
func EmailAddress() *Config {

	min := 6.0
	max := 254.0

	return NewConfig(
		CONST_MOCK_EMAIL,
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

			_, err := mail.ParseAddress(s); if req.Log().Error(err) {
				return req.Respond(400, "PARAM IS INVALID EMAIL ADDRESS"), nil
			}

			return nil, s
		},
		func (req web.RequestInterface, param interface{}) (status *web.ResponseStatus, _ interface{}) {

			s, ok := param.(string); if !ok { return nil, nil }

			status, s = checkString(
				req,
				min,
				max,
				strings.TrimSpace(strings.ToLower(s)),
			)
			if status != nil {
				return status, nil
			}

			_, err := mail.ParseAddress(s); if req.Log().Error(err) {
				return req.Respond(400, "PARAM IS INVALID EMAIL ADDRESS"), nil
			}

			return nil, s
		},
		min,
		max,
	)
}

// Returns a validation object which checks for valid email address
func EmailAddressOptional() *Config {

	min := 0.0
	max := 254.0

	return NewConfig(
		CONST_MOCK_EMAIL,
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

			if len(s) >= 6 {
				_, err := mail.ParseAddress(s); if req.Log().Error(err) {
					return req.Respond(400, "PARAM IS INVALID EMAIL ADDRESS"), nil
				}
			}

			return nil, s
		},
		func (req web.RequestInterface, param interface{}) (status *web.ResponseStatus, _ interface{}) {

			s, ok := param.(string); if !ok { return req.Respond(400, "PARAM IS NOT A STRING"), nil }

			status, s = checkString(
				req,
				min,
				max,
				strings.TrimSpace(strings.ToLower(s)),
			)
			if status != nil {
				return status, nil
			}

			if len(s) >= 6 {
				_, err := mail.ParseAddress(s); if req.Log().Error(err) {
					return req.Respond(400, "PARAM IS INVALID EMAIL ADDRESS"), nil
				}
			}

			return nil, s
		},
		min,
		max,
	)
}
