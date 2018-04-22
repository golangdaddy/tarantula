package manifest

import (
	"log"
	"reflect"
	"github.com/golangdaddy/tarantula/router/common"
)

type Endpoint struct {
	manifest *Manifest `json:"-"`
	Label string `json:label`
	Method string `json:"method"`
	Endpoint string `json:"endpoint"`
	PathArgs map[string]string `json:"pathArgs"`
	BodyArgs map[string]string `json:"bodyArgs"`
	BodyLiterals map[string]interface{} `json:"bodyLiterals"`
	Use map[string]string `json:"use"`
	Spec *common.HandlerSpec `json:"spec"`
	DelayBefore int `json:"delayBefore"`
	DelayAfter int `json:"delayAfter"`
	eval *func (*Manifest, interface{}) bool
}

func (endpoint *Endpoint) SecDelayBefore(x int) *Endpoint {
	endpoint.DelayBefore = x
	return endpoint
}

func (endpoint *Endpoint) SecDelayAfter(x int) *Endpoint {
	endpoint.DelayAfter = x
	return endpoint
}

func (endpoint *Endpoint) NewPathArgs(args map[string]string) *Endpoint {
	endpoint.PathArgs = args
	return endpoint
}

func (endpoint *Endpoint) NewBodyArgs(args map[string]string) *Endpoint {
	endpoint.BodyArgs = args
	return endpoint
}

func (endpoint *Endpoint) NewBodyLiterals(args map[string]interface{}) *Endpoint {
	endpoint.BodyLiterals = args
	return endpoint
}

func (endpoint *Endpoint) NewUsage(args map[string]string) *Endpoint {
	endpoint.Use = args
	return endpoint
}

func (endpoint *Endpoint) Eval(f func (*Manifest, interface{}) bool) *Endpoint {
	endpoint.eval = &f
	return endpoint
}

func (endpoint *Endpoint) Evaluate(manifest *Manifest, x interface{}) bool {

	switch v := x.(type) {

		case []interface{}:

			if endpoint.eval == nil {
				return true
			}
			return (*endpoint.eval)(manifest, v)

		case map[string]interface{}:

			if endpoint.eval == nil {
				return true
			}
			return (*endpoint.eval)(manifest, v)
	}

	log.Printf("INVALID TYPE: %s", reflect.TypeOf(x).String())

	return false
}
