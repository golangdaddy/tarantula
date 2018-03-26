package common

import	(
		"os"
		"fmt"
		"sync"
		"time"
		"strings"
		//
		"github.com/golangdaddy/tarantula/web"
		"github.com/golangdaddy/tarantula/web/validation"
		)

type Node struct {
	Config *Config
	parent *Node
	path string
	param *Node
	requestParams map[string]interface{}
	routes map[string]*Node
	methods map[string]*Handler
	module *Module
	modules []*Module
	middleware []*Middleware
	Validation *validation.Config
	Validations []*validation.Config
	sync.RWMutex
}

func (node *Node) new(path string) *Node {

	n := &Node{
		parent: node,
		Config: node.Config,
		requestParams: map[string]interface{}{},
		routes: map[string]*Node{},
		methods: map[string]*Handler{},
		modules: node.modules,
		// inherited properties
		path: path,
		Validations: node.Validations,
	}

	return n
}

func (node *Node) addHandler(method string, handler *Handler) {

	handler.Method = method
	handler.Config = node.Config
	handler.Node = node

	node.Lock()
		node.methods[method] = handler
	node.Unlock()

	node.Config.Lock()
		node.Config.Handlers = append(
			node.Config.Handlers,
			handler,
		)
	node.Config.Unlock()
}

// returns the base param map including node params
func (node *Node) RequestParams() map[string]interface{} {

  m := map[string]interface{}{}

  node.RLock()

  for k, v := range node.requestParams { m[k] = v }

  node.RUnlock()

  return m
}

// Returns the node's path string
func (node *Node) Path() string {

	return node.path
}

// Returns the node's full path string
func (node *Node) FullPath() string {

	parent := node

	path := node.path

	for {
		parent = parent.parent

		if parent == nil { break }

		path = parent.path + "/" + path
	}

	if len(path) == 0 { path = "/" }

	return path
}

// Adds a new node to the tree
func (node *Node) Add(path string, pathKeys ...string) *Node {

	path = strings.TrimSpace(strings.Replace(path, "/", "", -1))

	node.RLock()

		p := node.routes[path]

	node.RUnlock()

	if existing := p; existing != nil {
		return existing
	}

	n := node.new(path)

	node.Lock()

		node.routes[path] = n

	node.Unlock()

	if len(pathKeys) > 0 {
		n.Lock()
			for _, key := range pathKeys {
				n.requestParams[key] = path
			}
		n.Unlock()
	}

	return n
}

// Adds a new param-node
func (node *Node) Param(vc *validation.Config, keys ...string) *Node {

	if len(keys) == 0 { panic("NO KEYS SUPPLIED FOR NEW PARAMETER") }

	node.RLock()
		p := node.param
	node.RUnlock()

	if p != nil { return p }

	n := node.new(":" + keys[0])

	vc.Keys = keys

	n.Lock()
		n.Validation = vc
		n.Validations = append(n.Validations, vc)
	n.Unlock()

	node.Lock()
		node.param = n
	node.Unlock()

	return n
}

func (node *Node) newModule(function ModuleFunction, arg interface{}) *Module {

	//if node.Config.ModuleRegistry == nil { panic("Config has no ModuleRegistry setting!") }

	return &Module{
		config:					node.Config,
		function:				function,
		arg:					arg,
	}
}

// Adds a module that will be executed at the point it is added to the route
func (node *Node) Init(function ModuleFunction, arg interface{}) *Node {

	module := node.newModule(function, arg)

	node.Lock()
		node.module = module
	node.Unlock()

	return node
}

// Adds a module that will be executed upon reaching a handler
func (node *Node) Mod(function ModuleFunction, arg interface{}) *Node {

	if function == nil { panic("INVALID MODULE FUNC") }

	module := node.newModule(function, arg)

	node.Lock()
		node.modules = append(node.modules, module)
	node.Unlock()

	return node
}

// Adds a module that will be executed upon reaching a handler
func (node *Node) Mid(function MiddlewareFunction, arg interface{}) *Node {

	middleware := &Middleware{
		node.Config,
		function,
		arg,
	}

	node.Lock()
		node.middleware = append(node.middleware, middleware)
	node.Unlock()

	return node
}

