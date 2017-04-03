package validation

import (
	"github.com/golangdaddy/go.uuid"
	//
	"github.com/golangdaddy/tarantula/web"
)

// Returns a validation object that checks for a valid UUID version 4
func UUIDv4() *Config {

	min := 36.0
	max := 36.0

	return NewConfig(
		"98ceed88-756e-4823-96ef-2815eafc0c1e",
		func (req web.RequestInterface, uid string) (*web.ResponseStatus, interface{}) {

			u, err := uuid.FromString(uid)
			if req.Log().Error(err) {
				return req.Respond(400, err), nil
			}

			return nil, u.String()
		},
		func (req web.RequestInterface, param interface{}) (*web.ResponseStatus, interface{}) {

			uid, ok := param.(string); if !ok { return req.Respond(400, ERR_NOT_STRING), nil }

			u, err := uuid.FromString(uid)
			if req.Log().Error(err) {
				return req.Respond(400, err), nil
			}

			return nil, u.String()
		},
		min,
		max,
	)
}
