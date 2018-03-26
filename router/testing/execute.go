package testing

import (
	"fmt"
	"time"
	"errors"
	"reflect"
	"io/ioutil"
	"encoding/json"
	//
	"github.com/fatih/color"
	//
	"github.com/golangdaddy/tarantula/httpclient"
	"github.com/golangdaddy/tarantula/router/common"
	"github.com/golangdaddy/tarantula/router/testing/manifest"
)

func Execute(m *manifest.Manifest, endpoints ...*manifest.Endpoint) error {

	app := &App{
		GetHandlers(m.Spec),
		m,
		httpclient.NewClient(),
	}

	for _, e := range endpoints {

		time.Sleep(time.Duration(int64(time.Millisecond) * 1000 * int64(e.DelayAfter)))

		err := app.startExecution(e)
		if err != nil {
			return err
		}

		time.Sleep(time.Duration(int64(time.Millisecond) * 1000 * int64(e.DelayAfter)))

	}

	m.AddEndpoints(endpoints...)


	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("manifest.latest.json", b, 0666)
	if err != nil {
		panic(err)
	}


	return nil
}

func (app *App) startExecution(endpoint *manifest.Endpoint) error {

	pathArgs := map[string]interface{}{}
	for k, v := range endpoint.PathArgs {
		x := app.Manifest.Variables[v]
		if x == nil {
			return errors.New("PATH PARAM REFERENCE NOT FOUND: "+v)
		}
		pathArgs[k] = x
	}

	bodyArgs := map[string]interface{}{}
	for k, v := range endpoint.BodyArgs {
		x := app.Manifest.Variables[v]
		if x == nil {
			return errors.New("BODY PARAM REFERENCE NOT FOUND: "+v)
		}
		bodyArgs[k] = x
	}
	for k, v := range endpoint.BodyLiterals {
		bodyArgs[k] = v
	}

	endpoint.Spec = app.GetHandler(
		endpoint.Method,
		endpoint.Endpoint,
		pathArgs,
		bodyArgs,
	)
	if endpoint.Spec == nil {
		fmt.Println(endpoint)
		return fmt.Errorf("FAILED TO EXECUTE NIL SPEC %s %s", endpoint.Method, endpoint.Endpoint)
	}

	obj, err := app.execute(endpoint.Spec)
	if err != nil {
		return err
	}

	if obj == nil {
		return nil
	}

	if object, ok := obj.(map[string]interface{}); !ok {
		if object, ok := obj.([]interface{}); !ok {
			return errors.New("TYPE ASSERTION FAILED: "+reflect.TypeOf(obj).String())
		} else {
			app.Manifest.Variables["array"] = object
		}
	} else {
		for used, as := range endpoint.Use {
			value, ok := object[used].(string)
			if !ok {
				fmt.Println(object)
				panic("TYPE ASSERT FAILED")
			}
			fmt.Println("USING VARIABLE " + used + " AS " + as)
			app.Manifest.Variables[as] = value
		}
	}

	return nil
}

func (app *App) execute(spec *common.HandlerSpec) (interface{}, error) {

	var dst interface{}
	b, err := json.Marshal(spec)
	color.Blue("TESTING SPEC: "+string(b))

	responseSchema, _ := json.Marshal(spec.ResponseSchema)
	switch string(responseSchema[0]) {

		case "{":

			obj := map[string]interface{}{}
			dst = &obj

		case "[":

			array := []interface{}{}
			dst = &array

		default:

	}

	authHeader := map[string]string{
		"Authorization": "Bearer "+app.Manifest.Token,
	}

	switch spec.Method {

		case "GET":

			_, err = app.Get(app.Manifest.Host + spec.MockEndpoint, dst, authHeader)

		case "POST":

			b, _ := json.Marshal(spec.MockPayload)
			color.Yellow(string(b))

			_, err = app.Post(app.Manifest.Host + spec.MockEndpoint, spec.MockPayload, dst, authHeader)

		case "DELETE":

			_, err = app.Delete(app.Manifest.Host + spec.MockEndpoint, nil, authHeader)

	}
	if err != nil {
		return nil, err
	}

	if dst == nil {
		return nil, nil
	}

	b, _ = json.Marshal(dst)
	color.Green(string(b))

	output, ok := dst.(*[]interface{})
	if ok {
		return *output, nil
	}

	return *((dst).(*map[string]interface{})), nil
}
