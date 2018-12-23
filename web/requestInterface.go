package web

import 	(
		"io"
		"strconv"
		"reflect"
		"net/http"
		//
		"github.com/golangdaddy/tarantula/log"
		)

type RequestInterface interface {
	UID() (string, error)
	FullPath() string
	IsTLS() bool
	Method() string
	Device() string
	Body(string) interface{}
	// accesses the request params of the payload
	Param(string) interface{}
	Params() map[string]interface{}
	SetParam(string, interface{})
	SetParams(map[string]interface{})
	BodyParam(string) interface{}
	BodyParams() map[string]interface{}
	SetBodyParam(string, interface{})
	SetBodyParams(map[string]interface{})
	SetHeader(string, string)
	GetHeader(string) string
	RawBody() (*ResponseStatus, []byte)
	ReadBodyObject() *ResponseStatus
	ReadBodyArray() *ResponseStatus
	BodyArray() []interface{}
	BodyObject() map[string]interface{}
	Redirect(string, int) *ResponseStatus
	ServeFile(string)
	HttpError(string, int)
	Writer() io.Writer
	Write([]byte)
	Fail() *ResponseStatus
	Respond(args ...interface{}) *ResponseStatus
	// logging
	Log() logging.Logger
	//
	Res() http.ResponseWriter
	R() interface{}
}

type ResponseStatus struct {
	Value interface{} `json:"value,omitempty"`
	Code int `json:"code"`
	Message interface{} `json:"message"`
}

func (status *ResponseStatus) MessageString() string {

	switch v := status.Message.(type) {

		case nil:

			return "null"

		case error:

			return v.Error()

		case string:

			return v

	}

	return "INVALID STATUS MESSAGE TYPE: "+reflect.TypeOf(status.Message).String()
}

// returns a standard 500 http error status
func Fail() *ResponseStatus {

	return Respond(500, "UNEXPECTED APPLICATION ERROR")
}

func Respond(args ...interface{}) *ResponseStatus {

	var ok bool
	s := &ResponseStatus{}

	l := len(args)

	switch l {

		case 1:

			s.Value = args[0]
			s.Code = 200
			return s

		case 2, 3:

			s.Code, ok = args[0].(int); if !ok {
				return &ResponseStatus{nil, 501, "Respond(...) METHOD HAS 2 ARGS; UNEXPECTED ARG 0 TYPE: " + reflect.TypeOf(args[0]).String()}
			}

			// argument 1 is now an interface, so we can handle errors

			if args[1] == nil {
				panic("2nd ARGUEMENT TO RESPOND IS NIL")
			}

			s.Message = args[1]

			if l == 3 {
				s.Value = args[2]
			}

			return s

		default:

			return &ResponseStatus{nil, 400, "INVALID STATUS ARGS LENGTH: "+strconv.Itoa(len(args))}

	}

	return nil // Unreachable code warning
}
