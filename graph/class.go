package graph

import (
	"fmt"
	"sync"
	//
	"github.com/golangdaddy/tarantula/router/common"
	"github.com/golangdaddy/tarantula/web/validation"
)

type Class struct {
	// root schema
	DB *Database						`json:"-"`
	// deterministic id
	Uid int64
	// class name
	Name string
	// primary keys
	PKs []Constraint
	// foreign keys
	FKs map[string]Constraint
	// class schema
	Properties  Properties
	PropertyMap map[string]*Property
	// gf node
//	Node *gf.Node
	// gf node
//	InstanceNode *gf.Node
	// is internal class
	IsUser   bool
	IsAnchor bool
	Internal bool
	sync.RWMutex
}

func (db *Database) newClass(name string) *Class {

	return &Class{
		db,
		0,
		name,
		[]Constraint{},
		map[string]Constraint{},
		Properties{},
		map[string]*Property{},
//		db.Root.Add(name, "$class"),
//		nil,
		false,
		false,
		false,
		sync.RWMutex{},
	}
}

// adds a class marked as internal
func (db *Database) addInternalClass(name string) *Class {

	class := db.newClass(name)

	class.Internal = true

	return class
}

// adds a custom class to the schema
func (db *Database) AddClass(name string) *Class {

	class := db.ClassNameIndex(name)
	if class != nil {
		e := fmt.Sprintf("CLASS %s ALREADY EXISTS IN SCHEMA!", name)
		panic(e)
	}

	class = db.newClass(name)

	db.Classes.Add(class)

	return class
}

func (class *Class) Table() string {

	return class.DB.Name + "." + class.Name
}

func (class *Class) NewVertex(x string, data map[string]interface{}) *Vertex {

	if data == nil {
		data = map[string]interface{}{}
	}

	vertex := &Vertex{
		class.DB,
		class,
		0,
		x,
		data,
		sync.RWMutex{},
	}

	return vertex
}

func (class *Class) HasProperty(key string) (bool, *Property) {

	class.RLock()

		p := class.PropertyMap[key]

	class.RUnlock()

	return (p != nil), p
}

func (class *Class) NewProperty(key string, value *validation.Config) *Property {

	return &Property{class.DB, 0, key, value, 0, map[Constraint]bool{}, sync.RWMutex{}}
}

func (class *Class) Column(i int) (bool, *Property) {

	p := class.Properties[i]

	return (p != nil), p
}

func (class *Class) InsertColumns() Properties {

	props := Properties{&Property{}}

	props = append(props, class.Properties[1:]...)

	return props
}

func (class *Class) Validation() *common.Payload {

	schema := common.Payload{}

	for _, property := range class.Properties {

		schema[property.Key] = property.Value

	}

	return &schema
}

func (class *Class) AddProperty(key string, value *validation.Config, options ...interface{}) {

	var propExists bool

	class.RLock()

	for _, prop := range class.Properties {

		if key == prop.Key {
			propExists = true
		}

	}

	class.RUnlock()

	if propExists {

		e := fmt.Sprintf("PROPERTY WITH KEY %s ALREADY EXISTS IN CLASS %s!", key, class.Name)

		class.DB.Log.Panic(e)

	}

	// property = Null, Key, Value, Max
	property := class.NewProperty(key, value)

	// apply options
	property.Options(options)

	class.Lock()

	class.Properties = append(class.Properties, property)

	class.Unlock()

	if class.Internal { property.SQL_Datatype() }
}

func (class *Class) Link(predicate string, classes ...*Class) {

	if len(classes) == 0 {
		class.DB.Log.Panic("THIS METHOD REQUIRES AT LEAST 1 CLASS ARGUMENT")
	}

	for _, counterparty := range classes {

		relationship := &Relationship{
			In:        class,
			Out:       counterparty,
			Predicate: predicate,
		}

		// create predicate struct and create predicate table
		class.DB.AddPredicate(predicate)

		// register relationship
		class.DB.AddRelationship(relationship)

	}

}

func (db *Database) ClassNameIndex(name string) *Class {

	db.Classes.RLock()

	class := db.Classes.NameIndex[name]

	db.Classes.RUnlock()

	return class
}

func (db *Database) ClassUidIndex(uid int64) *Class {

	db.Classes.RLock()

	class := db.Classes.UidIndex[uid]

	db.Classes.RUnlock()

	//if class == nil { panic("CLASS NOT FOUND: "+fmt.Sprintf("%v", uid)) }

	return class
}
