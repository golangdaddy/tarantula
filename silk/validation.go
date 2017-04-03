package silk

import (
	"fmt"
	"strconv"
	//
	"github.com/golangdaddy/tarantula/graph"
	"github.com/golangdaddy/tarantula/web"
	"github.com/golangdaddy/tarantula/web/validation"
)

func (system *System) Subject() *validation.Config {

	return validation.NewConfig(
		&graph.Vertex{},
		func(req web.RequestInterface, id string) (*web.ResponseStatus, interface{}) {

			uid, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				return req.Respond(400, "ERR_PARSE_FLOAT64"), nil
			}

			ok, entity := system.DB.Client.QueryVertex(uid)
			if !ok {
				return req.Respond(400, fmt.Sprintf("VERTEX NOT FOUND: %v", uid)), nil
			}

			return nil, entity
		},
		func(req web.RequestInterface, id interface{}) (status *web.ResponseStatus, _ interface{}) {

			uid := int64(id.(float64))

			ok, entity := system.DB.Client.QueryVertex(uid)
			if !ok {
				status = req.Respond(400, fmt.Sprintf("VERTEX NOT FOUND: %v", uid))
			}

			return status, entity
		},
	)
}

func (system *System) DatabaseClass() *validation.Config {

	return validation.NewConfig(
		&graph.Class{},
		func(req web.RequestInterface, name string) (*web.ResponseStatus, interface{}) {

			if len(name) > graph.MAX_CLASS_NAME {
				return req.Respond(400, "ERR_RANGE_EXCEED"), nil
			}

			class := system.DB.ClassNameIndex(name)

			return nil, class
		},
		func(req web.RequestInterface, name interface{}) (*web.ResponseStatus, interface{}) {

			s, ok := name.(string); if !ok { return req.Respond(400, "ERR_NOT_STRING"), nil }

			if len(s) > graph.MAX_CLASS_NAME {
				return req.Respond(400, "ERR_RANGE_EXCEED"), nil
			}

			class := system.DB.ClassNameIndex(s)

			return nil, class
		},
	)
}

func (system *System) Phase() *validation.Config {

	return validation.NewConfig(
		false,
		func(req web.RequestInterface, name string) (*web.ResponseStatus, interface{}) {

			switch name {

			case "in":
				return nil, true

			case "out":
				return nil, false

			}

			return req.Respond(400, "INVALID PARAM"), false
		},
		func(req web.RequestInterface, name interface{}) (*web.ResponseStatus, interface{}) {

			n, ok := name.(string)
			if !ok {
				return req.Respond(400, "ERR_NOT_STRING"), false
			}

			switch n {

			case "in":
				return nil, true

			case "out":
				return nil, false

			}

			return req.Respond(400, "INVALID PARAM"), false
		},
	)
}
