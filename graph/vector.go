package graph

import (
	"fmt"
	"sync"
//	"time"
	"sort"
)

type StateIndex struct {
	predicate *Predicate 				`json:"-"`
	index map[int64]*Link
	sync.RWMutex
}

func (predicate *Predicate) newStateIndex() *StateIndex {

	return &StateIndex{
		index: map[int64]*Link{},
	}
}

func (si *StateIndex) Add(state *Link, target int64) bool {

	// init the state if not existing

	si.Lock()

		s := si.index[target]

		if s == nil {

			si.index[target] = state

			si.Unlock()

			return true

		}

	si.Unlock()

	// update the existing state

	s.Lock()

		s.State = state.State
		s.M = state.M

	s.Unlock()

	return false
}

type Vector struct {
	DB *Database
	predicate *Predicate
	in bool
	states map[int64]*StateIndex
	order map[int64][]*Link
	sync.RWMutex
}

func (vector *Vector) Add(state *Link) {

	var subject int64
	var target int64

	if vector.in {

		subject = state.Iuid
		target = state.Ouid

	} else {

		subject = state.Iuid
		target = state.Ouid

	}

	vector.Lock()

		index := vector.states[subject]

		if index == nil {

			index = vector.predicate.newStateIndex()

			vector.states[subject] = index

		}

	vector.Unlock()

	if index.Add(state, target) {

		order := vector.order[subject]

		order = append(order, state)

		// sort the order by modified time

		vector.SortOrder(subject)

	}

}

func (vector *Vector) SortOrder(uid int64) {

	order := []*Link{}

	vector.RLock()

		si := vector.states[uid]

	vector.RUnlock()

	si.RLock()

		for _, state := range si.index {

			if state.State { order = append(order, state) }

		}

	si.RUnlock()

	sort.Slice(order, func (a, b int) bool {

		return order[a].M.After(order[b].M)
	})

	vector.Lock()

		vector.order[uid] = order

	vector.Unlock()

	//vector.DB.Log.Debugf("SORTED : %v", x)
}

func (vector *Vector) Count() int64 {

	vector.RLock()

		x := int64(len(vector.order))

	vector.RUnlock()

	return x
}

func (vector *Vector) ActiveCount(uid int64) int64 {

	vector.RLock()

		x := int64(len(vector.order[uid]))

	vector.RUnlock()

	return x
}

func (vector *Vector) ActiveEdges(uid int64, limit int64) []*Link {

	vector.RLock()

		results := vector.order[uid]

	vector.RUnlock()

	vector.DB.Log.Debugf("LENGTH OF EDGE RESULTS: %v", len(results))

	return results
}

func (vector *Vector) Active(uid int64, phase bool, limit int64) []*Vertex {

	vector.RLock()

		results := make([]*Vertex, len(vector.order[uid]))

		for i, state := range vector.order[uid] {

			if int64(i) == limit { break }

			var ok bool

			if phase {

				if state.O == nil {
					ok, state.O = vector.DB.Client.QueryVertex(state.Ouid); if !ok {
						vector.DB.Log.NewError(fmt.Sprintf("FAILED TO GET VERTEX %v", state.Ouid))
						vector.DB.Log.ErrorJSON(state)
						continue
					}
				}

				results[i] = state.O

			} else {

				if state.I == nil {
					ok, state.I = vector.DB.Client.QueryVertex(state.Iuid); if !ok {
						vector.DB.Log.NewError(fmt.Sprintf("FAILED TO GET VERTEX %v", state.Iuid))
						vector.DB.Log.ErrorJSON(state)
						continue
					}
				}

				results[i] = state.I

			}

		}

	vector.RUnlock()

	vector.DB.Log.Debugf("LENGTH OF RESULTS: %v", len(results))

	return results
}

func (predicate *Predicate) Add(state *Link) {

	predicate.In.Add(state)
	predicate.Out.Add(state)

}


