package graph

import (
	"fmt"
	"time"
	"sync"
	//
)

type Link struct {
	Predicate *Predicate 				`json:"-"`
	// history uid
	Uid int64
	// edge id
	Id string
	// in
	I      *Vertex
	Iuid   int64
	Iclass int64
	// out
	O      *Vertex
	Ouid   int64
	Oclass int64
	// state
	State bool
	// update timestamp
	M time.Time
	// number of props
	H int64
	//
	Props int
	//
	sync.RWMutex
}

func (link *Link) EdgeID() string {

	input := fmt.Sprintf("%v * %v", link.I, link.O)

	return iD(input)
}

func (link *Link) HistoryID() string {

	input := fmt.Sprintf("%v * %v * %v", link.EdgeID(), link.State, link.M)

	return iD(input)
}

func (link *Link) LoadVertex(phase bool) bool {

	if phase {

		ok, vtx := link.Predicate.DB.Client.QueryVertex(link.Iuid); if !ok {
			
			link.Predicate.DB.Log.NewErrorf("FAILED TO LOAD VERTEX INTO EDGE: %v", link.Iuid)
			link.Predicate.DB.Log.ErrorJSON(link)

			return false
		}

		link.I = vtx

	} else {

		ok, vtx := link.Predicate.DB.Client.QueryVertex(link.Ouid); if !ok {

			link.Predicate.DB.Log.NewErrorf("FAILED TO LOAD VERTEX INTO EDGE: %v", link.Ouid)
			link.Predicate.DB.Log.ErrorJSON(link)
			
			return false
		}

		link.O = vtx

	}

	return true
}

func (link *Link) ToIndex() {

	if link.Predicate == nil {

		panic("*Link.ToIndex: LINK PREDICATE IS NIL")
	}

	link.LoadVertex(true)
	link.LoadVertex(false)

	//link.Predicate.DB.Log.Debug("*Link.ToIndex")

	link.Predicate.In.Add(link)
	link.Predicate.Out.Add(link)

}