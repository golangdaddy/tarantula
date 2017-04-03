package mysql

import (
	"fmt"
	"database/sql"
	"encoding/json"
	//
	"github.com/golangdaddy/tarantula/graph"
)

func (mysql *Client) Exec(q string, params ...interface{}) (bool, sql.Result) {

	mysql.Log.Debug(q)

	result, err := mysql.client.Exec(q, params...)

	if mysql.Log.Error(err) {
		mysql.Log.NewError(q)
	}

	return (err == nil), result
}

func (mysql *Client) CreateDatabase(name string) bool {

	q := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", name)

	ok, _ := mysql.Exec(q)

	return ok
}

func (mysql *Client) InsertClass(name string) (bool, int64) {

	q := fmt.Sprintf("INSERT INTO %s (name) VALUES (?);", mysql.Table(graph.TABLE_CLASSES))

	ok, r := mysql.Exec(q, name)
	if !ok {
		mysql.Log.NewError("FAILED TO INSERT CLASS: " + name)
		return false, 0
	}

	uid, _ := r.LastInsertId()

	return true, uid
}

func (mysql *Client) InsertVertex(class *graph.Class, X string) (bool, int64) {

	q := fmt.Sprintf("INSERT INTO %v (c, x) VALUES (?, ?);", mysql.Table(graph.TABLE_VERTICES))

	ok, r := mysql.Exec(q, class.Uid, X)
	if !ok {
		return false, 0
	}

	uid, _ := r.LastInsertId()

	return true, uid
}

func (mysql *Client) InsertProperty(vertex *graph.Vertex, key string, value interface{}) bool {

	v, err := json.Marshal(value)
	if mysql.Log.Error(err) {
		return false
	}

	q := fmt.Sprintf("INSERT INTO %v VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE k = ?, v = ?;", vertex.DB.Table(graph.TABLE_PROPERTIES))

	if ok, _ := mysql.Exec(q, vertex.Uid, key, v, key, v); !ok {
		return false
	}

	return true
}

func (mysql *Client) DeleteVertex(uid int64) bool {

	q := fmt.Sprintf("DELETE FROM %v WHERE uid = ?;", mysql.Table(graph.TABLE_VERTICES))

	ok, _ := mysql.Exec(q, uid)
	if !ok {
		return false
	}

	return true
}
