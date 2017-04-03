package graph

import (
	"sync"
)

type Classes struct {
	DB            *Database				`json:"-"`
	Index         []*Class
	UidIndex      map[int64]*Class
	NameIndex     map[string]*Class
	Predicates    map[string]*Predicate
	Relationships []*Relationship
	sync.RWMutex
}

func (classes *Classes) Add(class *Class) {

	ok, uid := class.DB.Client.QueryClassUID(class.Name); if !ok {

		ok, uid = class.DB.Client.InsertClass(class.Name); if !ok { class.DB.Log.Panic("*Classes.Add() FAILED") }

	}

	class.Uid = uid

	// insert class into classes structure

	classes.Lock()

	classes.Index = append(classes.Index, class)
	classes.UidIndex[class.Uid] = class
	classes.NameIndex[class.Name] = class

	classes.Unlock()
}

func (db *Database) newClasses() *Classes {

	return &Classes{
		db,
		[]*Class{},
		map[int64]*Class{},
		map[string]*Class{},
		map[string]*Predicate{},
		[]*Relationship{},
		sync.RWMutex{},
	}
}
