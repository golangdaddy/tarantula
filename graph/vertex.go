package graph

import (
	"sync"
	//
	"github.com/golangdaddy/tarantula/web"
)

type Vertex struct {
	DB *Database						`json:"-"`
	Class *Class						`json:"-"`
	Uid   int64
	X     string
	Data  map[string]interface{}
	sync.RWMutex
}

func (vertex *Vertex) Properties(key string) map[string]interface{} {

	vertex.RLock()

	m := vertex.Data

	vertex.RUnlock()

	return m
}

func (vertex *Vertex) GetProperty(key string) interface{} {

	vertex.RLock()

	v := vertex.Data[key]

	vertex.RUnlock()

	return v
}
 
func (vertex *Vertex) SetProperty(key string, value interface{}) {

	vertex.Lock()

	vertex.Data[key] = value

	vertex.Unlock()
}

func (vertex *Vertex) LoadProperties(args ...string) bool {

	properties := Properties{}

	for _, key := range args {

		exists, property := vertex.Class.HasProperty(key)

		if exists {

			properties = append(properties, property)

		}

	}

	ok, data := vertex.DB.Client.QueryProperties(vertex, properties.Keys()...)
	if !ok {
		vertex.DB.Log.NewError("FAILED TO QUERY PROPERTIES")
		return false
	}

	vertex.Lock()
	defer vertex.Unlock()

	if len(properties) > 0 {

		for _, prop := range properties {

			if data[prop.Key] == nil {

				vertex.Data[prop.Key] = prop.Value.Model

			} else {

				vertex.Data[prop.Key] = data[prop.Key]

			}

		}

	} else {

		vertex.Data = map[string]interface{}{}

		for _, property := range vertex.Class.Properties {

			if data[property.Key] == nil {

				vertex.Data[property.Key] = property.Value.Model

			} else {

				vertex.Data[property.Key] = data[property.Key]

			}

		}

	}

	return true
}

func (vertex *Vertex) LoadPayload(req web.RequestInterface) *web.ResponseStatus {

	payload := map[string]interface{}{}

	for _, property := range vertex.Class.Properties {

		prop := req.Param("_"+property.Key)

		if prop == nil { return req.Respond(400, "EXPECTED FIELD IS NIL: "+property.Key) }

		status, value := property.Value.BodyFunction(req, prop); if status != nil { return status }

		payload[property.Key] = value

	}

	vertex.Data = payload

	return nil
}

func (db *Database) ExportVertexList(list []*Vertex) {

	group := sync.WaitGroup{}

	group.Add(len(list))

	for _, vertex := range list {

		go func(vtx *Vertex) {

			vtx.LoadProperties()

			group.Done()

		}(vertex)

	}

	group.Wait()

}
