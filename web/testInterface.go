package web

import (
	"io"
	"os"
	"net/http"
	//
	"github.com/golangdaddy/tarantula/log"
	"github.com/golangdaddy/tarantula/log/testing"
)

func NewTestInterface(method, path string) RequestInterface {

	return &TestInterface{
		method: method,
		path: path,
		device: "Mobile",
		params: map[string]interface{}{},
		bodyParams: map[string]interface{}{},
		headers: map[string]string{},
		log: logs.NewClient().NewLogger("TestInterface"),
	}
}

type TestInterface struct {
	method string
	device string
	path string
	params map[string]interface{}
	bodyParams map[string]interface{}
	headers map[string]string
	log logging.Logger
}

func (ti *TestInterface) FullPath() string { return ti.path }

func (ti *TestInterface) IsTLS() bool { return false }

func (ti *TestInterface) Method() string { return ti.method }

func (ti *TestInterface) Device() string { return ti.method }

func (ti *TestInterface) Body(s string) interface{} { return 0 }

func (ti *TestInterface) Param(k string) interface{} { return ti.params[k] }
func (ti *TestInterface) Params() map[string]interface{} { return ti.params }
func (ti *TestInterface) SetParam(k string, i interface{}) { ti.params[k] = i }
func (ti *TestInterface) SetParams(m map[string]interface{}) { ti.params = m }

func (ti *TestInterface) BodyParam(k string) interface{} { return ti.bodyParams[k] }
func (ti *TestInterface) BodyParams() map[string]interface{} { return ti.bodyParams }
func (ti *TestInterface) SetBodyParam(k string, i interface{}) { ti.bodyParams[k] = i }
func (ti *TestInterface) SetBodyParams(m map[string]interface{}) { ti.bodyParams = m }

func (ti *TestInterface) SetHeader(k, v string) { ti.headers[k] = v }
func (ti *TestInterface) GetHeader(k string) string { return ti.headers[k] }

func (ti *TestInterface) RawBody() (*ResponseStatus, []byte) { return nil, []byte{} }

func (ti *TestInterface) ReadBodyObject() *ResponseStatus { return nil }
func (ti *TestInterface) ReadBodyArray() *ResponseStatus { return nil }

func (ti *TestInterface) BodyObject() map[string]interface{} { return map[string]interface{}{} }
func (ti *TestInterface) BodyArray() []interface{} { return []interface{}{} }

func (ti *TestInterface) Redirect(s string, x int) *ResponseStatus { return nil }

func (ti *TestInterface) ServeFile(s string) { }

func (ti *TestInterface) HttpError(s string, x int) { }

func (ti *TestInterface) Writer() io.Writer { return &rW{} }
func (ti *TestInterface) Write(b []byte) { }

func (ti *TestInterface) Fail() *ResponseStatus { return Fail() }

func (ti *TestInterface) Respond(args ...interface{}) *ResponseStatus { return Respond(args...) }

func (ti *TestInterface) Log() logging.Logger { return ti.log }

func (ti *TestInterface) Res() http.ResponseWriter { return &rW{} }

func (ti *TestInterface) R() *http.Request { return &http.Request{} }

type rW struct {
	status int
	size   int
	http.ResponseWriter
}

func (w *rW) Status() int {
	return w.status
}

func (w *rW) Size() int {
	return w.size
}

func (w *rW) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w *rW) Write(data []byte) (int, error) {

	written, err := os.Stdin.Write(data)
	w.size += written

	return written, err
}

func (w *rW) WriteHeader(statusCode int) {

	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
