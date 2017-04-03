package girlfriend

import 	(
		"errors"
		"reflect"
		"encoding/json"
		//
		"github.com/fatih/color"
		)

func (req *Request) Debug(msg string) { color.Blue(req.FullPath() + ": %v", msg) }

func (req *Request) NewError(msg string) error {

	err := errors.New(req.FullPath() + ": " + msg)

	color.Red(req.FullPath() + ": "+err.Error())

	return err
}

func (req *Request) Error(msg error) { color.Red(req.FullPath() + ": %v", msg) }

func (req *Request) Reflect(e interface{}) {

	msg := "REFLECT VALUE IS NIL"
	if e != nil { msg = "REFLECT VALUE IS "+reflect.TypeOf(e).String() }

	req.NewError(msg)
}

func (req *Request) DebugJSON(i interface{}) {
	b, err := json.Marshal(i); if err != nil { req.Error(err); return }
	req.Debug(string(b))
}
