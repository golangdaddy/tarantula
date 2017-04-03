package mysql

import (
	"fmt"
	"strings"
	//
	"github.com/golangdaddy/tarantula/graph"
)

func (mysql *Client) InsertInternalTable(class *graph.Class) bool {

	a := []string{
		fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v (", class.Table()),
	}

	keys := []string{}

	props := len(class.Properties)

	//class.DB.log.DebugJSON(class.Properties)

	for i, property := range class.Properties {

		for constraint, _ := range property.Constraints {

			switch c := constraint.(type) {

				case graph.FOREIGN_KEY:

					class.FKs[property.Key] = c

				case graph.PRIMARY_KEY:

					class.PKs = append(class.PKs, graph.PRIMARY_KEY{property.Key})

			}
		}

		if property.Constraints[graph.PRIMARY_KEY{}] {
			keys = append(keys, property.Key)
		}

		a = append(a, property.SQL_Datatype())

		x := (props - i) - 1

		if x > 0 {
			a = append(a, ", ")
		}
	}

	if len(class.PKs) > 0 {

		columns := []string{}

		for _, primaryKey := range class.PKs {

			//class.DB.log.Debug(fmt.Sprintf("ADDING PRIMARY KEY %v TO CLASS %v!", primaryKey.Value(), class.Name))

			columns = append(columns, primaryKey.Value())

		}

		a = append(a, fmt.Sprintf(", PRIMARY KEY(%v)", strings.Join(columns, ", ")))

	}

	for key, constraint := range class.FKs {

		a = append(a, fmt.Sprintf(", FOREIGN KEY(%v) REFERENCES %v(uid) ON DELETE CASCADE", key, constraint.Value()))

	}

	a = append(a, ");")

	query := strings.Join(a, "")

	ok, _ := mysql.Exec(query)

	return ok
}
