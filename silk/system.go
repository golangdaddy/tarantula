package silk

import (
	"sync"
	//
	"github.com/golangdaddy/tarantula/graph"
	"github.com/golangdaddy/tarantula/graph/mysql"
	"github.com/golangdaddy/tarantula/log"
	"github.com/golangdaddy/tarantula/log/gcp"
	"github.com/golangdaddy/tarantula/router/common"
	"github.com/golangdaddy/tarantula/router/standard"
	"github.com/golangdaddy/tarantula/graph/query"
)

type System struct {
	Log logging.Logger
	DB *graph.Database
	// Random ID for this execution instance
	InstanceID string
	Classes map[string]*SystemClass
	DBClient graph.Client
	QC *query.Client
	Root *common.Node
	sync.RWMutex
}

func NewSystem(subdomain, domain string, credentials graph.Credentials) (*System, interface{}) {

	version := "silk"

	log := logs.NewClient(credentials.ProjectID()).NewLogger(false, version)

	dbClient := mysql.NewClient(credentials.(*mysql.Credentials))

	db := graph.NewDatabase(log, credentials.DatabaseName(), dbClient)

	rootNode, router := router.NewRouter(log, subdomain)

	system := &System{
		logging.Logger(log),
		db,
		random(),
		map[string]*SystemClass{},
		dbClient,
		query.NewClient(db),
		rootNode,
		sync.RWMutex{},
	}

	return system, router
}

func (system *System) Class(name string) *SystemClass {

	system.RLock()

		sysClass := system.Classes[name]

	system.RUnlock()

	return sysClass
}

func (system *System) AddClass(key string) *SystemClass {

	class := system.DB.AddClass(key)

	classNode := system.Root.Add(class.Name, "$class")

	sysClass := &SystemClass{
		system,
		system.DB,
		class,
		classNode,
		classNode.Param(system.Subject(), "$subject"),
	}

	system.Lock()

		system.Classes[class.Name] = sysClass

	system.Unlock()

	return sysClass
}
