package validation

import (
	"strconv"
	//
	"github.com/golangdaddy/tarantula/web"
)

// Returns a validation object that checks for an int which parses correctly
func Int() *Config {

	return NewConfig(
		1,
		func (req web.RequestInterface, param string) (status *web.ResponseStatus, _ interface{}) {

			val, err := strconv.Atoi(param); if req.Log().Error(err) { status = req.Respond(400, "PARAM FAILED TO PARSE AS INT") }

			return status, val
		},
		func (req web.RequestInterface, param interface{}) (status *web.ResponseStatus, _ interface{}) {

			f, ok := param.(float64)

			if !ok {

				i, ok := param.(int); if !ok { status = req.Respond(400, "PARAM IS NOT NUMERIC") }

				return status, i
			}

			return status, int(f)
		},
	)
}

// Returns a validation object that checks for an int which parses correctly and is positive
func PositiveInt() *Config {

	e := "PARAM IS NOT POSITIVE"

	return NewConfig(
		0,
		func (req web.RequestInterface, param string) (status *web.ResponseStatus, _ interface{}) {

			val, err := strconv.Atoi(param); if req.Log().Error(err) { status = req.Respond(400, "PARAM FAILED TO PARSE AS INT") }

			if val <= 0 { status = req.Respond(400, e) }

			return status, val
		},
		func (req web.RequestInterface, param interface{}) (status *web.ResponseStatus, _ interface{}) {

			val, ok := param.(float64); if !ok { status = req.Respond(400, "PARAM IS NOT NUMERIC") }

			if val <= 0 { status = req.Respond(400, e) }

			return status, int(val)
		},
	)
}

// Returns a validation object that checks for an int which parses correctly and is negative
func NegativeInt() *Config {

	e := "PARAM IS NOT NEGATIVE"

	return NewConfig(
		0,
		func (req web.RequestInterface, param string) (status *web.ResponseStatus, _ interface{}) {

			val, err := strconv.Atoi(param); if req.Log().Error(err) { status = req.Respond(400, "PARAM FAILED TO PARSE AS INT") }

			if val >= 0 { status = req.Respond(400, e) }

			return status, val
		},
		func (req web.RequestInterface, param interface{}) (status *web.ResponseStatus, _ interface{}) {

			val, ok := param.(float64); if !ok { status = req.Respond(400, "PARAM IS NOT NUMERIC") }

			if val >= 0 { status = req.Respond(400, e) }

			return status, int(val)
		},
	)
}

// Returns a validation object that checks for an int which parses correctly and is zero or above
func IntOptimistic() *Config {

	e := "PARAM IS NOT OPTIMISTIC"

	return NewConfig(
		0,
		func (req web.RequestInterface, param string) (status *web.ResponseStatus, _ interface{}) {

			val, err := strconv.Atoi(param); if req.Log().Error(err) { status = req.Respond(400, "PARAM FAILED TO PARSE AS INT") }

			if val < 0 { status = req.Respond(400, e) }

			return status, val
		},
		func (req web.RequestInterface, param interface{}) (status *web.ResponseStatus, _ interface{}) {

			val, ok := param.(float64); if !ok { status = req.Respond(400, "PARAM IS NOT NUMERIC") }

			if val < 0 { status = req.Respond(400, e) }

			return status, int(val)
		},
	)
}

// Returns a validation object that checks for an int which parses correctly and is zero or lower
func IntPessimistic() *Config {

	e := "PARAM IS NOT PESSIMISTIC"

	return NewConfig(
		0,
		func (req web.RequestInterface, param string) (status *web.ResponseStatus, _ interface{}) {

			val, err := strconv.Atoi(param); if req.Log().Error(err) { status = req.Respond(400, "PARAM FAILED TO PARSE AS INT") }

			if val > 0 { status = req.Respond(400, e) }

			return status, val
		},
		func (req web.RequestInterface, param interface{}) (status *web.ResponseStatus, _ interface{}) {

			val, ok := param.(float64); if !ok { status = req.Respond(400, "PARAM IS NOT NUMERIC") }

			if val > 0 { status = req.Respond(400, e) }

			return status, int(val)
		},
	)
}
