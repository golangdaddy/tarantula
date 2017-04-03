package graph

import (
	"fmt"
	"time"
	//
)

func (predicate *Predicate) Sync() bool {

	predicate.RLock()

		state := predicate.sync

	predicate.RUnlock()

	// make sure predicate is in sync and cache is turned on

	state = state && predicate.DB.MemGraph

	if state { predicate.DB.Log.Debug("CACHE HIT: "+predicate.Name) }

	return state
}

func (predicate *Predicate) PredicateService() {

	for predicate.DB.MemGraph {

		time.Sleep(time.Second)

		currentRowCount := predicate.DB.Client.CountAllEdges(predicate)

		predicate.DB.Log.Debug(fmt.Sprintf("COUNTED %v ROWS FROM %s", currentRowCount, predicate.Name))

		if currentRowCount != predicate.edgeCount {

			if !predicate.SyncEdges(predicate.edgeCount) { continue }

			predicate.edgeCount = currentRowCount

			predicate.Lock()

				predicate.sync = true

			predicate.Unlock()

		}

		predicate.DB.Log.Debugf("PREDICATE %s IN - %v", predicate.Name, predicate.In.Count())
		predicate.DB.Log.Debugf("PREDICATE %s OUT - %v", predicate.Name, predicate.Out.Count())

		time.Sleep(time.Minute)
	}

}

func (predicate *Predicate) SyncEdges(offset int64) bool {

	predicate.DB.Log.Debug("SYNCING EDGES: "+predicate.Name)

	ok, edges := predicate.DB.Client.QueryAllEdges(predicate); if !ok { return false }

	for _, link := range edges {

		if !link.LoadVertex(true) { continue }
		if !link.LoadVertex(false) { continue }

		link.ToIndex()

	}

	return true
}