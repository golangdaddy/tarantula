package query

import (
	"sync"
	//
	"github.com/golangdaddy/tarantula/graph"
)

type Cursor struct {
	Traversal *Traversal
	Stage int
	Index int
	Parent *Cursor
	sync.RWMutex
}

func (qx *QueryExec) LoadList(cursor *Cursor, vtx *graph.Vertex) bool {

	ok, list := cursor.Traversal.Exec(vtx); if !ok { return false }

	qx.Query.DB.Log.DebugJSON(list)

	qx.Lock()

		qx.Lists[cursor.Stage] = list

	qx.Unlock()
	
	return true
}

func (qx *QueryExec) List(cursor *Cursor) []*graph.Link {

	qx.RLock()

		list := qx.Lists[cursor.Stage]

	qx.RUnlock()

	return list
}

func (qx *QueryExec) Vertex(cursor *Cursor) *graph.Vertex {

	qx.RLock()

	edge := qx.Lists[cursor.Stage][cursor.Index]

	qx.RUnlock()

	if cursor.Traversal.Phase {
		return edge.O
	}

	return edge.I
}

func (qx *QueryExec) Next(cursor *Cursor, vtx *graph.Vertex) bool {

	if vtx == nil {

		qx.Query.DB.Log.Debug("USING CURSOR VERTEX")
		vtx = qx.Vertex(cursor)

		qx.Query.DB.Log.Debugf("TRAVERSING - STAGE: %v, INDEX: %v, RESULTS LEN: %v, VERTEX: %v", cursor.Stage, cursor.Index, len(qx.Results), vtx.Uid)

	} else {

		qx.Query.DB.Log.Debugf("TRAVERSING - STAGE: %v, INDEX: %v, RESULTS LEN: %v, VERTEX: %v", cursor.Stage, cursor.Index, len(qx.Results), vtx.Uid)

	}

	if cursor.Index >= len(qx.List(cursor)) {

		if cursor.Parent == nil {

			qx.Query.DB.Log.Debug("TRAVERSE QUIT A")

			return false

		}

		cursor.Index = 0

		cursor.Parent.Index++

		return qx.Next(cursor.Parent, nil)
	}

	// if not last round of traverse
	if cursor.Stage < (qx.Query.TraversalCount - 1) {

		nextCursor := qx.Query.Cursors[cursor.Stage + 1]

		if !qx.LoadList(nextCursor, vtx) { panic("FAILED TO LOAD CURSOR ARRAY") }

		return qx.Next(nextCursor, nil)
	}

	results := make([]*graph.Vertex, len(qx.List(cursor)))

	for i, edge := range qx.List(cursor) {

		//qx.Query.DB.Log.Debugf("READING VERTEX %v %v %v/%v", cursor.Stage, cursor.Index, i, len(cursor.array))

		if cursor.Traversal.Phase {

			results[i] = edge.O

		} else {
			
			results[i] = edge.I

		}

		if int64(len(qx.Results) + i + 1) == qx.Query.Limit { qx.Query.DB.Log.Debug("TRAVERSE QUIT B"); return true }

	}

	qx.Results = append(qx.Results, results...)

	if cursor.Parent == nil { qx.Query.DB.Log.Debug("TRAVERSE QUIT C"); return false }

	return qx.Next(cursor.Parent, nil)
}

type Cursors []*Cursor

func (cursors Cursors) Last() *Cursor {

	return cursors[len(cursors) - 1]
}

func (query *Query) Exec(inputs ...*graph.Vertex) (bool, *QueryExec) {

	qx := &QueryExec{
		query,
		make([][]*graph.Link, query.TraversalCount),
		[]*graph.Vertex{},
		sync.RWMutex{},
	}

	for _, vtx := range inputs {

		cursor := query.Cursors[0]

		if !qx.LoadList(cursor, vtx) { return false, qx }

		qx.Next(cursor, vtx)

	}

	if query.Export {

		query.DB.ExportVertexList(qx.Results)

	}

	return true, qx
}
