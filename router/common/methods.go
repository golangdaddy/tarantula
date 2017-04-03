package common

import (
	"github.com/golangdaddy/tarantula/web"
)

// Allows POST requests to the node's handler
func (node *Node) HEAD(functions ...interface{}) *Handler {

	handler := &Handler{}

	if len(functions) > 0 { handler.function = functions[0].(func (req web.RequestInterface) *web.ResponseStatus) }

	node.addHandler("HEAD", handler)

	return handler
}

// Allows GET requests to the node's handler
func (node *Node) GET(functions ...interface{}) *Handler {

	handler := &Handler{}

	if len(functions) > 0 { handler.function = functions[0].(func (req web.RequestInterface) *web.ResponseStatus) }

	node.addHandler("GET", handler)

	return handler
}

// Allows POST requests to the node's handler
func (node *Node) POST(functions ...interface{}) *Handler {

	handler := &Handler{}

	if len(functions) > 0 { handler.function = functions[0].(func (req web.RequestInterface) *web.ResponseStatus) }

	node.addHandler("POST", handler)

	return handler
}

// Allows PUT requests to the node's handler
func (node *Node) PUT(functions ...interface{}) *Handler {

	handler := &Handler{}

	if len(functions) > 0 { handler.function = functions[0].(func (req web.RequestInterface) *web.ResponseStatus) }

	node.addHandler("PUT", handler)

	return handler
}

// Allows POST requests to the node's handler
func (node *Node) DELETE(functions ...interface{}) *Handler {

	handler := &Handler{}

	if len(functions) > 0 { handler.function = functions[0].(func (req web.RequestInterface) *web.ResponseStatus) }

	node.addHandler("DELETE", handler)

	return handler
}

// Allows POST requests to the node's handler
func (node *Node) PATCH(functions ...interface{}) *Handler {

	handler := &Handler{}

	if len(functions) > 0 { handler.function = functions[0].(func (req web.RequestInterface) *web.ResponseStatus) }

	node.addHandler("PATCH", handler)

	return handler
}
