package graph

var global *Database

func DB() *Database {

	return global
}