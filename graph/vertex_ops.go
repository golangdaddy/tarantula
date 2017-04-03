package graph

import (
	//
)

// updates a state of a link between two vertices
func (vertex *Vertex) SetEdgeState(out *Vertex, predicateName string, state bool) bool {

	predicate := vertex.DB.GetPredicate(predicateName)

	ok, link := predicate.NewLink(vertex, out, state); if !ok { return false }

	// cast new edge event
	vertex.DB.Channels.newEdgeState(link)

	return true
}

// publishes the new vertex
func (vertex *Vertex) Insert() bool {

	// insert vertex

	ok, uid := vertex.DB.Client.InsertVertex(vertex.Class, vertex.X); if !ok { return false }

	vertex.Uid = uid

	// insert properties

	vertex.SaveData()

	// cast new vertex event
	vertex.DB.Channels.insertVertex(vertex)

	return true
}

// deletes the existing vertex
func (vertex *Vertex) Delete() bool {

	// delete vertex
	if !vertex.DB.Client.DeleteVertex(vertex.Uid) { return false }

	// cast new vertex event
	vertex.DB.Channels.deleteVertex(vertex)

	return true
}

// saves the properties of the vertex
func (vertex *Vertex) Patch(props map[string]interface{}) bool {

	client := vertex.DB.Client

	for k, v := range props {

		if !client.InsertProperty(vertex, k, v) { return false }

	}

	return true
}

// saves the properties of the vertex
func (vertex *Vertex) SaveData() bool {

	log := vertex.DB.Log

	props := map[string]interface{}{}

	for _, property := range vertex.Class.Properties {

		prop := vertex.Data[property.Key]

		if prop == nil {
			
			log.NewError("NO DATA FOUMD FOR PARAMETER: "+property.Key)
			return false
		}

		props[property.Key] = prop

	}

	client := vertex.DB.Client

	for k, v := range props {

		if !client.InsertProperty(vertex, k, v) { return false }

	}

	return true
}

func (vertex *Vertex) InList(predicateName string, export bool, limit int64) (bool, []*Vertex) {
	
	db := vertex.DB

	predicate := db.GetPredicate(predicateName)

	if db.MemGraph && predicate.Sync() {
		return true, predicate.Traverse(vertex.Uid, true, export, limit)
	}

	return db.Client.Traverse(vertex.Uid, predicate, true, export, limit)
}

func (vertex *Vertex) InCount(predicateName string) int64 {

	db := vertex.DB

	predicate := db.GetPredicate(predicateName)

	if db.MemGraph && predicate.Sync() {
		return predicate.TraverseCount(vertex.Uid, true)
	}
	
	return db.Client.TraverseCount(vertex.Uid, predicate, true)
}

func (vertex *Vertex) OutList(predicateName string, export bool, limit int64) (bool, []*Vertex) {

	db := vertex.DB

	predicate := db.GetPredicate(predicateName)

	if db.MemGraph && predicate.Sync() {
		return true, predicate.Traverse(vertex.Uid, false, export, limit)
	}

	return db.Client.Traverse(vertex.Uid, predicate, false, export, limit)
}

func (vertex *Vertex) OutCount(predicateName string) int64 {

	db := vertex.DB

	predicate := vertex.DB.GetPredicate(predicateName)

	if db.MemGraph && predicate.Sync() {
		return predicate.TraverseCount(vertex.Uid, false)
	}
	
	return db.Client.TraverseCount(vertex.Uid, predicate, false)
}

