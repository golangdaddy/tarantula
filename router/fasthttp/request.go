package router

import 	(
		"io"
		"net/http"
		"encoding/json"
		//
		"github.com/valyala/fasthttp"
		"github.com/golangdaddy/go.uuid"
		//"github.com/hjmodha/goDevice"
		//
		"github.com/golangdaddy/tarantula/log"
		"github.com/golangdaddy/tarantula/router/common"
		"github.com/golangdaddy/tarantula/web"
		)

type Request struct {
	ctx *fasthttp.RequestCtx
	config *common.Config
	path string
	Node *common.Node
	method string
	params map[string]interface{}
	bodyParams map[string]interface{}
	Object map[string]interface{}
	Array []interface{}
}

func NewRequestObject(node *common.Node, ctx *fasthttp.RequestCtx) *Request {

	return &Request{
		ctx:			ctx,
		config:			node.Config,
		Node:			node,
		method:			string(ctx.Method()),
		params:			node.RequestParams(),
		bodyParams:		map[string]interface{}{},
		Object:			common.Object{},
		Array:			common.Array{},
	}
}

func (req *Request) UID() (string, error) {

	uid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return uid.String(), nil
}

func (req *Request) Log() logging.Logger {

	return req.config.Log
}

func (req *Request) Config() *common.Config {

	return req.config
}

func (req *Request) Res() http.ResponseWriter {

	x := new(http.ResponseWriter)

	return *x
}

func (req *Request) R() interface{} {

	return req.ctx
}

func (req *Request) IsTLS() bool {

	return req.ctx.IsTLS()
}

func (req *Request) BodyArray() []interface{} {

	return req.Array
}

func (req *Request) BodyObject() map[string]interface{} {

	return req.Object
}

func (req *Request) FullPath() string {

	if len(req.path) == 0 {

		req.path = req.Node.FullPath()

	}

	return req.path
}

func (req *Request) Method() string {

	return req.method
}

func (req *Request) Device() string {

	//r := req.R().(*fasthttp.RequestCtx)

	return "?"
}

func (req *Request) Writer() io.Writer {
	return req.ctx.Response.BodyWriter()
}

func (req *Request) Write(b []byte) {

	req.ctx.Write(b)
}

func (req *Request) ServeFile(path string) {

	fasthttp.ServeFile(req.ctx, path)
}

func (req *Request) Body(k string) interface{} {

	return req.Object[k]
}

func (req *Request) Param(k string) interface{} { return req.params[k] }
func (req *Request) Params() map[string]interface{} { return req.params }
func (req *Request) SetParam(k string, v interface{}) { req.params[k] = v }
func (req *Request) SetParams(m map[string]interface{}) { req.params = m }

func (req *Request) BodyParam(k string) interface{} { return req.bodyParams[k] }
func (req *Request) BodyParams() map[string]interface{} { return req.bodyParams }
func (req *Request) SetBodyParam(k string, v interface{}) { req.bodyParams[k] = v }
func (req *Request) SetBodyParams(m map[string]interface{}) { req.bodyParams = m }

func (req *Request) GetHeader(k string) string {

	return string(req.ctx.Request.Header.Peek(k))
}

func (req *Request) SetHeader(k, v string) {

	req.ctx.Request.Header.Set(k, v)
}

func (req *Request) RawBody() (*web.ResponseStatus, []byte) {

	b := req.ctx.PostBody()

	if b == nil { return web.Respond(400, "BODY IS NIL"), nil }

	return nil, b
}

func (req *Request) ReadBodyObject() *web.ResponseStatus {

	err := json.Unmarshal(req.ctx.PostBody(), &req.Object); if err != nil { return web.Respond(400, err.Error()) }

	return nil
}

func (req *Request) ReadBodyArray() *web.ResponseStatus {

	err := json.Unmarshal(req.ctx.PostBody(), &req.Array); if err != nil { return web.Respond(400, err.Error()) }

	return nil
}

func (req *Request) Fail() *web.ResponseStatus {

	return web.Fail()
}

func (req *Request) Respond(args ...interface{}) *web.ResponseStatus {

	return web.Respond(args...)
}

func (req *Request) Redirect(path string, code int) *web.ResponseStatus {

	req.ctx.Redirect(path, code)

	return nil
}

func (req *Request) HttpError(msg string, code int) {

	req.ctx.Error(msg, code)
	req.NewError(msg)
}
