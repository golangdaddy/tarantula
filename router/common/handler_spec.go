package common

import	(
		"fmt"
		"strings"
		"reflect"
		//
		"github.com/golangdaddy/tarantula/web"
		"github.com/golangdaddy/tarantula/web/validation"
		)

type HandlerSpec struct {
	Method string `json:"method"`
	Endpoint string `json:"endpoint"`
	MockEndpoint string `json:"mockEndpoint"`
	MockPayload interface{} `json:"mockPayload,omitempty"`
	PayloadSchema interface{} `json:"payloadSchema,omitempty"`
	OptionalPayloadSchema interface{} `json:"optionalPayloadSchema,omitempty"`
	ResponseSchema interface{} `json:"responseSchema,omitempty"`
	RouteParams map[string]*validation.Config `json:"routeParams,omitempty"`
	IsFile bool `json:"isFile"`
	FilePath string `json:"filePath,omitempty"`
	Description string `json:"description,omitempty"`
}

func (handler *Handler) Spec(req web.RequestInterface) *HandlerSpec {

	// build useful payloadSchema for the spec

	routeParams := map[string]*validation.Config{}
	for _, validationConfig := range handler.Node.Validations {
		routeParams[validationConfig.Keys[0]] = validationConfig
	}

	spec := &HandlerSpec{
		Method:					handler.Method,
		Endpoint:				handler.Node.FullPath(),
		ResponseSchema:			handler.responseSchema,
		IsFile:					handler.IsFile,
		FilePath:				handler.filePath,
		RouteParams:			routeParams,
		Description:			handler.Description,
	}

	for _, schema := range handler.payloadSchema {

		switch data := schema.(type) {

			case nil:

			case *Payload:

				var payloadSchema map[string]*validation.Config
				if schema, ok := spec.PayloadSchema.(map[string]*validation.Config); ok {
					payloadSchema = schema
				} else {
					payloadSchema = map[string]*validation.Config{}
				}
				for k, cfg := range *data { payloadSchema[k] = cfg }
				spec.PayloadSchema = payloadSchema

			case *Optional:

				var payloadSchema map[string]*validation.Config
				if schema, ok := spec.PayloadSchema.(map[string]*validation.Config); ok {
					payloadSchema = schema
				} else {
					payloadSchema = map[string]*validation.Config{}
				}
				for k, cfg := range *data {
					payloadSchema[k] = cfg
				}
				spec.PayloadSchema = payloadSchema

			case *Object:

				spec.PayloadSchema = data

			case *Array:

				spec.PayloadSchema = data

			default:

				spec.PayloadSchema = fmt.Sprintf(
					"unknown_type(%s)",
					reflect.TypeOf(handler.payloadSchema).String(),
				)

		}

	}

	spec.NewMockPath()
	if spec.Method != "GET" {
		spec.NewMockPayload()
	}

	return spec
}

func (spec *HandlerSpec) NewMockPath(patches ...map[string]interface{}) {

	output := []string{}

	path := string(spec.Endpoint[1:])

	for _, part := range strings.Split(path, "/") {

		if string(part[0]) == ":" {

			key := part[1:]

			param := spec.RouteParams[key]

			value := fmt.Sprintf("%v", param.Model)

			for _, patch := range patches {
				v, ok := patch[key]
				if ok {
					value = fmt.Sprintf("%v", v)
				}
			}

			output = append(
				output,
				value,
			)

		} else {
			output = append(
				output,
				part,
			)
		}
	}

	spec.MockEndpoint = "/" + strings.Join(output, "/")
}


func (spec *HandlerSpec) NewMockPayload(patches ...map[string]interface{}) {

	var ok bool
	src := map[string]interface{}{}

	switch schema := spec.PayloadSchema.(type) {

		case nil:

		case []interface{}:

		case map[string]*validation.Config:

			for key, cfg := range schema {

				src[key] = cfg.Model

				for _, patch := range patches {
					v, ok := patch[key]
					if ok {
						src[key] = v
					}
				}

			}

		case map[string]interface{}:

			for key, cfg := range schema {

				src[key], ok = cfg.(map[string]interface{})["model"]
				if !ok {
					panic("map[string]interface{} SCHEMA IS INVALID")
				}

				for _, patch := range patches {
					v, ok := patch[key]
					if ok {
						src[key] = v
					}
				}

			}

		default:

			panic("INVALID PAYLOAD SCHEMA: "+reflect.TypeOf(schema).String())

	}

	if len(src) > 0 {
		spec.MockPayload = src
	}
}

type HandlerArray []*HandlerSpec

func (a HandlerArray) Len() int { return len(a) }

func (a HandlerArray) Swap(x, y int) { a[x], a[y] = a[y], a[x] }

func (a HandlerArray) Less(x, y int) bool {

	if strings.Compare(a[x].Endpoint, a[y].Endpoint) == 1 { return true }

	return false
}
