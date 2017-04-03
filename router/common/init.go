package common

import 	(
		"fmt"
		//
		"github.com/golangdaddy/tarantula/web/validation"
		)

// for internal debugging use
func debug(s string) { fmt.Println(s) }

var globalNode *Node

func rootNode() *Node {

	return &Node{
		routes:	map[string]*Node{},
		methods: map[string]*Handler{},
		requestParams: map[string]interface{}{},
		modules: []*Module{},
		Validations: []*validation.Config{},
	}

}

func init() {

	globalNode = rootNode()
	globalNode.Config = newConfig()

}

func newConfig() *Config {
	return &Config{
		cacheFiles: true,
		SubdomainTrees: map[string]*Node{},
		Handlers: []*Handler{},
		lDelim: "{{",
		rDelim: "}}",
	}
}

func Root() *Node {

	return globalNode
}
