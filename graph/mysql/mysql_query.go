package mysql

import (
	"fmt"
	"time"
	"strings"
	"database/sql"
	"encoding/json"
	//
	"github.com/golangdaddy/tarantula/graph"
)

func (mysql *Client) QueryRows(q string, params ...interface{}) (bool, *sql.Rows) {

	mysql.Log.Debug(q)

	results, err := mysql.client.Query(q, params...)

	if mysql.Log.Error(err) { mysql.Log.NewError(q) }

	return (err == nil), results
}

func (mysql *Client) CountRows(q string, args ...interface{}) int64 {

	mysql.Log.Debug(q)

	var count int64

	rows := mysql.client.QueryRow(q, args...)

	rows.Scan(&count)

	return count
}

func (mysql *Client) QueryClassUID(name string) (bool, int64) {

	log := mysql.Log

	q := fmt.Sprintf("SELECT uid FROM %v WHERE name = ?;", mysql.Table(graph.TABLE_CLASSES))

	ok, rows := mysql.QueryRows(q, name)
	if ok {
		defer rows.Close()
	} else {
		return false, 0
	}

	for rows.Next() {

		var uid int64 = 0

		if err := rows.Scan(&uid); err != nil {

			log.NewError("GetClassUID: FAILED")
			log.Error(err)
			return false, 0
		}

		return true, uid

	}

	return false, 0
}

func (mysql *Client) QueryUser(q string, args ...interface{}) (bool, *graph.User) {

	log := mysql.Log

	ok, rows := mysql.QueryRows(q, args...)
	if ok {
		defer rows.Close()
	} else {
		return false, nil
	}

	for rows.Next() {

		user := &graph.User{
			DB: mysql.DB,
		}

		var vertexUid int64

		if err := rows.Scan(&user.Uid, &user.Name, &user.PasswordHash, &vertexUid); err != nil {

			log.NewError("QueryUser: FAILED")
			log.Error(err)
			return false, nil
		}

		ok, user.Vertex = mysql.QueryVertex(vertexUid)
		if !ok {
			return false, nil
		}

		return true, user

	}

	return false, nil
}

func (mysql *Client) QuerySession(q string, args ...interface{}) (bool, *graph.Session) {

	log := mysql.Log

	ok, rows := mysql.QueryRows(q, args...)
	if ok {
		defer rows.Close()
	} else {
		return false, nil
	}

	for rows.Next() {

		sesh := &graph.Session{
			DB: mysql.DB,
		}

		if err := rows.Scan(&sesh.Uid, &sesh.TokenHash, &sesh.Timestamp); err != nil {

			log.NewError("QueryUser: FAILED")
			log.Error(err)
			return false, nil
		}

		return true, sesh

	}

	return false, nil
}

func (mysql *Client) QueryAllEdges(predicate *graph.Predicate) (bool, []*graph.Link) {

	q := fmt.Sprintf("SELECT * FROM %s;", predicate.Edges.Table())

	return mysql.QueryEdges(predicate, q)
}

func (mysql *Client) CountAllEdges(predicate *graph.Predicate) int64 {

	q := fmt.Sprintf("SELECT COUNT(*) FROM %s;", predicate.Edges.Table())

	return mysql.CountRows(q)
}

func (mysql *Client) QueryEdges(predicate *graph.Predicate, q string, args ...interface{}) (bool, []*graph.Link) {

	log := mysql.Log

	ok, rows := mysql.QueryRows(q, args...)
	if ok {
		defer rows.Close()
	} else {
		return false, nil
	}

	links := []*graph.Link{}

	for rows.Next() {

		link := &graph.Link{
			Predicate: predicate,
		}

		var timeSerial string

		err := rows.Scan(&link.Id, &link.Iuid, &link.Iclass, &link.Ouid, &link.Oclass, &link.State, &timeSerial, &link.H); if err != nil {

			log.NewError("QueryEdges: FAILED")
			log.Error(err)
			return false, nil
		}

		link.M, err = time.Parse("2006-01-02 15:04:05", timeSerial); if mysql.Log.Error(err) { return false, nil }

		links = append(links, link)

	}

	return true, links
}

