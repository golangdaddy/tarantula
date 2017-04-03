package silk

import (
	"fmt"
	//
	"github.com/golangdaddy/tarantula/graph"
	"github.com/golangdaddy/tarantula/web"
)

// authentication

func (system *System) gen_auth_register(req web.RequestInterface) *web.ResponseStatus {

	class := system.DB.ClassNameIndex(req.Param("$class").(string))

	username := req.Param("_username").(string)
	password := req.Param("_password").(string)

	passwordHash, _ := newHashedPassword(username, password)

	userQuery := fmt.Sprintf("SELECT * FROM %v WHERE name = ?;", system.DB.Table(graph.TABLE_USERS))

	ok, _ := system.DB.Client.QueryUser(userQuery, username)
	if ok {
		return req.Respond(400, "USER ALREADY EXISTS")
	}

	userVertex := class.NewVertex(username, nil)
	if !userVertex.Insert() {
		return req.Fail()
	}

	q := fmt.Sprintf("INSERT INTO %v VALUES (null, ?, ?, ?);", system.DB.Table(graph.TABLE_USERS))

	if ok, _ := system.DB.Client.Exec(q, username, passwordHash, userVertex.Uid); !ok {
		return req.Fail()
	}

	ok, user := system.DB.Client.QueryUser(userQuery, username)
	if !ok {
		return req.Respond(404, "USER NOT FOUND")
	}

	return req.Respond(user.Vertex)
}

func (system *System) gen_auth_login(req web.RequestInterface) *web.ResponseStatus {

	username := req.Param("_username").(string)
	password := req.Param("_password").(string)

	_, passwordBytes := newHashedPassword(username, password)

	q := fmt.Sprintf("SELECT * FROM %v WHERE name = ?;", system.DB.Table(graph.TABLE_USERS))

	ok, user := system.DB.Client.QueryUser(q, username)
	if !ok {
		return req.Respond(404, "USER NOT FOUND")
	}

	if err := comparePassword(user.PasswordHash, passwordBytes); err != nil {

		req.Log().Error(err)

		return req.Respond(403, "LOGIN ATTEMPT FAIL FOR USER: "+username)
	}

	ok, session := user.NewSession()
	if !ok {
		return req.Fail()
	}

	return req.Respond(session)
}
