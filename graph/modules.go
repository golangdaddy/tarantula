package graph

import (
	"fmt"
	"strings"
	//
	"github.com/golangdaddy/tarantula/web"
	//"github.com/golangdaddy/tarantula/router/common"
)

func (db *Database) MOD_authenticate(req web.RequestInterface, arg interface{}) *web.ResponseStatus {

	auths := strings.Split(req.GetHeader("Authorization"), " ")

	username := auths[0]
	token := auths[1]

	ok, session := db.FindSession(token)
	if !ok {
		return req.Respond(403, "NO SESSION FOUND FOR: "+username)
	}

	q := fmt.Sprintf("SELECT * FROM %v WHERE uid = ?;", TABLE_USERS)

	ok, user := db.Client.QueryUser(q, session.User)
	if !ok {
		return req.Fail()
	}

	if username != user.Name {
		return req.Respond(403, "UNAUTHORIZED API CALL")
	}

	req.SetParam("$user", user)

	return nil
}

func (db *Database) MOD_return(req web.RequestInterface, arg interface{}) *web.ResponseStatus {

	key := arg.(string)

	return req.Respond(req.Param(key))
}

func (db *Database) MOD_lookupVertex(req web.RequestInterface, arg interface{}) *web.ResponseStatus {

	outputKey := arg.(string)

	uid := req.Param("uid").(int64)

	ok, vertex := db.Client.QueryVertex(uid)
	if !ok {
		return req.Fail()
	}

	req.SetParam(outputKey, vertex)

	return nil
}

type NewVertex struct {
	Class    *Class
	ParamKey string
}

func (db *Database) MOD_createVertex(req web.RequestInterface, arg interface{}) *web.ResponseStatus {

	args := arg.(*NewVertex)

	var X string

	if req.BodyParam("_") != nil {

		X = req.BodyParam("_").(string)

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
			req.Log().NewError("FAILED TO FIND BODY PARAM  " + args.Class.Name + ":" + property.Key)
			continue
		}

		data[property.Key] = param

	}

	vertex := args.Class.NewVertex(X, data)
	if !vertex.Insert() {
		return req.Fail()
	}

	req.SetParam(args.ParamKey, vertex)

	req.Log().Debug(fmt.Sprintf("CREATED VERTEX: %v", vertex.Uid))

	return nil
}

type Edge struct {
	Phase     bool
	InKey     string
	Predicate string
	OutKey    string
	State     bool
}

func (db *Database) MOD_edgeState(req web.RequestInterface, arg interface{}) *web.ResponseStatus {

	args := arg.(*Edge)

	predicate := db.GetPredicate(args.Predicate)

	in := req.Param(args.InKey).(*Vertex)

	out := req.Param(args.OutKey).(*Vertex)

	if args.Phase {

		in.SetEdgeState(out, predicate.Name, args.State)

	} else {

		out.SetEdgeState(in, predicate.Name, args.State)

	}

	return nil
}

func (db *Database) MOD_getProperty(req web.RequestInterface, arg interface{}) *web.ResponseStatus {

	key := arg.(string)

	subject := req.Param("$subject").(*Vertex)

	if !subject.LoadProperties(key) {
		return req.Fail()
	}

	return req.Respond(subject.GetProperty(key))
}
