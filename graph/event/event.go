package event

import (
	"github.com/golangdaddy/tarantula/graph"
)

type Client interface {
	NewEventState(*graph.Link) bool
	InsertVertex(*graph.Vertex) bool
	UpdateVertex(*graph.Vertex) bool
	DeleteVertex(*graph.Vertex) bool
}

