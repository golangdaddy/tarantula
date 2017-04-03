package graph

import "fmt"

func (db *Database) Table(id string) string {
	
	return fmt.Sprintf("%s.%s", db.Name, id)
}