func (mysql *Client) QueryEdge(predicate *graph.Predicate, q string, args ...interface{}) (bool, *graph.Link) {

	log := mysql.Log

	ok, rows := mysql.QueryRows(q, args...)
	if ok {
		defer rows.Close()
	} else {
		return false, nil
	}

	for rows.Next() {

		link := &graph.Link{
			Predicate: predicate,
		}

		var timeSerial string

		err := rows.Scan(&link.Id, &link.Iuid, &link.Iclass, &link.Ouid, &link.Oclass, &link.State, &timeSerial, &link.H); if err != nil {

			log.NewError("QueryEdge: FAILED")
			log.Error(err)
			return false, nil
		}

		link.M, err = time.Parse("2006-01-02 15:04:05", timeSerial); if mysql.Log.Error(err) { return false, nil }		

		return true, link
	}

	return false, nil
}

func (mysql *Client) QueryLinks(q string, args ...interface{}) (bool, []*graph.Link) {

	log := mysql.Log

	ok, rows := mysql.QueryRows(q, args...)
	if ok {
		defer rows.Close()
	} else {
		return false, nil
	}

	links := []*graph.Link{}

	for rows.Next() {

		link := &graph.Link{}

		if err := rows.Scan(&link.Id, &link.State, &link.M); err != nil {

			log.NewError("QueryLinks: FAILED")
			log.Error(err)
			return false, nil
		}

		links = append(links, link)

	}

	return true, links
}

func (mysql *Client) QueryLink(q string, args ...interface{}) (bool, *graph.Link) {

	log := mysql.Log

	ok, rows := mysql.QueryRows(q, args...)
	if ok {
		defer rows.Close()
	} else {
		return false, nil
	}

	for rows.Next() {

		link := &graph.Link{}

		if err := rows.Scan(&link.Id, &link.State, &link.M); err != nil {

			log.NewError("QueryLink: FAILED")
			log.Error(err)
			return false, nil
		}

		return true, link

	}

	return false, nil
}

func (mysql *Client) QueryProperties(vertex *graph.Vertex, args ...string) (bool, map[string]interface{}) {

	log := mysql.Log

	var q string
	vals := []interface{}{}

	if len(args) == 0 {

		q = fmt.Sprintf("SELECT k, v FROM %v WHERE vertex = ?;", mysql.Table(graph.TABLE_PROPERTIES))

		vals = append(vals, vertex.Uid)

	} else {

		list := []string{}

		for _, item := range args {

			vals = append(vals, item)
			list = append(list, "k = ?")

		}

		q = fmt.Sprintf("SELECT k, v FROM %v WHERE vertex = %v AND (%v);", mysql.Table(graph.TABLE_PROPERTIES), vertex.Uid, strings.Join(list, " OR "))

	}

	ok, rows := mysql.QueryRows(q, vals...)
	if ok {
		defer rows.Close()
	} else {
		return false, nil
	}

	props := map[string]interface{}{}

	for rows.Next() {

		var key string
		var value string

		if err := rows.Scan(&key, &value); err != nil {

			log.NewError("QueryLinks: FAILED")
			log.Error(err)
			return false, nil
		}

		x := new(interface{})

		err := json.Unmarshal([]byte(value), &x)
		if log.Error(err) {
			return false, nil
		}

		props[key] = *x

	}

	return true, props
}

func (mysql *Client) QueryProperty(q string) (bool, interface{}) {

	log := mysql.Log

	log.Debug(q)

	ok, rows := mysql.QueryRows(q)
	if ok {
		defer rows.Close()
	} else {
		return false, nil
	}

	for rows.Next() {

		var key string
		var value string

		if err := rows.Scan(&key, &value); err != nil {

			log.NewError("QueryLinks: FAILED")
			log.Error(err)
			return false, nil
		}

		x := new(interface{})

		err := json.Unmarshal([]byte(value), &x)
		if log.Error(err) {
			return false, nil
		}

		return true, *x

	}

	log.NewError("NO PROPERTY FOUND: " + q)

	return false, nil
}

func (mysql *Client) QueryState(in, out *graph.Vertex, predicate *graph.Predicate, state bool) bool {

	link := &graph.Link{
		I: in,
		O: out,
	}

	edgeID := link.EdgeID()

	q := fmt.Sprintf("SELECT * FROM %v WHERE id = %v;", predicate.Edges.Table(), edgeID)

	ok, edge := mysql.QueryEdge(predicate, q)
	if !ok {
		return false
	}

	return edge.State
}

