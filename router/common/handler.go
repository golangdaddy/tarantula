package common

import	(
		"fmt"
		"sync"
		"path"
		"mime"
		"bytes"
		"reflect"
		"strings"
		"net/http"
		"io/ioutil"
		"encoding/json"
		//
		"github.com/golangdaddy/tarantula/web"
		"github.com/golangdaddy/tarantula/web/validation"
		)

type Handler struct {
	Config *Config
	Node *Node
	Method string
	Description string
	function func (req web.RequestInterface) *web.ResponseStatus
	IsFile bool
	filePath string
	fileType string
	fileCache []byte
	responseSchema interface{}
	payloadSchema []interface{}
	patchSchema []interface{}
	clientJS *bytes.Buffer
	sync.RWMutex
}

func (handler *Handler) DetectContentType(req web.RequestInterface, filePath string) *web.ResponseStatus {

	if handler.fileCache == nil || !handler.Node.Config.cacheFiles {

		// handle potential trailing slash on folder path declaration
		filePath := strings.Replace(filePath, "//", "/", -1)

		b, err := ioutil.ReadFile(filePath); if err != nil { return req.Respond(404, err.Error()) }

		handler.fileCache = b

		handler.fileType = mime.TypeByExtension(path.Ext(filePath))
		if handler.fileType == "" {

			handler.fileType = http.DetectContentType(b)

		}
	}

	return nil
}

func (handler *Handler) ApiUrl() string {

	var name string

	parts := strings.Split(handler.Node.FullPath(), "/")

	for _, part := range parts {

		if len(part) == 0 { continue }

		if string(part[0]) == ":" {

			part = "'+" + part[1:] + "+'"

		}

		name += "/" + part

	}

	return "'" + name + "'"
}

// Describes the function via the spec JSON
func (handler *Handler) Describe(descr string) *Handler {

	handler.Description = descr

	return handler
}

// Applies model which describes request payload
func (handler *Handler) Body(objects ...*Payload) *Handler {

	for _, object := range objects {
		handler.payloadSchema = append(
			handler.payloadSchema,
			object,
		)
	}

	return handler
}

// Applies model which describes request payload
func (handler *Handler) Patch(objects ...*Patch) *Handler {

	for _, object := range objects {
		handler.patchSchema = append(
			handler.payloadSchema,
			object,
		)
	}

	return handler
}

// Applies model which describes request payload
func (handler *Handler) BodyIsObject() *Handler {

	handler.payloadSchema = append(
		handler.payloadSchema,
		&Object{},
	)

	return handler
}

// Applies model which describes request payload
func (handler *Handler) BodyIsArray() *Handler {

	handler.payloadSchema = append(
		handler.payloadSchema,
		&Array{},
	)

	return handler
}

// Applies model which describes request payload
func (handler *Handler) OptionalBody(obj *Optional) *Handler {

	handler.payloadSchema = append(
		handler.payloadSchema,
		obj,
	)

	return handler
}

// Applys model which describes response schema
func (handler *Handler) Response(schema ...interface{}) *Handler {

	handler.responseSchema = schema[0]

	return handler
}

// Validates any payload present in the request body, according to the payloadSchema
func (handler *Handler) ReadPayload(req web.RequestInterface) *web.ResponseStatus {

	// handle payload

	var paramCount int
	var optionalCount int
	var readBodyObject bool
	bodyParams := map[string]interface{}{}
	statusMessages := map[string]*web.ResponseStatus{}

	for _, schema := range handler.payloadSchema {

		switch params := schema.(type) {

			case nil:

				// do nothing

			case []byte:

			// do nothing

			case map[string]interface{}:

				if !readBodyObject {
					status := req.ReadBodyObject(); if status != nil { return status }
					readBodyObject = true
				}

			case []interface{}:

				if !readBodyObject {
					status := req.ReadBodyObject(); if status != nil { return status }
					readBodyObject = true
				}

			case *Array:

				if !readBodyObject {
					status := req.ReadBodyObject(); if status != nil { return status }
					readBodyObject = true
				}

				array := *params

				switch len(array) {

					case 1:

						return req.Respond(400, "INVALID TYPE FOR ARRAY PAYLOAD SCHEMA, EXPECTS 0 OR 2 ARGS (*ValidationConfig, paramKey)")

					case 2:

						vc, ok := array[0].(*validation.Config); if !ok { return req.Respond(500, "INVALID ARRAY PAYLOAD SCHEMA VALIDATION CONFIG") }

						paramKey, ok := array[1].(string); if !ok { return req.Respond(500, "INVALID ARRAY PAYLOAD SCHEMA PARAM KEY") }

						status, array := vc.BodyFunction(req, req.BodyArray()); if status != nil {

							req.Log().DebugJSON(req.BodyArray())
							//return req.Respond(400, "INVALID TYPE FOR ARRAY PAYLOAD ITEM, EXPECTED: "+vc.Type())

							return status
						}

						req.SetParam(paramKey, array)

				}

			case *Object:

				if !readBodyObject {
					status := req.ReadBodyObject(); if status != nil { return status }
					readBodyObject = true
				}


			case *Payload:

				if !readBodyObject {
					status := req.ReadBodyObject(); if status != nil { return status }
					readBodyObject = true
				}

				object := *params

				for key, vc := range object {
					paramCount++
					status, x := vc.BodyFunction(
						req,
						req.Body(key),
					)
					if status != nil {
						status.Message = fmt.Sprintf("%s KEY(%v)", status.MessageString(), key)
						statusMessages[key] = status
					} else {
						bodyParams[key] = x
					}
				}

			case *Optional:

				if !readBodyObject {
					status := req.ReadBodyObject(); if status != nil { return status }
					readBodyObject = true
				}

				object := *params

				for key, vc := range object {
					optionalCount++
					status, x := vc.BodyFunction(
						req,
						req.Body(key),
					)
					if status == nil {
						bodyParams[key] = x
					}
				}

			default:

				return req.Respond(500, "INVALID OPTIONAL PAYLOAD SCHEMA CONFIG TYPE: "+reflect.TypeOf(params).String())

		}

	}

	if len(statusMessages) > 0 {
		b, _ := json.Marshal(statusMessages)
		return req.Respond(500, string(b))
	}

	lp := len(bodyParams)
	if len(bodyParams) < paramCount {
		return req.Respond(
			400,
			fmt.Sprintf(
				"INVALID PAYLOAD FIELD COUNT %v EXPECTED %v/%v",
				lp,
				paramCount,
				paramCount+optionalCount,
			),
		)
	}

	req.SetBodyParams(bodyParams)

	return nil
}

func (handler *Handler) UseFunction(f interface{}) {

	switch v := f.(type) {

		case func(web.RequestInterface) *web.ResponseStatus:

		  handler.function = v

		default:

		  panic("INVALID ARGUMENT TYPE FOR HANDLER METHOD FUNCTION DECLARATION")

	}
}
