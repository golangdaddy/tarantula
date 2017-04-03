package silk

import (
	"fmt"
	"sync"
	//
	"github.com/golangdaddy/tarantula/graph"
	"github.com/golangdaddy/tarantula/router/common"
	"github.com/golangdaddy/tarantula/web"
	"github.com/golangdaddy/tarantula/web/validation"
)

func (system *System) AddListMethods(node *common.Node, f func(req web.RequestInterface) *web.ResponseStatus) {

	list := node.Add("list", "mode")

		list.GET(f)

		limitList := list.Param(validation.Int64(), "limit")

			limitList.GET(f)
			limitList.Param(validation.Int64(), "page").GET(f)

	list = node.Add("export", "mode")

		list.GET(f)

		limitList = list.Param(validation.Int64(), "limit")

			limitList.GET(f)
			limitList.Param(validation.Int64(), "page").GET(f)

}

// Generates API endpoints.
func (system *System) GenerateAPI() {

	db := system.DB

	system.Root.Add("_anchors").GET(func(req web.RequestInterface) *web.ResponseStatus {

		outputMap := map[string]int64{}

		for _, class := range db.Classes.Index {

			if class.IsAnchor {

				q := fmt.Sprintf("SELECT * FROM %v WHERE c = %d;", db.Table(graph.TABLE_VERTICES), class.Uid)

				ok, vertices := db.Client.QueryVertices(q)
				if !ok {
					return req.Fail()
				}

				for _, vtx := range vertices {
					outputMap[vtx.X] = vtx.Uid
				}

				break
			}

		}

		return req.Respond(outputMap)
	})

	system.Root.Add("_schema").GET(func(req web.RequestInterface) *web.ResponseStatus {

		outputMap := map[string]interface{}{}

		for _, class := range db.Classes.Index {

			outputMap[class.Name] = class.Properties

		}

		return req.Respond(outputMap)
	})

	group := &sync.WaitGroup{}

	group.Add(len(db.Classes.Index))

	// init classes
	for _, class := range db.Classes.Index {

		//if class.Internal {
		//	system.Log.Debug("SKIPPING INTERNAL CLASS: "+class.Name)
		//	return
		//}

		go system.GenClassEndpoints(class, group)

	}

	group.Wait()

	// generate vector api

	group = &sync.WaitGroup{}

	group.Add(len(db.Classes.Relationships))

	for _, relationship := range db.Classes.Relationships {

		go system.VectorAPI(relationship.In, relationship.Out, relationship, group)

	}

	group.Wait()

	for _, predicate := range db.Classes.Predicates {

		go predicate.PredicateService()

	}

	system.Log.Debug("Thunderbirds are Go!")

}

func (system *System) GenClassEndpoints(class *graph.Class, group *sync.WaitGroup) {

	defer group.Done()

	sysClass := system.Class(class.Name)

	if sysClass == nil {
		system.Log.Fatal("FAILED TO GENERATE ENDPOINTS, MISSING CLASS: "+class.Name)
	}

	if class.IsUser {

		sysClass.Node.Add(
			"register",
		).POST(
			system.gen_auth_register,
		).Describe(
			"Register a new user.",
		).Body(
			class.Validation(),
		)

		sysClass.Node.Add(
			"login",
		).POST(
			system.gen_auth_login,
		).Describe(
			"User login.",
		).Body(
			class.Validation(),
		)

	} else {

		// creating new entities
		sysClass.Node.Add(
			"/create",
		).Mod(
			system.DB.MOD_createVertex,
			&graph.NewVertex{
				class,
				"$target",
			},
		).Mod(
			system.DB.MOD_return,
			"$target",
		).POST().Body(
			class.Validation(),
		)

		// deleting old entities
		sysClass.InstanceNode.Add("/delete").POST(
			func(req web.RequestInterface) *web.ResponseStatus {

				if !req.Param("$subject").(*graph.Vertex).Delete() {
					return req.Fail()
				}

				return nil
			},
		)

	}

	// count entities
	sysClass.Node.Add("/count").GET(system.gen_class_count_endpoint)

	// count entities
	sysClass.Node.Add("/x").Param(validation.StringExact(40), "x").GET(system.gen_class_x_endpoint)

	// adds list and export endpoints
	system.AddListMethods(sysClass.Node, system.gen_class_list_endpoint)

	// handles get request which returns the entity object
	sysClass.InstanceNode.GET(
		func(req web.RequestInterface) *web.ResponseStatus {

			subject := req.Param("$subject").(*graph.Vertex)

			if !subject.LoadProperties() {
				return req.Fail()
			}

			return req.Respond(subject)
		},
	)

	// function to lookup required properties associated to a vertex
	x := func(req web.RequestInterface) *web.ResponseStatus {

		subject := req.Param("$subject").(*graph.Vertex)

		if req.Param("keys") == nil {

			if !subject.LoadProperties() {
				return req.Fail()
			}

			return req.Respond(subject.Data)
		}

		props := req.Param("keys").([]string)

		if !subject.LoadProperties(props...) {
			return req.Fail()
		}

		return req.Respond(subject.Data)
	}

	propsEndpoint := sysClass.InstanceNode.Add("/properties")

	propsEndpoint.GET(x)

	propsEndpoint.Add("keys").Param(validation.StringSplit(","), "keys").GET(x)

}

