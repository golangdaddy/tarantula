package graph

import (
	"fmt"
	"sync"
	"encoding/json"
	//
	"github.com/golangdaddy/tarantula/web/validation"
)

type Predicate struct {
	DB          *Database					`json:"-"`
	Name        string
	Edges       *Class
	History     *Class
	Properties  *Class
	Initialized bool
	In 			*Vector
	Out 		*Vector
	edgeCount	int64
	sync 		bool
	sync.RWMutex
}

func (predicate *Predicate) Init() {

	if predicate.Initialized {
		return
	}

	// edge

	var q string

	predicate.Edges = predicate.DB.addInternalClass("_e_" + predicate.Name)
	predicate.Edges.Internal = true

	predicate.Edges.AddProperty("uid", validation.StringExact(SHA3_ID_LENGTH), SHA3_ID_LENGTH, NOT_NULL{}, PRIMARY_KEY{})
	// in uid
	predicate.Edges.AddProperty("i", validation.Int64(), NOT_NULL{})
	// in class
	predicate.Edges.AddProperty("ic", validation.Int64(), NOT_NULL{})
	// out uid
	predicate.Edges.AddProperty("o", validation.Int64(), 30, NOT_NULL{})
	// out class
	predicate.Edges.AddProperty("oc", validation.Int64(), 30, NOT_NULL{})
	// current state
	predicate.Edges.AddProperty("s", validation.Bool(), NOT_NULL{})
	// timestamp
	predicate.Edges.AddProperty("m", validation.SQLTimestamp(), NOT_NULL{})
	// latest history record uid
	predicate.Edges.AddProperty("hid", validation.Int64(), NOT_NULL{})

	if !predicate.DB.Client.InsertInternalTable(predicate.Edges) {
		panic("FAILED TO INIT PREDICATE: " + predicate.Name)
	}

	q = fmt.Sprintf("CREATE INDEX index_i_%v ON %v (i, s);", predicate.Edges.Name, predicate.Edges.Table())
	predicate.DB.Client.Exec(q)

	q = fmt.Sprintf("CREATE INDEX index_o_%v ON %v (o, s);", predicate.Edges.Name, predicate.Edges.Table())
	predicate.DB.Client.Exec(q)

	// history

	predicate.History = predicate.DB.addInternalClass("_h_" + predicate.Name)
	predicate.History.Internal = true

	predicate.History.AddProperty("uid", validation.Int64(), NOT_NULL{}, PRIMARY_KEY{}, AUTO_INCREMENT{})
	// foreign key of edge
	predicate.History.AddProperty("e", validation.StringExact(SHA3_ID_LENGTH), SHA3_ID_LENGTH, NOT_NULL{})
	// state
	predicate.History.AddProperty("s", validation.Bool(), NOT_NULL{})
	// properties
	predicate.History.AddProperty("p", validation.Int(), NOT_NULL{})
	// timestamp
	predicate.History.AddProperty("m", validation.SQLTimestamp(), NOT_NULL{})

	if !predicate.DB.Client.InsertInternalTable(predicate.History) {
		panic("FAILED TO INIT PREDICATE: " + predicate.Name)
	}

	// history

	predicate.Properties = predicate.DB.addInternalClass("_p_" + predicate.Name)
	predicate.Properties.Internal = true

	// vertex uid reference
	predicate.Properties.AddProperty("h", validation.Int64(), NOT_NULL{}, PRIMARY_KEY{})
	// property key
	predicate.Properties.AddProperty("k", validation.String(1, 32), 32, NOT_NULL{}, PRIMARY_KEY{})
	// property value
	predicate.Properties.AddProperty("v", validation.Json(), NOT_NULL{})
	// timestamp
	predicate.Properties.AddProperty("m", validation.SQLTimestamp(), NOT_NULL{})

	if !predicate.DB.Client.InsertInternalTable(predicate.Properties) {
		panic("FAILED TO INIT PREDICATE: " + predicate.Name)
	}

	// mark as created

	predicate.Initialized = true
}

