package mysql

import (
	"fmt"
	"strings"
	"database/sql"
	//
	"github.com/golangdaddy/tarantula/graph"
)

func (mysql *Client) scanVertex(fn string, rows *sql.Rows) *graph.Vertex {

	vertex := &graph.Vertex{
		DB:   graph.DB(),
		Data: map[string]interface{}{},
	}

	var classUid int64

	if err := rows.Scan(&vertex.Uid, &classUid, &vertex.X); err != nil {

		mysql.Log.NewError(fn + ": FAILED")
		mysql.Log.Error(err)

		return nil
	}

	vertex.Class = vertex.DB.ClassUidIndex(classUid)

	mysql.DB.SetVertex(vertex)

	return vertex
}

func (mysql *Client) QueryVertices(q string, args ...interface{}) (bool, []*graph.Vertex) {
	fn := "QueryVertices"

	ok, rows := mysql.QueryRows(q, args...)
	if !ok {

		mysql.Log.NewError("FAILED TO QUERY VERTICES")

		return false, nil
	}

	defer rows.Close()

	vertices := []*graph.Vertex{}

	for rows.Next() {

		vertices = append(vertices, mysql.scanVertex(fn, rows))

	}

	return true, vertices
}

func (mysql *Client) QueryVertex(uid int64) (bool, *graph.Vertex) {
	fn := "QueryVertex"

	vtx := mysql.DB.GetVertex(uid); if vtx != nil {
		return true, vtx
	}

	q := fmt.Sprintf("SELECT * FROM %v WHERE uid = ?;", mysql.Table(graph.TABLE_VERTICES))

	ok, rows := mysql.QueryRows(q, uid)
	if !ok {
		return false, nil
	}

	defer rows.Close()

	for rows.Next() {

		return true, mysql.scanVertex(fn, rows)

	}

	return false, nil
}

func (mysql *Client) QueryVertexByX(x string, classes ...*graph.Class) (bool, *graph.Vertex) {
	fn := "QueryVertexByX"

	var q string

	switch len(classes) {

		case 0:

			q = fmt.Sprintf("SELECT * FROM %v WHERE X = ?;", mysql.Table(graph.TABLE_VERTICES))

		default:

			a := []string{}
			for _, class := range classes { a = append(a, fmt.Sprintf("c = %d", class.Uid)) }

			q = fmt.Sprintf("SELECT * FROM %v WHERE X = ? AND %s;", mysql.Table(graph.TABLE_VERTICES), strings.Join(a, " OR "))

	}

	ok, rows := mysql.QueryRows(q, x)
	if !ok {
		return false, nil
	}

	defer rows.Close()

	for rows.Next() {

		return true, mysql.scanVertex(fn, rows)

	}

	return false, nil
}

