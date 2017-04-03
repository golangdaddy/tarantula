package validation

import (
	"time"
	//
	"github.com/golangdaddy/tarantula/web"
)

// Returns a validation object which checks for valid time
func SQLTimestamp() *Config {

	return NewConfig(
		time.Now(),
		func (req web.RequestInterface, param string) (*web.ResponseStatus, interface{}) {

			return nil, param
		},
		func (req web.RequestInterface, param interface{}) (*web.ResponseStatus, interface{}) {

			return nil, param
		},
	)
}

// Returns a validation object which outputs the current time
func Now() *Config {

	return NewConfig(
		time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC).UnixNano(),
		func (req web.RequestInterface, param string) (*web.ResponseStatus, interface{}) {

			return nil, time.Now().UTC().UnixNano()
		},
		func (req web.RequestInterface, param interface{}) (*web.ResponseStatus, interface{}) {

			return nil, time.Now().UTC().UnixNano()
		},
	)
}

// Returns a validation object which checks for valid time like 2017-06-23T00:00:00.000Z
func Time(layout string) *Config {

	return NewConfig(
		"",
		func (req web.RequestInterface, param string) (status *web.ResponseStatus, _ interface{}) {

			return verifyTime(req, layout, param)
		},
		func (req web.RequestInterface, param interface{}) (status *web.ResponseStatus, _ interface{}) {

			s, ok := param.(string); if !ok { status = req.Respond(400, ERR_NOT_STRING) }

			return verifyTime(req, layout, s)
		},
	)
}

func verifyTime(req web.RequestInterface, layout, input string) (*web.ResponseStatus, string) {

    _, err := time.Parse(layout, input); if req.Log().Error(err) { return req.Respond(400, ERR_PARSE_TIME), "" }

    return nil, input
}