package graph

import (
	"sync"
)

type VertexIndex struct {
	table map[int64]*Vertex
	sync.RWMutex
}

func (vi VertexIndex) New() *VertexIndex {

	return &VertexIndex{
		map[int64]*Vertex{},
		sync.RWMutex{},
	}
}

func (db *Database) GetVertex(uid int64) *Vertex {

	db.vertexIndex.RLock()

	vtx := db.vertexIndex.table[uid]

	db.vertexIndex.RUnlock()

	if vtx != nil { db.SetVertex(vtx) }

	return vtx
}

func (db *Database) SetVertex(vtx *Vertex) {

	db.vertexIndex.Lock()

	db.vertexIndex.table[vtx.Uid] = vtx

	db.vertexIndex.Unlock()
}
