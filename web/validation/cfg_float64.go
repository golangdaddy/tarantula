package validation

import (
	"strconv"
	//
	"github.com/golangdaddy/tarantula/web"
)

// Returns a validation object that checks for a float64 which parses correctly
func Float64(ranges ...int) *Config {

	var min float64
	var max float64

	switch len(ranges) {

		case 2:

			min = float64(ranges[0])
			max = float64(ranges[1])

	}

	cfg := NewConfig(
		0.123,
		func (req web.RequestInterface, param string) (status *web.ResponseStatus, _ interface{}) {

			val, err := strconv.ParseFloat(param, 64); if req.Log().Error(err) {
				return req.Respond(400, ERR_PARSE_FLOAT64), nil
			}

			if min + max == 0 { return nil, val }

			if (val > max || val < min) { status = req.Respond(400, ERR_RANGE_EXCEED) }

			return status, val
		},
		func (req web.RequestInterface, param interface{}) (status *web.ResponseStatus, _ interface{}) {

			val, ok := param.(float64); if !ok { return req.Respond(400, ERR_NOT_FLOAT64), 0 }

			if min + max == 0 { return nil, val }

			if (val > max || val < min) { status = req.Respond(400, ERR_RANGE_EXCEED) }

			return status, val
		},
		min,
		max,
	)


	return cfg
}
