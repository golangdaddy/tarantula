package query

import (
	"sync"
	//
	"github.com/golangdaddy/tarantula/graph"
)

type Client struct {
	DB *graph.Database
	Queries map[string]*Query
	sync.RWMutex
}

func NewClient(db *graph.Database) *Client {

	return &Client{
		DB: db,
		Queries: map[string]*Query{},
	}
}

type QueryExec struct {
	Query *Query
	Lists [][]*graph.Link
	Results []*graph.Vertex
	sync.RWMutex
}

type Query struct {
	DB *graph.Database
	Name string
	Cursors Cursors
	Traversals []*Traversal
	TraversalCount int
	Limit int64
	Export bool
}

func (client *Client) Query(name string) *Query {

	client.RLock()

		query := client.Queries[name]

	client.RUnlock()

	return query
}


func (client *Client) NewQuery(name string) *Query {

	q := &Query{
		DB: client.DB,
		Name: name,
		Traversals: []*Traversal{},
		Limit: -1,
	}

	client.Lock()

		client.Queries[name] = q

	client.Unlock()

	return q
}

func (query *Query) lastTraversal() *Traversal {

	x := len(query.Traversals)

	if x == 0 { panic("THIS QUERY OBJECT HAS NO TRAVERSALS TO FILTER") }

	return query.Traversals[x - 1]
}

func (query *Query) Class(class *graph.Class) *Query {

	traversal := query.lastTraversal()

	traversal.Class = class

	return query
}

func (query *Query) Filter(filters map[string]interface{}) *Query {
	
	traversal := query.lastTraversal()

	traversal.Filters = filters

	return query
}

func (query *Query) Exported() *Query {
	
	query.Export = true

	return query
}

type Traversal struct {
	Predicate *graph.Predicate
	Phase bool
	Filters map[string]interface{}
	Class *graph.Class
}

func (traversal *Traversal) Exec(vtx *graph.Vertex) (bool, []*graph.Link) {

	db := traversal.Predicate.DB

	var ok bool
	var array []*graph.Link

	if traversal.Phase {

		if traversal.Predicate.Sync() {

			array = traversal.Predicate.In.ActiveEdges(vtx.Uid, -1)

			db.Log.Debug("IN QUERY SYNC")
			db.Log.DebugJSON(array)

		} else {

			ok, array = db.Client.TraverseEdges(vtx.Uid, traversal.Predicate, traversal.Phase, -1); if !ok { db.Log.Debug("TRAVERSE QUIT 2"); return false, nil }

			db.Log.Debug("IN QUERY SLOW")
			db.Log.DebugJSON(array)

		}

	} else {

		if traversal.Predicate.Sync() {

			array = traversal.Predicate.Out.ActiveEdges(vtx.Uid, -1)

			db.Log.Debug("OUT QUERY SYNC")
			db.Log.DebugJSON(array)

		} else {

			ok, array = db.Client.TraverseEdges(vtx.Uid, traversal.Predicate, traversal.Phase, -1); if !ok { db.Log.Debug("TRAVERSE QUIT 3"); return false, nil }
		
			db.Log.Debug("OUT QUERY SLOW")
			db.Log.DebugJSON(array)

		}

	}

	return true, array
}

func newTraversal(predicate *graph.Predicate, phase bool) *Traversal {

	return &Traversal{
		Predicate: predicate,
		Phase: true,
		Filters: map[string]interface{}{},
	}
}

func (query *Query) updateTraversals(traversal *Traversal) *Query {

	query.Traversals = append(query.Traversals, traversal)

	query.TraversalCount = len(query.Traversals)

	cl := len(query.Cursors)

	cursor := &Cursor{
		Stage: len(query.Cursors),
		Traversal: traversal,
	}

	if cl > 0 {
		cursor.Parent = query.Cursors[cl - 1]
	}

	query.Cursors = append(query.Cursors, cursor)

	return query
}

func (query *Query) Traverse(predicate *graph.Predicate, phase bool) *Query {

	return query.updateTraversals(newTraversal(predicate, phase))
}

func (query *Query) In(predicate *graph.Predicate) *Query {

	return query.updateTraversals(newTraversal(predicate, true))
}

func (query *Query) Out(predicate *graph.Predicate) *Query {

	return query.updateTraversals(newTraversal(predicate, false))
}

