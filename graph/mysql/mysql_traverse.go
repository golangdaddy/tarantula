package mysql

import (
	"fmt"
	//
	"github.com/golangdaddy/tarantula/graph"
)

func (mysql *Client) TraverseEdges(in int64, predicate *graph.Predicate, phase bool, limit int64) (bool, []*graph.Link) {

	io := "i"
	if !phase {
		io = "o"
	}

	q := fmt.Sprintf("SELECT * FROM %v WHERE %s = %v AND s = true;", predicate.Edges.Table(), io, in)

	return mysql.QueryEdges(predicate, q)
}

func (mysql *Client) Traverse(in int64, predicate *graph.Predicate, phase, export bool, limit int64) (bool, []*graph.Vertex) {

	ok, edges := mysql.TraverseEdges(in, predicate, phase, limit); if !ok { return false, nil }

	results := []*graph.Vertex{}

	for i, link := range edges {

		if int64(i) == limit { break }

		if phase {

			if ok, vtx := mysql.QueryVertex(link.Ouid); ok { results = append(results, vtx) }

		} else {

			if ok, vtx := mysql.QueryVertex(link.Iuid); ok { results = append(results, vtx) }

		}

	}

	predicate.DB.ExportVertexList(results)

	return true, results
}

func (mysql *Client) TraverseCount(in int64, predicate *graph.Predicate, phase bool) int64 {

	io := "i"
	if !phase {
		io = "o"
	}

	q := fmt.Sprintf("SELECT COUNT(*) FROM %v WHERE %s = ? AND s = true;", predicate.Edges.Table(), io)

	return mysql.CountRows(q, in)
}