func (system *System) VectorAPI(class, counterClass *graph.Class, relationship *graph.Relationship, group *sync.WaitGroup) {

	defer group.Done()

	db := system.DB

	sysClass := system.Class(class.Name)
	sysCounterClass := system.Class(counterClass.Name)

	in := sysClass.InstanceNode.Add("/in", "$phase").Add(relationship.Predicate, "$predicate")

	in.Add("/set").Param(system.Subject(), "$target").Param(validation.Bool(), "state").POST(system.gen_edge_link_endpoint)

	in.Add("/state").Param(system.Subject(), "$target").POST(system.gen_edge_state_endpoint)

	in.Add("/count").GET(system.gen_edge_count_endpoint)


	// adds list and export endpoints
	system.AddListMethods(in, system.gen_edge_list_endpoint)

	// endpoint for testing queries
	in.Add("/test").GET(system.gen_edge_list_endpoint_test)

	// create vector node
	vector := in.Add(counterClass.Name)

	vector.Add("/count").GET(sysCounterClass.gen_edge_class_count_endpoint)

	// adds list and export endpoints
	system.AddListMethods(vector, sysCounterClass.gen_edge_class_list_endpoint)


	vector.Add("/create").Mod(
		db.MOD_createVertex,
		&graph.NewVertex{
			sysCounterClass.Class,
			"$target",
		},
	).Mod(
		db.MOD_edgeState,
		&graph.Edge{
			true,
			"$subject",
			relationship.Predicate,
			"$target",
			true,
		},
	).Mod(
		db.MOD_return,
		"$target",
	).POST().Body(
		sysCounterClass.Class.Validation(),
	)

	out := sysCounterClass.InstanceNode.Add("/out", "$phase").Add(relationship.Predicate, "$predicate")

	out.Add("/set").Param(system.Subject(), "$target").Param(validation.Bool(), "state").POST(system.gen_edge_link_endpoint)

	out.Add("/state").Param(system.Subject(), "$target").POST(system.gen_edge_state_endpoint)

	out.Add("/count").GET(system.gen_edge_count_endpoint)

	// adds list and export endpoints
	system.AddListMethods(out, system.gen_edge_list_endpoint)

	// endpoint for testing queries
	out.Add("/test").GET(system.gen_edge_list_endpoint_test)

	// create vector node
	vector = out.Add(class.Name)

	vector.Add("/count").GET(sysClass.gen_edge_class_count_endpoint)

	// adds list and export endpoints
	system.AddListMethods(vector, sysClass.gen_edge_class_list_endpoint)

	vector.Add("/create").Mod(
		db.MOD_createVertex,
		&graph.NewVertex{
			sysClass.Class,
			"$target",
		},
	).Mod(
		db.MOD_edgeState,
		&graph.Edge{
			false,
			"$subject",
			relationship.Predicate,
			"$target",
			true,
		},
	).Mod(
		db.MOD_return,
		"$target",
	).POST().Body(
		sysClass.Class.Validation(),
	)

}
