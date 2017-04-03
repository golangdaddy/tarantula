package mysql

import (
	"fmt"
	//
	"github.com/golangdaddy/tarantula/graph"
)

func (mysql *Client) QueryTable(tableName string) bool {

	q := fmt.Sprintf("SELECT table_name FROM information_schema.tables WHERE table_schema = ? AND table_name = ?;")

	ok, rows := mysql.QueryRows(q, mysql.Credentials.DatabaseName(), tableName)

	if !ok {
		return false
	} else {
		defer rows.Close()
	}

	var x int

	for rows.Next() {
		x++
	}

	return (x == 1)
}

func (mysql *Client) QueryClassList(class *graph.Class, limit, page int64) (bool, []*graph.Vertex) {

	var q string

	if limit < 0 {

		q = fmt.Sprintf("SELECT * FROM %s WHERE c = ?;", mysql.Table(graph.TABLE_VERTICES))

	} else {

		if page >= 0 {

			offset := limit * (page - 1)

			q = fmt.Sprintf("SELECT * FROM %s WHERE c = ? LIMIT %v OFFSET %v;", mysql.Table(graph.TABLE_VERTICES), limit, offset)

		} else {

			q = fmt.Sprintf("SELECT * FROM %s WHERE c = ? LIMIT %v;", mysql.Table(graph.TABLE_VERTICES), limit)

		}

	}

	return mysql.QueryVertices(q, class.Uid)
}

func (mysql *Client) QueryClassCount(class *graph.Class) int64 {

	q := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE c = %v;", mysql.Table(graph.TABLE_VERTICES), class.Uid)

	return mysql.CountRows(q)
}

func (mysql *Client) QueryInList(predicate *graph.Predicate, subject *graph.Vertex, limit int64) (bool, []*graph.Vertex) {

	var q string

	if limit <= 0 {

		q = fmt.Sprintf("SELECT a.* FROM %s a JOIN %s b ON b.o = a.uid WHERE b.i = %v AND b.s = true ORDER BY b.m DESC;", mysql.Table(graph.TABLE_VERTICES), predicate.Edges.Table(), subject.Uid)

	} else {

		q = fmt.Sprintf("SELECT a.* FROM %s a JOIN %s b ON b.o = a.uid WHERE b.i = %v AND b.s = true ORDER BY b.m DESC LIMIT %v;", mysql.Table(graph.TABLE_VERTICES), predicate.Edges.Table(), subject.Uid, limit)

	}

	return mysql.QueryVertices(q)
}

func (mysql *Client) QueryInCount(predicate *graph.Predicate, subject *graph.Vertex) int64 {

	q := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE i = %v AND s = true;", predicate.Edges.Table(), subject.Uid)

	return mysql.CountRows(q)
}

func (mysql *Client) QueryInClassList(class *graph.Class, predicate *graph.Predicate, subject *graph.Vertex, limit int64, page int64) (bool, []*graph.Vertex) {

	var q string

	if limit <= 0 {

		q = fmt.Sprintf("SELECT a.* FROM %s a JOIN %s b ON b.o = a.uid WHERE b.i = %v AND b.s = true AND a.c = %v ORDER BY b.m DESC;", mysql.Table(graph.TABLE_VERTICES), predicate.Edges.Table(), subject.Uid, class.Uid)

	} else {

		if page < 0 {

			q = fmt.Sprintf("SELECT a.* FROM %s a JOIN %s b ON b.o = a.uid WHERE b.i = %v AND b.s = true AND a.c = %v ORDER BY b.m DESC LIMIT %v;", mysql.Table(graph.TABLE_VERTICES), predicate.Edges.Table(), subject.Uid, class.Uid, limit)

		} else {

			offset := limit * (page - 1)

			q = fmt.Sprintf("SELECT a.* FROM %s a JOIN %s b ON b.o = a.uid WHERE b.i = %v AND b.s = true AND a.c = %v ORDER BY b.m DESC LIMIT %v OFFSET %v;", mysql.Table(graph.TABLE_VERTICES), predicate.Edges.Table(), subject.Uid, class.Uid, limit, offset)

		}

	}

	return mysql.QueryVertices(q)
}

func (mysql *Client) QueryInClassCount(class *graph.Class, predicate *graph.Predicate, subject *graph.Vertex) int64 {

	q := fmt.Sprintf("SELECT COUNT(*) FROM %s a JOIN %s b ON b.o = a.uid WHERE b.i = %v AND b.s = true AND a.c = %v;", mysql.Table(graph.TABLE_VERTICES), predicate.Edges.Table(), subject.Uid, class.Uid)

	return mysql.CountRows(q)
}

func (mysql *Client) QueryOutList(predicate *graph.Predicate, subject *graph.Vertex, limit int64) (bool, []*graph.Vertex) {

	var q string

	if limit <= 0 {

		q = fmt.Sprintf("SELECT a.* FROM %s a JOIN %s b ON b.i = a.uid WHERE b.o = %v AND b.s = true ORDER BY b.m DESC;", mysql.Table(graph.TABLE_VERTICES), predicate.Edges.Table(), subject.Uid)

	} else {

		q = fmt.Sprintf("SELECT a.* FROM %s a JOIN %s b ON b.i = a.uid WHERE b.o = %v AND b.s = true ORDER BY b.m DESC LIMIT %v;", mysql.Table(graph.TABLE_VERTICES), predicate.Edges.Table(), subject.Uid, limit)

	}

	return mysql.QueryVertices(q)
}

func (mysql *Client) QueryOutCount(predicate *graph.Predicate, subject *graph.Vertex) int64 {

	q := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE o = %v AND s = true;", predicate.Edges.Table(), subject.Uid)

	return mysql.CountRows(q)
}

func (mysql *Client) QueryOutClassList(class *graph.Class, predicate *graph.Predicate, subject *graph.Vertex, limit int64, page int64) (bool, []*graph.Vertex) {

	var q string

	if limit <= 0 {

		q = fmt.Sprintf("SELECT a.* FROM %s a JOIN %s b ON b.i = a.uid WHERE b.o = %v AND b.s = true AND a.c = %v ORDER BY b.m DESC;", mysql.Table(graph.TABLE_VERTICES), predicate.Edges.Table(), subject.Uid, class.Uid)

	} else {

		if page < 0 {

			q = fmt.Sprintf("SELECT a.* FROM %s a JOIN %s b ON b.i = a.uid WHERE b.o = %v AND b.s = true AND a.c = %v ORDER BY b.m DESC LIMIT %v;", mysql.Table(graph.TABLE_VERTICES), predicate.Edges.Table(), subject.Uid, class.Uid, limit)
		
		} else {

			offset := limit * (page - 1)

			q = fmt.Sprintf("SELECT a.* FROM %s a JOIN %s b ON b.i = a.uid WHERE b.o = %v AND b.s = true AND a.c = %v ORDER BY b.m DESC LIMIT %v;", mysql.Table(graph.TABLE_VERTICES), predicate.Edges.Table(), subject.Uid, class.Uid, limit, offset)

		}
	}

	return mysql.QueryVertices(q)
}

func (mysql *Client) QueryOutClassCount(class *graph.Class, predicate *graph.Predicate, subject *graph.Vertex) int64 {

	q := fmt.Sprintf("SELECT COUNT(*) FROM %s a JOIN %s b ON b.i = a.uid WHERE b.o = %v AND b.s = true AND a.c = %v;", mysql.Table(graph.TABLE_VERTICES), predicate.Edges.Table(), subject.Uid, class.Uid)

	return mysql.CountRows(q)
}
