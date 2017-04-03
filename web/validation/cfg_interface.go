package validation

import (
	//
	"github.com/golangdaddy/tarantula/web"
)

// Returns a validation object that allows any value
func Interface() *Config {

	return NewConfig(
		0,
		func (req web.RequestInterface, param string) (status *web.ResponseStatus, _ interface{}) {

			return status, param
		},
		func (req web.RequestInterface, param interface{}) (status *web.ResponseStatus, _ interface{}) {

			return status, param
		},
	)
}
