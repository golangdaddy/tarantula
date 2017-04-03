package graph

type EventClient interface {
	NewEventState(*Link) bool
	InsertVertex(*Vertex) bool
	UpdateVertex(*Vertex) bool
	DeleteVertex(*Vertex) bool
}

