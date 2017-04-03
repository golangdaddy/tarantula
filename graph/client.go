package graph

import (
	"database/sql"
)

type Client interface {
	SetDB(*Database)
	//
	Exec(string, ...interface{}) (bool, sql.Result)
	CreateDatabase(string) bool
	InsertInternalTable(*Class) bool
	InsertClass(string) (bool, int64)
	InsertVertex(*Class, string) (bool, int64)
	InsertProperty(*Vertex, string, interface{}) bool
	DeleteVertex(int64) bool
	//
	QueryClassUID(string) (bool, int64)
	//
	QueryRows(string, ...interface{}) (bool, *sql.Rows)
	CountRows(string, ...interface{}) int64
	//
	QueryUser(string, ...interface{}) (bool, *User)
	QuerySession(string, ...interface{}) (bool, *Session)
	//
	QueryVertex(int64) (bool, *Vertex)
	QueryVertexByX(string, ...*Class) (bool, *Vertex)
	QueryVertices(string, ...interface{}) (bool, []*Vertex)
	//
	QueryEdge(*Predicate, string, ...interface{}) (bool, *Link)
	QueryEdges(*Predicate, string, ...interface{}) (bool, []*Link)
	QueryAllEdges(*Predicate) (bool, []*Link)
	CountAllEdges(*Predicate) (int64)
	//
	QueryLink(string, ...interface{}) (bool, *Link)
	QueryLinks(string, ...interface{}) (bool, []*Link)	
	//
	SearchClassProperties(*Class, string, string) (bool, []*PropertyExport)
	SearchProperties(string, string) (bool, []interface{})
	QueryProperties(*Vertex, ...string) (bool, map[string]interface{})
	//
	QueryState(*Vertex, *Vertex, *Predicate, bool) bool
	QueryClassCount(*Class) int64
	QueryClassList(*Class, int64, int64) (bool, []*Vertex)
	// in
	QueryInCount(*Predicate, *Vertex) int64
	QueryInList(*Predicate, *Vertex, int64) (bool, []*Vertex)
	QueryInClassCount(*Class, *Predicate, *Vertex) int64
	QueryInClassList(*Class, *Predicate, *Vertex, int64, int64) (bool, []*Vertex)
	// out
	QueryOutCount(*Predicate, *Vertex) int64
	QueryOutList(*Predicate, *Vertex, int64) (bool, []*Vertex)
	QueryOutClassCount(*Class, *Predicate, *Vertex) int64
	QueryOutClassList(*Class, *Predicate, *Vertex, int64, int64) (bool, []*Vertex)
	//
	TraverseEdges(int64, *Predicate, bool, int64) (bool, []*Link)
	Traverse(int64, *Predicate, bool, bool, int64) (bool, []*Vertex)
	TraverseCount(int64, *Predicate, bool) int64
}