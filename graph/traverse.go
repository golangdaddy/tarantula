package graph

import (
)

func (predicate *Predicate) TraverseEdges(uid int64, phase bool, limit int64) []*Link {

	var results []*Link

	if phase {

		results = predicate.In.ActiveEdges(uid, limit)

	} else {

		results = predicate.Out.ActiveEdges(uid, limit)

	}

	return results
}

func (predicate *Predicate) Traverse(uid int64, phase, export bool, limit int64) []*Vertex {

	var results []*Vertex

	if phase {

		results = predicate.In.Active(uid, phase, limit)

	} else {

		results = predicate.Out.Active(uid, phase, limit)

	}

	if export { predicate.DB.ExportVertexList(results) }

	return results
}

func (predicate *Predicate) TraverseCount(uid int64, phase bool) int64 {

	var result int64

	if phase {

		result = predicate.In.ActiveCount(uid)

	} else {

		result = predicate.Out.ActiveCount(uid)

	}

	return result
}


