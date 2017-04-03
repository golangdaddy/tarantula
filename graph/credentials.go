package graph

type Credentials interface {
	ProjectID() string
	ServiceName() string
	DatabaseName() string
}