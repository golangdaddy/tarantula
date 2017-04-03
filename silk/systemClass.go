package silk

import (
	//
	"github.com/golangdaddy/tarantula/graph"
	"github.com/golangdaddy/tarantula/router/common"
	"github.com/golangdaddy/tarantula/web/validation"
)

type SystemClass struct {
	System *System
	DB *graph.Database
	Class *graph.Class
	Node *common.Node
	InstanceNode *common.Node
}

// defines the class to use as a user vertex
func (sysClass *SystemClass) SetAsUser() *SystemClass {

	if sysClass.DB.UserClass != nil {
		sysClass.DB.Log.Panic("TRIED TO SET MULTIPLE CLASSES AS USER")
	}

	sysClass.DB.UserClass = sysClass.Class

	sysClass.Class.IsUser = true

	return sysClass
}

// defines the class as anchor vertices
func (sysClass *SystemClass) SetAsAnchor() *SystemClass {

	sysClass.Class.IsAnchor = true

	return sysClass
}

func (sysClass *SystemClass) AddProperty(key string, vc *validation.Config) {

	sysClass.Class.AddProperty(key, vc)
}

func (sysClass *SystemClass) Link(predicate string, otherClass *SystemClass) {

	sysClass.Class.Link(predicate, otherClass.Class)
}