func (db *Database) GetPredicate(name string) *Predicate {

	db.Classes.RLock()

	predicate := db.Classes.Predicates[name]

	db.Classes.RUnlock()

	return predicate
}

func (db *Database) SetPredicate(name string, predicate *Predicate) {

	db.Classes.Lock()

	db.Classes.Predicates[name] = predicate

	db.Classes.Unlock()
}

func (db *Database) AddPredicate(name string) *Predicate {

	db.Log.Debug("ADDING PREDICATE: " + name)

	p := db.GetPredicate(name)

	if p == nil {

		p = &Predicate{
			db,
			name,
			nil,
			nil,
			nil,
			false,
			nil,	
			nil,
			0,
			false,
			sync.RWMutex{},
		}

		p.In = &Vector{
			predicate: p,
			DB: db,
			in: true,
			states:	map[int64]*StateIndex{},
			order: map[int64][]*Link{},
		}

		p.Out = &Vector{
			predicate: p,
			DB: db,
			states:	map[int64]*StateIndex{},
			order: map[int64][]*Link{},
		}

		db.SetPredicate(name, p)

		p.Init()

	}

	db.Log.Debug("ADDED PREDICATE: " + name)

	return p
}


func (predicate *Predicate) NewLink(inVertex *Vertex, outVertex *Vertex, s bool, properties ...map[string]interface{}) (bool, *Link) {

	link := &Link{
		I: inVertex,
		O: outVertex,
		State: s,
		Predicate: predicate,
	}

	if len(properties) == 1 {

		link.Props = len(properties[0])

		go predicate.Update_Properties(link.Uid, properties[0])

	}

	if !predicate.Insert_History(link) {
		return false, nil
	}

	if !predicate.Insert_Edge(link) {
		return false, nil
	}

	return true, link
}

func (predicate *Predicate) Insert_Edge(link *Link) bool {

	q := fmt.Sprintf("SELECT uid, i, ic, o, oc, s, m, hid FROM %v WHERE i = %v AND o = %v;", predicate.Edges.Table(), link.I.Uid, link.O.Uid)

	ok, edges := predicate.DB.Client.QueryEdges(predicate, q)
	if !ok {
		return false
	}

	if len(edges) == 0 {

		q := fmt.Sprintf("INSERT INTO %v (uid, i, ic, o, oc, s, hid) VALUES (?, ?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE s = ?, m = Now();", predicate.Edges.Table())

		if ok, _ := predicate.DB.Client.Exec(q, link.EdgeID(), link.I.Uid, link.I.Class.Uid, link.O.Uid, link.O.Class.Uid, link.State, link.Uid, link.State); !ok {
			return false
		}

	} else {

		q := fmt.Sprintf("UPDATE %v SET s = %v WHERE i = %v AND o = %v;", predicate.Edges.Table(), link.State, link.I.Uid, link.O.Uid)

		if ok, _ := predicate.DB.Client.Exec(q); !ok {
			return false
		}

	}

	return true
}

func (predicate *Predicate) Insert_History(link *Link) bool {

	q := fmt.Sprintf("INSERT INTO %v (e, s, p) VALUES (?, ?, ?);", predicate.History.Table())

	ok, r := predicate.DB.Client.Exec(q, link.EdgeID(), link.State, link.Props)
	if !ok {
		return false
	}

	link.Uid, _ = r.LastInsertId()

	return true
}

func (predicate *Predicate) Update_Properties(historyId int64, payload map[string]interface{}) bool {

	tableName := predicate.Properties.Table()

	for k, v := range payload {

		x, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}

		value := string(x)

		q := fmt.Sprintf("INSERT INTO %v (h, k, v) VALUES (%v, ?, ?) ON DUPLICATE KEY UPDATE k = ?, v = ?;", tableName, historyId)

		if ok, _ := predicate.DB.Client.Exec(q, k, value, k, value); !ok {
			return false
		}

	}

	return true
}
