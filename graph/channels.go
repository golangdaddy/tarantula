package graph

import ()

type ChannelConfig struct {
	Channel chan interface{}
	Listeners []*EventClient
}

func newChannelConfig() *ChannelConfig {

	return &ChannelConfig{
		Channel: make(chan interface{}, 999),
		Listeners: []*EventClient{},
	}
}

type Channels struct {
	EdgeState *ChannelConfig
	InsertVertex *ChannelConfig
	DeleteVertex *ChannelConfig
	UpdateVertex *ChannelConfig
}

// add listeners to channels

func (c *Channels) ListenEdgeState(listener *EventClient) {

	if c.EdgeState == nil { c.EdgeState = newChannelConfig() }

	c.EdgeState.Listeners = append(c.EdgeState.Listeners, listener)

}

func (c *Channels) ListenInsertVertex(listener *EventClient) {

	if c.InsertVertex == nil { c.InsertVertex = newChannelConfig() }

	c.InsertVertex.Listeners = append(c.InsertVertex.Listeners, listener)

}

func (c *Channels) ListenDeleteVertex(listener *EventClient) {

	if c.DeleteVertex == nil { c.DeleteVertex = newChannelConfig() }

	c.DeleteVertex.Listeners = append(c.DeleteVertex.Listeners, listener)

}

func (c *Channels) ListenUpdateVertex(listener *EventClient) {

	if c.UpdateVertex == nil { c.UpdateVertex = newChannelConfig() }

	c.UpdateVertex.Listeners = append(c.UpdateVertex.Listeners, listener)

}

// edges

func (chans Channels) newEdgeState(link *Link) {

	if chans.EdgeState != nil { chans.EdgeState.Channel <- link }
}

// vertices

func (chans Channels) insertVertex(vertex *Vertex) {

	if chans.InsertVertex != nil { chans.InsertVertex.Channel <- vertex }
}

func (chans Channels) deleteVertex(vertex *Vertex) {

	if chans.DeleteVertex != nil { chans.DeleteVertex.Channel <- vertex }
}

func (chans Channels) updateVertex(vertex *Vertex) {

	if chans.UpdateVertex != nil { chans.UpdateVertex.Channel <- vertex }
}

