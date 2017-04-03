package main

import (
	"os"
	"fmt"
	"net/http"
	//
	"github.com/golangdaddy/tarantula/silk"
	"github.com/golangdaddy/tarantula/graph/mysql"
	//"github.com/golangdaddy/tarantula/web"
	"github.com/golangdaddy/tarantula/web/validation"
)

const (
	SUBDOMAIN = "example"
	DOMAIN = "gettarantula.com"
)

func main() {

	var projectId string
	var zone string
	var dbName string
	var password string

	args := os.Args

	switch len(args) {

		case 5:

			projectId = args[1]
			zone = args[2]
			dbName = args[3]
			password = args[4]

		default:

			panic("USAGE: ./example <PROJECT_ID> <ZONE> <DBNAME> <PASSWORD>")

	}

	credentials := &mysql.Credentials{
		ProjectId:  projectId,
		Connection: fmt.Sprintf("%s:%s:%s", projectId, zone, dbName),
		Database:   "silktest",
		Username:   "root",
		Password:   password,
	}

	system, router := silk.NewSystem(SUBDOMAIN, DOMAIN, credentials)

	users := system.AddClass("user")

		users.SetAsUser()

		users.Class.AddProperty("name", validation.Username(3, 16))


	users.Link("following", users)

	// auto generate endpoints for the schema

	system.GenerateAPI()

	// serve the router

	panic(http.ListenAndServe(":80", router.(http.Handler)))
}
