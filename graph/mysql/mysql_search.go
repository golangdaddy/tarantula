package mysql

import (
	"fmt"
	"encoding/json"
	//
	"github.com/golangdaddy/tarantula/graph"
)

func (mysql *Client) SearchClassProperties(class *graph.Class, propertyName, searchInput string) (bool, []*graph.PropertyExport) {

	log := mysql.Log

	searchInput = "%" + searchInput + "%"

	q := fmt.Sprintf(`SELECT b.uid, a.k, a.v FROM %s a JOIN %s b ON b.uid = a.vertex WHERE b.c = %v AND k = ? AND JSON_EXTRACT(v, "$") LIKE ?;`, mysql.Table(graph.TABLE_PROPERTIES), mysql.Table(graph.TABLE_VERTICES), class.Uid)

	ok, rows := mysql.QueryRows(q, propertyName, searchInput)
	if ok {
		defer rows.Close()
	} else {
		return false, nil
	}

	props := []*graph.PropertyExport{}

	for rows.Next() {

		prop := &graph.PropertyExport{}

		if err := rows.Scan(&prop.Vertex, &prop.K, &prop.V); err != nil {

			log.NewError("QueryLinks: FAILED")
			log.Error(err)
			return false, nil
		}

		props = append(props, prop)

	}

	return true, props
}


func (mysql *Client) SearchProperties(propertyName, searchInput string) (bool, []interface{}) {

	log := mysql.Log

	searchInput = "%" + searchInput + "%"

	q := fmt.Sprintf(`SELECT * FROM %v WHERE k = ? AND JSON_EXTRACT(v, "$") LIKE ?;`, mysql.Table(graph.TABLE_PROPERTIES))

	ok, rows := mysql.QueryRows(q, propertyName, searchInput)
	if ok {
		defer rows.Close()
	} else {
		return false, nil
	}

	props := []interface{}{}

	for rows.Next() {

		var key string
		var value string
		var vertex int64

		if err := rows.Scan(&vertex, &key, &value); err != nil {

			log.NewError("QueryLinks: FAILED")
			log.Error(err)
			return false, nil
		}

		x := new(interface{})

		err := json.Unmarshal([]byte(value), &x)
		if log.Error(err) {
			return false, nil
		}

		props = append(props, *x)

	}

	return true, props
}