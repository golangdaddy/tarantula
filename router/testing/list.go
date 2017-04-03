package testing

import (
	"fmt"
	//
	"github.com/golangdaddy/tarantula/router/common"
)

func GetHandlers(url string) (list []*common.HandlerSpec) {

	app := NewApp()

	services := map[string][]*common.HandlerSpec{}
	_, err := app.Get(url, &services)
	if err != nil {
		panic(err)
	}

	var i int
	for serviceName, handlers := range services {
		for _, handler := range handlers {

			fmt.Println(i, serviceName, handler.Method, handler.Endpoint)

			list = append(list, handler)

			i++

		}
	}

	return list
}

func (app *App) GetHandler(method, path string, params ...map[string]interface{}) *common.HandlerSpec {

	for _, spec := range app.List {
		if spec.Method == method && spec.Endpoint == path {

			s := *spec
			newSpec := &s

			if len(params) > 0 {
				newSpec.NewMockPath(params[0])
			}
			if len(params) > 1 {
				newSpec.NewMockPayload(params[1])
			}

			return newSpec
		}
	}

	return nil
}
