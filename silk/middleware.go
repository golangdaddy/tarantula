package silk

import (
	"fmt"
	"strings"
	//
	"github.com/golangdaddy/tarantula/router/common"
	"github.com/golangdaddy/tarantula/graph"
	"github.com/golangdaddy/tarantula/web"
)

func (system *System) MW_authenticate(ware *common.Middleware, req web.RequestInterface, arg interface{}) *web.ResponseStatus {

	auths := strings.Split(req.GetHeader("Authorization"), " ")

	username := auths[0]
	token := auths[1]

	ok, session := system.DB.FindSession(token)
	if !ok {
		return req.Respond(403, "NO SESSION FOUND FOR: "+username)
	}

	q := fmt.Sprintf("SELECT * FROM %v WHERE uid = ?;", graph.TABLE_USERS)

	ok, user := system.DB.Client.QueryUser(q, session.User)
	if !ok {
		return req.Fail()
	}

	if username != user.Name {
		return req.Respond(403, "UNAUTHORIZED API CALL")
	}

	req.SetParam("$user", user)

	return nil
}

func (system *System) MW_return(ware *common.Middleware, req web.RequestInterface, arg interface{}) *web.ResponseStatus {

	key := arg.(string)

	return req.Respond(req.Param(key))
}

func (system *System) MW_lookupVertex(ware *common.Middleware, req web.RequestInterface, arg interface{}) *web.ResponseStatus {

	outputKey := arg.(string)

	uid := req.Param("uid").(int64)

	ok, vertex := system.DB.Client.QueryVertex(uid)
	if !ok {
		return req.Fail()
	}

	req.SetParam(outputKey, vertex)

	return nil
}

type NewVertex struct {
	Class    *graph.Class
	ParamKey string
}

func (system *System) MW_createVertex(ware *common.Middleware, req web.RequestInterface, arg interface{}) *web.ResponseStatus {

	for k, v := range req.BodyParams() {
		system.Log.Debug(k)
		system.Log.DebugJSON(v)
	}

	args := arg.(*NewVertex)

	var X string

	if req.Param("__") != nil {

		X = req.Param("__").(string)

		if len(X) == 0 {
			return req.Respond(400, "VERTEX _ (X) REFERENCE IS ZERO LENGTH")
		}

	} else {

		X = random()

	}

	data := map[string]interface{}{}

	for _, property := range args.Class.Properties {

		param := req.BodyParam(property.Key)

		if param == nil {
			system.Log.NewError("FAILED TO FIND BODY PARAM  " + args.Class.Name + " : " + property.Key)
			continue
		}

		data[property.Key] = param

	}

	vertex := args.Class.NewVertex(X, data)

	if !vertex.Insert() {
		return req.Fail()
	}

	req.SetParam(args.ParamKey, vertex)

	return nil
}

type Edge struct {
	Phase     bool
	InKey     string
	Predicate string
	OutKey    string
	State     bool
}

func (system *System) MW_edgeState(ware *common.Middleware, req web.RequestInterface, arg interface{}) *web.ResponseStatus {

	args := arg.(*graph.Edge)

	predicate := system.DB.GetPredicate(args.Predicate)

	in := req.Param(args.InKey).(*graph.Vertex)

	out := req.Param(args.OutKey).(*graph.Vertex)

	if args.Phase {

		in.SetEdgeState(out, predicate.Name, args.State)

	} else {

		out.SetEdgeState(in, predicate.Name, args.State)

	}

	return nil
}

func (system *System) MW_getProperty(ware *common.Middleware, req web.RequestInterface, arg interface{}) *web.ResponseStatus {

	key := arg.(string)

	subject := req.Param("$subject").(*graph.Vertex)

	if !subject.LoadProperties(key) {
		return req.Fail()
	}

	return req.Respond(subject.GetProperty(key))
}
