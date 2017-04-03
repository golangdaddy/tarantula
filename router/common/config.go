package common

import 	(
		"sync"
		"sort"
		//
		"github.com/golangdaddy/tarantula/log"
		"github.com/golangdaddy/tarantula/web"
		)

type Config struct {
	Log logging.Logger
// ProjectName is for App Engine apps
	ProjectName string
	Host string
	RootRegistry Registry
	SubdomainTrees map[string]*Node
	headers Headers
	lDelim, rDelim string
	Handlers []*Handler
	cacheFiles bool
	forceTLS bool
	sync.RWMutex
}

func (config *Config) NoCache() {

	config.Lock()
		config.cacheFiles = false
	config.Unlock()
}

func (config *Config) SubTree(subdomain string) *Node {

	config.RLock()
		tree := config.SubdomainTrees[subdomain]
	config.RUnlock()

	if tree == nil {

		tree = rootNode()

		newConfig := *config

		tree.Config = &newConfig

		config.Lock()
			config.SubdomainTrees[subdomain] = tree
		config.Unlock()
	}

	return tree
}

// block all non-https requests
func (config *Config) ForceTLS() {

	config.Lock()
		config.forceTLS = true
	config.Unlock()
}


func (config *Config) SetDelims(l, r string) {

	config.Lock()
		config.lDelim = l
		config.rDelim = r
	config.Unlock()
}

// Sets the root registry to the specified map
func (config *Config) SetRootRegistry(reg Registry) {

	config.Lock()
		config.RootRegistry = reg
	config.Unlock()
}

// Sets the http preflight headers to the specified map
func (config *Config) SetHeaders(h Headers) {

	config.Lock()
		config.headers = h
	config.Unlock()
}

// Returns the root function if present in the registry
func (config *Config) GetRootFunction(functionKey string) func (req web.RequestInterface) *web.ResponseStatus {

	if config.RootRegistry == nil { return nil }

	config.RLock()
		function := config.RootRegistry[functionKey]
	config.RUnlock()

	return function
}

// Builds the handler documentation object.
func (config *Config) buildSpec(req web.RequestInterface) []*HandlerSpec {

	a := HandlerArray{}

	config.RLock()

	for _, handler := range config.Handlers {
		a = append(a, handler.Spec(req))
	}

	config.RUnlock()

	sort.Sort(a)

	return a
}
