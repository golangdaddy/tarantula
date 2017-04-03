package graph

import (
	"fmt"
	//
	"github.com/golangdaddy/tarantula/web/validation"
)

type InternalClasses struct {
	Constraints Constraints
	users       *Class
	sessions    *Class
	classes     *Class
	vertices    *Class
	properties  *Class
}

func (db *Database) newInternalClasses() *InternalClasses {

	log := db.Log

	internal := &InternalClasses{
		Constraints{
			PRIMARY_KEY{},
			FOREIGN_KEY{},
			NOT_NULL{},
			DEFAULT{},
			UNIQUE{},
			CHECK{},
		},
		db.addInternalClass(TABLE_USERS),
		db.addInternalClass(TABLE_SESSIONS),
		db.addInternalClass(TABLE_CLASSES),
		db.addInternalClass(TABLE_VERTICES),
		db.addInternalClass(TABLE_PROPERTIES),
	}

	var q string

	// global users table

	// primary key
	internal.users.AddProperty("uid", validation.Int64(), NOT_NULL{}, PRIMARY_KEY{}, AUTO_INCREMENT{})
	// user vertex uid
	internal.users.AddProperty("name", validation.Username(3, 16), 16, NOT_NULL{})
	internal.users.AddProperty("password", validation.PasswordWeak(), 128, NOT_NULL{})
	internal.users.AddProperty("vertex", validation.Int64(), 16, NOT_NULL{})

	if !db.Client.InsertInternalTable(internal.users) {
		log.Panic("FAILED TO INSERT CLASS: " + internal.users.Name)
	}
	q = fmt.Sprintf("CREATE INDEX usersIndex ON %s (uid, name);", db.Table(TABLE_USERS))
	db.Client.Exec(q)

	// global users sessions

	// primary key
	internal.sessions.AddProperty("uid", validation.Int64(), NOT_NULL{}, PRIMARY_KEY{}, AUTO_INCREMENT{})
	// class name
	internal.sessions.AddProperty("user", validation.Int64(), 16, NOT_NULL{})
	internal.sessions.AddProperty("tokenHash", validation.StringExact(64), 64, NOT_NULL{})
	internal.sessions.AddProperty("timestamp", validation.SQLTimestamp(), NOT_NULL{})

	if !db.Client.InsertInternalTable(internal.sessions) {
		log.Panic("FAILED TO INSERT CLASS: " + internal.sessions.Name)
	}
	q = fmt.Sprintf("CREATE INDEX sessionsIndex ON %s (uid, user);", db.Table(TABLE_SESSIONS))
	db.Client.Exec(q)

	// global classes table

	// primary key
	internal.classes.AddProperty("uid", validation.Int64(), NOT_NULL{}, PRIMARY_KEY{}, AUTO_INCREMENT{})
	// class name
	internal.classes.AddProperty("name", validation.String(1, 16), 16, UNIQUE{}, NOT_NULL{})

	if !db.Client.InsertInternalTable(internal.classes) {
		log.Panic("FAILED TO INSERT CLASS: " + internal.classes.Name)
	}
	q = fmt.Sprintf("CREATE INDEX classesIndex ON %s (uid, name);", db.Table(TABLE_CLASSES))
	db.Client.Exec(q)

	// global vertex table

	// vertex uid
	internal.vertices.AddProperty("uid", validation.Int64(), NOT_NULL{}, PRIMARY_KEY{}, AUTO_INCREMENT{})
	// vertex class
	internal.vertices.AddProperty("c", validation.Int64(), NOT_NULL{})
	// vertex arbitrary reference
	internal.vertices.AddProperty("x", validation.String(1, 64), 64, NOT_NULL{})

	if !db.Client.InsertInternalTable(internal.vertices) {
		log.Panic("FAILED TO INSERT CLASS: " + internal.vertices.Name)
	}
	q = fmt.Sprintf("CREATE INDEX verticesIndex ON %s (uid, c, x);", db.Table(TABLE_VERTICES))
	db.Client.Exec(q)

	// global properties table

	// vertex uid reference
	internal.properties.AddProperty("vertex", validation.Int64(), NOT_NULL{}, PRIMARY_KEY{})
	// property key
	internal.properties.AddProperty("k", validation.String(1, 32), 32, NOT_NULL{}, PRIMARY_KEY{})
	// property value
	internal.properties.AddProperty("v", validation.Json(), NOT_NULL{})

	if !db.Client.InsertInternalTable(internal.properties) {
		log.Panic("FAILED TO INSERT CLASS: " + internal.properties.Name)
	}
	q = fmt.Sprintf("CREATE INDEX propertiesIndex ON %s (vertex, k);", db.Table(TABLE_PROPERTIES))
	db.Client.Exec(q)

	return internal
}