// execute init function added with .Init(...)
func (node *Node) RunModule(req web.RequestInterface) *web.ResponseStatus {

	node.RLock()
		module := node.module
	node.RUnlock()

	if module != nil {

		status := module.Run(req); if status != nil { return status }
	}

	return nil
}

// execute all module functions added with .Mod(...)
func (node *Node) RunModules(req web.RequestInterface) *web.ResponseStatus {

	for _, module := range node.modules {

		status := module.Run(req); if status != nil { return status }
	}

	return nil
}

// traversal

// finds next node according to supplied URL path segment
func (node *Node) Next(req web.RequestInterface, pathSegment string) (*Node, *web.ResponseStatus) {

	// execute any init module(s)

	if status := node.RunModule(req); status != nil {
		return nil, status
	}

	// check for child routes

	next := node.routes[pathSegment]

	if next != nil { return next, nil }

	// check for path param

	next = node.param; if next == nil { return nil, nil }

	if next.Validation != nil {

		status, value := next.Validation.PathFunction(req, pathSegment)
		if status != nil {

			status.Message = fmt.Sprintf("%s KEY(%v)", status.MessageString(), pathSegment)

			//return nil, &web.ResponseStatus{nil, 400, fmt.Sprintf("UNEXPECTED VALUE  %v, %v", pathSegment, next.Validation.Expecting())}
			return nil, status
		}

		// write route params into request object

		for _, key := range next.Validation.Keys { req.SetParam(key, value) }

	}

	return next, nil
}

// Returns the handler assciated with the HTTP request method.
func (node *Node) handler(req web.RequestInterface) *Handler {

	node.RLock()

		handler := node.methods[req.Method()]

	node.RUnlock()

	return handler
}

// Adds a file to be served from the specified path.
func (node *Node) File(path string) *Node {

	h := &Handler{
		IsFile:	true,
		filePath: path,
	}

	node.addHandler("GET", h)

	return node
}

// Walks through the specified folder to mirror the file structure for files containing all filters
func (node *Node) Folder(directoryPath string, filters ...string) *Node {

	// remove trailing slash from the directory path if existing
	directoryPath = strings.TrimSuffix(directoryPath, "/")

	go func () {

		for {
			f, err := os.Open(directoryPath); if err != nil { panic(err) }

			names, err := f.Readdirnames(-1)
			f.Close()
			if err != nil { panic(err) }

			for _, name := range names {

				path := strings.Replace(directoryPath + "/" + name, "//", "/", -1)

				node.checkFile(name, path, filters)

			}

			time.Sleep(5 * time.Second)
		}

	}()

	return node
}

// Walks through the specified folder to mirror the file structure for files containing all filters
func (node *Node) StaticFolder(directoryPath string, filters ...string) *Node {

	// remove trailing slash from the directory path if existing
	directoryPath = strings.TrimSuffix(directoryPath, "/")

	f, err := os.Open(directoryPath); if err != nil { panic(err) }

	names, err := f.Readdirnames(-1)
	f.Close()
	if err != nil { panic(err) }

	for _, name := range names {

		path := strings.Replace(directoryPath + "/" + name, "//", "/", -1)

		node.checkFile(name, path, filters)

	}

	return node
}

// Checks if file or folder, adding any files
func (node *Node) checkFile(name, path string, filters []string) {

	info, err := os.Lstat(path)
	if err != nil {
		panic(err)
	}
	// check if the file is a directory
	if info.Mode() & os.ModeSymlink != 0 {
		linked, err := os.Readlink(path)
		if err != nil {
			panic(err)
		}
		s := strings.Split(path, "/")
		path = strings.Join(
			append(
				s[:len(s)-1],
				linked,
			),
			"/",
		)
		info, err = os.Lstat(path)
		if err != nil {
			panic(err)
		}
	}
	// check if path is a directory
	if info.IsDir() {
		node.Add(name).StaticFolder(path)
		return
	}

	for _, filter := range filters {
		if !strings.Contains(name, filter) { return }
	}

	node.Add(name).File(path)

}
