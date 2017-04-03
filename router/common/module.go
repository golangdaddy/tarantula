package common

import 	(
	"github.com/golangdaddy/tarantula/web"
)

type Middleware struct {
	config *Config
	function MiddlewareFunction
	arg interface{}
}

func (mid *Middleware) Run(req web.RequestInterface) *web.ResponseStatus {

	return mid.function(mid, req, mid.arg)
}

type Module struct {
	config *Config
	function ModuleFunction
	arg interface{}
}

func (mod *Module) Run(req web.RequestInterface) *web.ResponseStatus {

	return mod.function(req, mod.arg)
}
