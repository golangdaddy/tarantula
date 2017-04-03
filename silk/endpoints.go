package silk

import (
	//
	"github.com/golangdaddy/tarantula/graph"
	"github.com/golangdaddy/tarantula/web"
)

const (
	EXPORT_HARD_LIMIT = 30
)

func (system *System) parseOptions(req web.RequestInterface) (*web.ResponseStatus, *graph.Vertex, *graph.Vertex, *graph.Predicate, bool) {

	//db.log.DebugJSON(req.Params())

	subject := req.Param("$subject").(*graph.Vertex)

	var predicate *graph.Predicate
	if req.Param("$predicate") != nil {
		predicate = system.DB.GetPredicate(req.Param("$predicate").(string))
	}

	var target *graph.Vertex
	if req.Param("$target") != nil {
		target = req.Param("$target").(*graph.Vertex)
	}

	var phase bool

	if req.Param("$phase") != nil {

		switch req.Param("$phase").(string) {

		case "in":

			phase = true

		case "out":

			phase = false

		default:

			return req.Respond(400, "INVALID PHASE IDENTIFIER: "+req.Param("$phase").(string)), nil, nil, nil, false

		}

	}

	return nil, subject, target, predicate, phase
}

// count edges
func (system *System) gen_edge_count_endpoint(req web.RequestInterface) *web.ResponseStatus {

	status, subject, _, predicate, phase := system.parseOptions(req)
	if status != nil {
		return status
	}

	if phase {
		return req.Respond(system.DBClient.QueryInCount(predicate, subject))
	}

	return req.Respond(system.DBClient.QueryOutCount(predicate, subject))
}

// list edges
func (system *System) gen_edge_list_endpoint_test(req web.RequestInterface) *web.ResponseStatus {

	status, subject, _, _, _ := system.parseOptions(req)
	if status != nil {
		return status
	}

	ok, qx := system.QC.Query("testQuery").Exec(subject); if !ok { return req.Fail() }

	return req.Respond(qx.Results)
}

// list edges
func (system *System) gen_edge_list_endpoint(req web.RequestInterface) *web.ResponseStatus {

	status, subject, _, predicate, phase := system.parseOptions(req)
	if status != nil {
		return status
	}

	var ok bool
	var limit int64 = -1

	switch req.Param("mode").(string) {

		case "list":

			var results []*graph.Vertex

			if phase {
				ok, results = system.DBClient.QueryInList(predicate, subject, limit); if !ok { break }
			} else {
				ok, results = system.DBClient.QueryOutList(predicate, subject, limit); if !ok { break }
			}

			return req.Respond(results)

		case "export":

			if limit < 0 || limit > EXPORT_HARD_LIMIT { limit = EXPORT_HARD_LIMIT }

			var results []*graph.Vertex

			if phase {
				ok, results = system.DBClient.QueryInList(predicate, subject, limit); if !ok { break }
			} else {
				ok, results = system.DBClient.QueryOutList(predicate, subject, limit); if !ok { break }
			}

			system.DB.ExportVertexList(results)

			return req.Respond(results)

		default:

			return req.Respond(400, "INVALID LIST MODE")

	}

	return req.Fail()
}

// change state (in)
func (system *System) gen_edge_link_endpoint(req web.RequestInterface) *web.ResponseStatus {

	state := req.Param("state").(bool)

	status, subject, target, predicate, phase := system.parseOptions(req)
	if status != nil {
		return status
	}

	if phase {

		predicate.NewLink(subject, target, state)
		return nil

	} else {

		predicate.NewLink(target, subject, state)
		return nil

	}

	return req.Respond(400, "INVALID PHASE IDENTIFIER: "+req.Param("$phase").(string))
}

// states

func (system *System) gen_edge_state_endpoint(req web.RequestInterface) *web.ResponseStatus {

	status, subject, target, predicate, phase := system.parseOptions(req)
	if status != nil {
		return status
	}

	if phase {
		return req.Respond(system.DBClient.QueryState(subject, target, predicate, phase))
	}

	return req.Respond(system.DBClient.QueryState(subject, target, predicate, phase))
}

// find entity of a class with x value as the argument
func (system *System) gen_class_x_endpoint(req web.RequestInterface) *web.ResponseStatus {

	x := req.Param("x").(string)

	class := system.DB.ClassNameIndex(req.Param("$class").(string))

	ok, vtx := system.DBClient.QueryVertexByX(x, class); if !ok { return req.Respond(404, "VERTEX NOT FOUND") }

	return req.Respond(vtx)
}

// count entities belonging to a class
func (system *System) gen_class_count_endpoint(req web.RequestInterface) *web.ResponseStatus {

	class := system.DB.ClassNameIndex(req.Param("$class").(string))

	return req.Respond(system.DBClient.QueryClassCount(class))
}

// list entities belonging to a class
func (system *System) gen_class_list_endpoint(req web.RequestInterface) *web.ResponseStatus {

	class := system.DB.ClassNameIndex(req.Param("$class").(string))

	var limit int64 = -1
	var page int64 = -1

	if req.Param("limit") != nil { limit = req.Param("limit").(int64) }
	if req.Param("page") != nil { page = req.Param("page").(int64) }

	switch req.Param("mode").(string) {

		case "list":

			ok, results := system.DBClient.QueryClassList(class, limit, page); if !ok { return req.Fail() }
			return req.Respond(results)

		case "export":

			if limit > EXPORT_HARD_LIMIT { limit = EXPORT_HARD_LIMIT }

			ok, results := system.DBClient.QueryClassList(class, limit, page); if !ok { return req.Fail() }

			system.DB.ExportVertexList(results)

			return req.Respond(results)

	}

	return req.Fail()
}

// list class edges
func (sysClass *SystemClass) gen_edge_class_list_endpoint(req web.RequestInterface) *web.ResponseStatus {

	status, subject, _, predicate, phase := sysClass.System.parseOptions(req)
	if status != nil {
		return status
	}

	var limit int64 = -1
	var page int64 = -1
	var export bool

	if req.Param("mode").(string) == "export" { export = true }
	if req.Param("limit") != nil { limit = req.Param("limit").(int64) }
	if req.Param("page") != nil { page = req.Param("page").(int64) }

	var ok bool
	var results []*graph.Vertex

	if phase {

		ok, results = sysClass.DB.Client.QueryInClassList(sysClass.Class, predicate, subject, limit, page); if !ok { return req.Fail() }

	} else {

		ok, results = sysClass.DB.Client.QueryOutClassList(sysClass.Class, predicate, subject, limit, page); if !ok { return req.Fail() }

	}

	if export {
		sysClass.DB.ExportVertexList(results)
	}

	return req.Respond(results)
}

// count edges
func (sysClass *SystemClass) gen_edge_class_count_endpoint(req web.RequestInterface) *web.ResponseStatus {

	status, subject, _, predicate, phase := sysClass.System.parseOptions(req)
	if status != nil {
		return status
	}

	if phase {
		return req.Respond(sysClass.DB.Client.QueryInClassCount(sysClass.Class, predicate, subject))
	}

	return req.Respond(sysClass.DB.Client.QueryOutClassCount(sysClass.Class, predicate, subject))
}
