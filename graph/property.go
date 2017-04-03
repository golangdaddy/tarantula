package graph

import (
	"database/sql"
	"fmt"
	"reflect"
	"sync"
	"time"
	//
	"github.com/golangdaddy/tarantula/web/validation"
	//"github.com/mitchellh/hashstructure"
)

type URL struct{}

type Properties []*Property

func (properties Properties) Keys() []string {

	keys := make([]string, len(properties))

	for i, property := range properties {

		keys[i] = property.Key

	}

	return keys
}

type PropertyExport struct {
	Vertex int64
	K string
	V interface{}
}

type Property struct {
	// root schema
	DB *Database					`json:"-"`
	// vertex uid
	Vertex int64
	// name of this field
	Key string
	// arbitrary value
	Value *validation.Config
	// max length for strings
	Max int
	// is primary key
	Constraints map[Constraint]bool `json:"-"`
	sync.RWMutex
}

func (prop *Property) Type() reflect.Type { return reflect.TypeOf(prop.Value.Model) }

func (property *Property) Options(options []interface{}) *Property {

	for _, x := range options {

		switch v := x.(type) {

		case AUTO_INCREMENT:

			property.Constraints[v] = true

		case NOT_NULL:

			property.Constraints[v] = true

		case UNIQUE:

			property.Constraints[v] = true

		case PRIMARY_KEY:

			property.Constraints[v] = true

		case FOREIGN_KEY:

			property.Constraints[v] = true

		case CHECK:

			property.Constraints[v] = true

		case DEFAULT:

			property.Constraints[v] = true

		// max length spec
		case int:

			property.Max = v

		case []string:

			for _, item := range v {

				l := len(item)

				if l > property.Max {
					property.Max = l
				}

			}

		}

	}

	return property
}

// SQL

func (property *Property) SQL_Datatype() string {

	var s string

	switch property.Value.Model.(type) {

	case time.Time:

		s += fmt.Sprintf("%v TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP", property.Key)

	case validation.URL:

		s += fmt.Sprintf("%v VARCHAR(%v)", property.Key, property.Max)

	case validation.JSON:

		s += fmt.Sprintf("%v JSON", property.Key)

	case sql.NullInt64:

		s += fmt.Sprintf("%v BIGINT", property.Key)

	case sql.NullFloat64:

		s += fmt.Sprintf("%v DOUBLE", property.Key)

	case sql.NullString:

		datatype := "VARCHAR"

		if property.Max == 1 {
			datatype = "CHARACTER"
		}

		s += fmt.Sprintf("%v %v(%v)", property.Key, datatype, property.Max)

	case int:

		s += fmt.Sprintf("%v INT", property.Key)

	case int64:

		s += fmt.Sprintf("%v BIGINT", property.Key)

	case float32:

		s += fmt.Sprintf("%v FLOAT", property.Key)

	case float64:

		s += fmt.Sprintf("%v DOUBLE", property.Key)

	case string:

		datatype := "VARCHAR"

		if property.Max == 1 || property.Max == SHA3_ID_LENGTH {
			datatype = "CHARACTER"
		}

		s += fmt.Sprintf("%v %v(%v)", property.Key, datatype, property.Max)

	case bool:

		s += fmt.Sprintf("%v TINYINT", property.Key)

	default:

		panic("(*Property).SQL_Datatype(): PROPERTY VALUE TYPE IS UNSUPPORTED: " + property.Type().String())

	}

	for constraint, _ := range property.Constraints {

		switch constraint.(type) {

		case PRIMARY_KEY:
			continue

		case FOREIGN_KEY:
			continue

		}

		s += " " + constraint.SQL()

	}

	return s
}
