package common

import (
	"github.com/golangdaddy/tarantula/web"
	"github.com/golangdaddy/tarantula/web/validation"
)

type Router interface{
  Serve(int)
}

type Array []interface{}

type Object map[string]interface{}

type Headers map[string]string

type Registry map[string]func (req web.RequestInterface) *web.ResponseStatus

type ModuleFunction func (web.RequestInterface, interface{}) *web.ResponseStatus

type MiddlewareFunction func (*Middleware, web.RequestInterface, interface{}) *web.ResponseStatus

type ModuleRegistry map[string]ModuleFunction

type Patch map[string]*validation.Config
type Payload map[string]*validation.Config

func (payload Payload) WithFields(fields Payload) Payload {

  for k, v := range payload { fields[k] = v }

  return fields
}

type Optional map[string]*validation.Config
